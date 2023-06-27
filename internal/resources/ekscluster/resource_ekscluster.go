/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

var ignoredTagsPrefix = "tmc.cloud.vmware.com/"

const defaultTimeout = 3 * time.Minute

func ResourceTMCEKSCluster() *schema.Resource {
	return &schema.Resource{
		Schema:        clusterSchema,
		CreateContext: resourceClusterCreate,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCEKSClusterRead(helper.GetContextWithCaller(ctx, helper.RefreshState), d, m)
		},
		UpdateContext: resourceClusterInPlaceUpdate,
		DeleteContext: resourceClusterDelete,
		Description:   "Tanzu Mission Control EKS Cluster Resource",
	}
}

var clusterSchema = map[string]*schema.Schema{
	CredentialNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the AWS Credential in Tanzu Mission Control",
		Required:    true,
		ForceNew:    true,
	},
	RegionKey: {
		Type:        schema.TypeString,
		Description: "AWS Region of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        clusterSpecSchema,
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero",
		Default:     "default",
		Optional:    true,
	},
}

var clusterSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group to which this cluster belongs",
				Default:     clusterGroupDefaultValue,
				Optional:    true,
			},
			proxyNameKey: {
				Type:        schema.TypeString,
				Description: "Optional proxy name is the name of the Proxy Config to be used for the cluster",
				Optional:    true,
			},
			configKey: configSchema,
			nodepoolKey: {
				Type:        schema.TypeList,
				Description: "Nodepool definitions for the cluster",
				Required:    true,
				MinItems:    1,
				Elem:        nodepoolDefinitionSchema,
			},
		},
	},
}

var configSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "EKS config for the cluster control plane",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleArnKey: {
				Type:        schema.TypeString,
				Description: "ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations",
				Required:    true,
				ForceNew:    true,
			},
			kubernetesVersionKey: {
				Type:        schema.TypeString,
				Description: "Kubernetes version of the cluster",
				Required:    true,
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "The metadata to apply to the cluster to assist with categorization and organization",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.Contains(k, ignoredTagsPrefix)
				},
			},
			kubernetesNetworkConfigKey: {
				Type:        schema.TypeList,
				Description: "Kubernetes Network Config",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						serviceCidrKey: {
							Type:        schema.TypeString,
							Description: "Service CIDR for Kubernetes services",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			loggingKey: {
				Type:        schema.TypeList,
				Description: "EKS logging configuration",
				Optional:    true,
				ForceNew:    false,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						apiServerKey: {
							Type:        schema.TypeBool,
							Description: "Enable API server logs",
							Optional:    true,
							ForceNew:    false,
						},
						auditKey: {
							Type:        schema.TypeBool,
							Description: "Enable audit logs",
							Optional:    true,
							ForceNew:    false,
						},
						authenticatorKey: {
							Type:        schema.TypeBool,
							Description: "Enable authenticator logs",
							Optional:    true,
							ForceNew:    false,
						},
						controllerManagerKey: {
							Type:        schema.TypeBool,
							Description: "Enable controller manager logs",
							Optional:    true,
							ForceNew:    false,
						},
						schedulerKey: {
							Type:        schema.TypeBool,
							Description: "Enable scheduler logs",
							Optional:    true,
							ForceNew:    false,
						},
					},
				},
			},
			vpcKey: vpcSchema,
		},
	},
}

var vpcSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VPC config",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			enablePrivateAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable Kubernetes API requests within your cluster's VPC (such as node to control plane communication) use the private VPC endpoint (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))",
				Optional:    true,
				ForceNew:    false,
			},
			enablePublicAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable cluster API server access from the internet. You can, optionally, limit the CIDR blocks that can access the public endpoint using public_access_cidrs (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))",
				Optional:    true,
				ForceNew:    false,
			},
			publicAccessCidrsKey: {
				Type:        schema.TypeSet,
				Description: "Specify which addresses from the internet can communicate to the public endpoint, if public endpoint is enabled (see [Amazon EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html))",
				Optional:    true,
				ForceNew:    false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			securityGroupsKey: {
				Type:        schema.TypeSet,
				Description: "Security groups for the cluster VMs",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    5,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			subnetIdsKey: {
				Type:        schema.TypeSet,
				Description: "Subnet ids used by the cluster (see [Amazon EKS VPC and subnet requirements and considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html#network-requirements-subnets))",
				Required:    true,
				ForceNew:    true,
				MinItems:    2,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

func constructEksClusterSpec(d *schema.ResourceData) (spec *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) {
	spec = &eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec{
		ClusterGroupName: clusterGroupDefaultValue,
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ClusterGroupName, clusterGroupKey)
	}

	if v, ok := specData[proxyNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ProxyName, proxyNameKey)
	}

	if v, ok := specData[configKey]; ok {
		configData, _ := v.([]interface{})
		spec.Config = constructConfig(configData)
	}

	return spec
}

func constructConfig(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
	config := &eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig{}

	if len(data) == 0 || data[0] == nil {
		return config
	}

	configData, _ := data[0].(map[string]interface{})

	if v, ok := configData[roleArnKey]; ok {
		helper.SetPrimitiveValue(v, &config.RoleArn, roleArnKey)
	}

	if v, ok := configData[kubernetesVersionKey]; ok {
		helper.SetPrimitiveValue(v, &config.Version, kubernetesVersionKey)
	}

	if v, ok := configData[tagsKey]; ok {
		data, _ := v.(map[string]interface{})
		config.Tags = constructStringMap(data)
	}

	if v, ok := configData[kubernetesNetworkConfigKey]; ok {
		data, _ := v.([]interface{})
		config.KubernetesNetworkConfig = constructNetworkData(data)
	}

	if v, ok := configData[loggingKey]; ok {
		data, _ := v.([]interface{})
		config.Logging = constructLogging(data)
	}

	if v, ok := configData[vpcKey]; ok {
		data, _ := v.([]interface{})
		config.Vpc = constructVpc(data)
	}

	return config
}

func constructNetworkData(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig {
	nc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig{}
	if len(data) == 0 || data[0] == nil {
		return nc
	}

	networkData, _ := data[0].(map[string]interface{})
	if v, ok := networkData[serviceCidrKey]; ok {
		helper.SetPrimitiveValue(v, &nc.ServiceCidr, serviceCidrKey)
	}

	return nc
}

func constructLogging(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging {
	logging := &eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging{}

	if len(data) == 0 || data[0] == nil {
		return logging
	}

	loggingData, _ := data[0].(map[string]interface{})

	if v, ok := loggingData[apiServerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.APIServer, apiServerKey)
	}

	if v, ok := loggingData[auditKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Audit, auditKey)
	}

	if v, ok := loggingData[authenticatorKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Authenticator, authenticatorKey)
	}

	if v, ok := loggingData[controllerManagerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.ControllerManager, controllerManagerKey)
	}

	if v, ok := loggingData[schedulerKey]; ok {
		helper.SetPrimitiveValue(v, &logging.Scheduler, schedulerKey)
	}

	return logging
}

func constructVpc(data []interface{}) *eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig {
	vpc := &eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig{}

	if len(data) == 0 || data[0] == nil {
		return vpc
	}

	vpcData, _ := data[0].(map[string]interface{})

	if v, ok := vpcData[enablePrivateAccessKey]; ok {
		helper.SetPrimitiveValue(v, &vpc.EnablePrivateAccess, enablePrivateAccessKey)
	}

	if v, ok := vpcData[enablePublicAccessKey]; ok {
		helper.SetPrimitiveValue(v, &vpc.EnablePublicAccess, enablePublicAccessKey)
	}

	if v, ok := vpcData[publicAccessCidrsKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			vpc.PublicAccessCidrs = constructStringList(data.List())
		}
	}

	if v, ok := vpcData[securityGroupsKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			vpc.SecurityGroups = constructStringList(data.List())
		}
	}

	if v, ok := vpcData[subnetIdsKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			vpc.SubnetIds = constructStringList(data.List())
		}
	}

	return vpc
}

func constructStringMap(data map[string]interface{}) map[string]string {
	out := make(map[string]string)

	for k, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, valueKey)

		out[k] = value
	}

	return out
}

func constructStringList(data []interface{}) []string {
	out := make([]string, 0, len(data))

	for _, v := range data {
		var value string

		helper.SetPrimitiveValue(v, &value, "")

		out = append(out, value)
	}

	return out
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, ok := m.(authctx.TanzuContext)
	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	clusterFn := constructClusterFullname(d)
	clusterSpec := constructEksClusterSpec(d)

	clusterReq := &eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
		EksCluster: &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
			FullName: clusterFn,
			Meta:     common.ConstructMeta(d),
			Spec:     clusterSpec,
		},
	}

	var eksCluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster

	clusterResponse, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceCreate(clusterReq)
	if err != nil {
		if !clienterrors.IsAlreadyExistsError(err) {
			return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
		}

		clusterResponse, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(clusterFn)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
		}

		eksCluster = clusterResponse.EksCluster
	} else {
		eksCluster = clusterResponse.EksCluster
	}

	d.SetId(eksCluster.Meta.UID)

	return dataSourceTMCEKSClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceDelete(constructClusterFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	clusterFn := constructClusterFullname(d)
	getClusterResourceRetryableFn := func() (retry bool, err error) {
		_, err = config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(clusterFn)
		if err == nil {
			log.Printf("[DEBUG] cluster(%s) deletion in progress", clusterFn.ToString())
			return true, errors.New("cluster deletion in progress")
		}

		if !clienterrors.IsNotFoundError(err) {
			return true, err
		}

		return false, nil
	}

	timeoutDuration := getRetryTimeout(d)

	_, err = helper.RetryUntilTimeout(getClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	if err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s EKS cluster resource clean up", d.Get(NameKey)))
	}

	return diags
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	// Get call to initialise the cluster struct
	getResp, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(constructClusterFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	clusterSpec := constructEksClusterSpec(d)
	errcl := handleClusterDiff(config, getResp.EksCluster, common.ConstructMeta(d), clusterSpec)
	if errcl != nil {
		return diag.FromErr(errors.Wrapf(errcl, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	log.Printf("[INFO] cluster update successful")

	return dataSourceTMCEKSClusterRead(ctx, d, m)
}

func handleClusterDiff(config authctx.TanzuContext, tmcCluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, clusterSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) error {
	updateCluster := false

	if meta.Description != tmcCluster.Meta.Description ||
		!reflect.DeepEqual(meta.Labels, tmcCluster.Meta.Labels) {
		updateCluster = true
		tmcCluster.Meta.Description = meta.Description
		tmcCluster.Meta.Labels = meta.Labels
	}

	if !clusterSpecEqual(clusterSpec, tmcCluster.Spec) {
		updateCluster = true
	}

	// The TF update request was only for nodepools.
	// No need to update Cluster, it will error out.
	if !updateCluster {
		return nil
	}

	// there is some translation error, which results
	// in mismatch on the server.
	tmcCluster.Meta.CreationTime = strfmt.DateTime{}

	newCluster := &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
		FullName: tmcCluster.FullName,
		Meta:     tmcCluster.Meta,
		Spec:     clusterSpec,
	}

	_, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceUpdate(
		&eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
			EksCluster: newCluster,
		},
	)

	if err != nil {
		return errors.Wrapf(err, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", tmcCluster.FullName.Name)
	}

	return nil
}

func constructClusterFullname(d *schema.ResourceData) (fullname *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) {
	fullname = &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{}

	fullname.CredentialName, _ = d.Get(CredentialNameKey).(string)
	fullname.Region, _ = d.Get(RegionKey).(string)
	fullname.Name, _ = d.Get(NameKey).(string)

	return fullname
}

func getRetryTimeout(d *schema.ResourceData) time.Duration {
	timeoutValueData, _ := d.Get(waitKey).(string)
	if timeoutValueData != "default" {
		providedDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr == nil {
			return providedDuration
		}
	}

	return defaultTimeout
}

func flattenClusterSpec(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec, nodepools []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[clusterGroupKey] = item.ClusterGroupName

	if item.Config != nil {
		data[configKey] = flattenConfig(item.Config)
	}

	if len(nodepools) > 0 {
		data[nodepoolKey] = flattenNodePools(nodepools)
	}

	if item.ProxyName != "" {
		data[proxyNameKey] = item.ProxyName
	}

	return []interface{}{data}
}

func flattenConfig(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.KubernetesNetworkConfig != nil {
		data[kubernetesNetworkConfigKey] = flattenKubernetesNetworkConfig(item.KubernetesNetworkConfig)
	}

	if item.Logging != nil {
		data[loggingKey] = flattenLogging(item.Logging)
	}

	data[roleArnKey] = item.RoleArn
	data[tagsKey] = item.Tags
	data[kubernetesVersionKey] = item.Version

	if item.Vpc != nil {
		data[vpcKey] = flattenVpc(item.Vpc)
	}

	return []interface{}{data}
}

func flattenKubernetesNetworkConfig(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[serviceCidrKey] = item.ServiceCidr

	return []interface{}{data}
}

func flattenLogging(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[apiServerKey] = item.APIServer
	data[auditKey] = item.Audit
	data[authenticatorKey] = item.Authenticator
	data[controllerManagerKey] = item.ControllerManager
	data[schedulerKey] = item.Scheduler

	return []interface{}{data}
}

func flattenVpc(item *eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[enablePrivateAccessKey] = item.EnablePrivateAccess
	data[enablePublicAccessKey] = item.EnablePublicAccess

	if len(item.PublicAccessCidrs) > 0 {
		data[publicAccessCidrsKey] = item.PublicAccessCidrs
	}

	if len(item.SecurityGroups) > 0 {
		data[securityGroupsKey] = item.SecurityGroups
	}

	if len(item.SubnetIds) > 0 {
		data[subnetIdsKey] = item.SubnetIds
	}

	return []interface{}{data}
}
