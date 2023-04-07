package providerekscluster

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provider_ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func DataSourceTMCProviderEKSCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProviderClusterRead,
		Schema:      providerClusterSchema,
	}
}

func dataSourceProviderClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterGetProviderEksClusterResponse
		err   error
	)

	clusterFn := constructFullname(d)
	getEksClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.ProviderEKSClusterResourceService.ProviderEksClusterResourceServiceGet(clusterFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) {
				d.SetId("")
				return false, nil
			}

			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", clusterFn.Name)
		}

		d.SetId(resp.ProviderEksCluster.Meta.UID)

		if ctx.Value(contextMethodKey{}) == "create" &&
			resp.ProviderEksCluster.Status.Phase != nil &&
			*resp.ProviderEksCluster.Status.Phase != models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhaseMANAGED {
			log.Printf("[DEBUG] waiting for cluster(%s) to be in READY phase", clusterFn.ToString())
			return true, nil
		}

		return false, nil
	}

	timeoutValueData, _ := d.Get(waitKey).(string)

	if ctx.Value(contextMethodKey{}) != "create" {
		timeoutValueData = "do_not_retry"
	}

	switch timeoutValueData {
	case "do_not_retry":
		_, err = getEksClusterResourceRetryableFn()
	case "":
		fallthrough
	case "default":
		timeoutValueData = "5m"

		fallthrough
	default:
		timeoutDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr != nil {
			log.Printf("[INFO] unable to prase the duration value for the key %s. Defaulting to 5 minutes(5m)"+
				" Please refer to 'https://pkg.go.dev/time#ParseDuration' for providing the right value", waitKey)

			timeoutDuration = defaultTimeout
		}

		_, err = helper.RetryUntilTimeout(getEksClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	}

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control Provider EKS cluster entry, name : %s", clusterFn.Name))
	}

	// always run
	d.SetId(resp.ProviderEksCluster.Meta.UID)

	status := map[string]interface{}{
		// TODO: add condition
		"phase": resp.ProviderEksCluster.Status.Phase,
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.ProviderEksCluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenClusterSpec(resp.ProviderEksCluster.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func flattenClusterSpec(item *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[clusterGroupKey] = item.ClusterGroupName

	if item.ProxyName != "" {
		data[proxyNameKey] = item.ProxyName
	}

	if item.AgentName != "" {
		data[agentNameKey] = item.AgentName
	}

	if item.Arn != "" {
		data[eksARNKey] = item.Arn
	}

	return []interface{}{data}
}
