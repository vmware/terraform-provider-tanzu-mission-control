/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretexportclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
	secretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	secretexportclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup/secretexport"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/spec"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/status"
)

type dataFromServer struct {
	UID                     string
	secretExportRespNil     bool
	secretExportErr         error
	meta                    *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta
	atomicSpec              *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec
	clusterScopeStatus      *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretStatus
	clusterGroupScopeStatus *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretStatus
}

func DataSourceSecret() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceSecretRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		Schema: getDataSourceSchema(),
	}
}

func dataSourceSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	secretName, ok := d.Get(NameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret name")
	}

	secretNamespaceName, ok := d.Get(NamespaceNameKey).(string)
	if !ok {
		return diag.Errorf("Unable to read secret namespace name")
	}

	scopedFullnameData := scope.ConstructScope(d, secretName, secretNamespaceName)

	if scopedFullnameData == nil {
		return diag.Errorf("Unable to get Tanzu Mission Control secret entry; Scope full name is empty")
	}

	secretDataFromServer, err := retrieveSecretDataFromServer(config, scopedFullnameData, d)
	if err != nil {
		if clienterrors.IsNotFoundError(err) && !helper.IsDataRead(ctx) {
			_ = schema.RemoveFromState(d, m)
			return diags
		}

		return diag.FromErr(err)
	}

	d.SetId(secretDataFromServer.UID)

	var password string

	var opaqueData map[string]interface{}

	if _, ok := d.GetOk(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.PasswordKey)); ok {
		password, _ = (d.Get(helper.GetFirstElementOf(spec.SpecKey, spec.DockerConfigjsonKey, spec.PasswordKey))).(string)
	}

	if opData, ok := d.GetOk(helper.GetFirstElementOf(spec.SpecKey, spec.OpaqueKey)); ok && opData != nil {
		opaqueData = opData.(map[string]interface{})
	}

	if d.Get(ExportKey).(bool) {
		if secretDataFromServer.secretExportErr != nil || secretDataFromServer.secretExportRespNil {
			switch {
			case clienterrors.IsNotFoundError(err):
				if err := d.Set(ExportKey, false); err != nil {
					return diag.FromErr(err)
				}
			default:
				return diag.FromErr(errors.Wrapf(secretDataFromServer.secretExportErr, "Unable to get Tanzu Mission Control secret export entry, name : %s", secretName))
			}
		}
	} else {
		switch {
		case secretDataFromServer.secretExportErr == nil && !(secretDataFromServer.secretExportRespNil):
			if err := d.Set(ExportKey, true); err != nil {
				return diag.FromErr(err)
			}
		case secretDataFromServer.secretExportErr != nil && clienterrors.IsNotFoundError(secretDataFromServer.secretExportErr):
			if err := d.Set(ExportKey, false); err != nil {
				return diag.FromErr(err)
			}
		default:
			return diag.FromErr(errors.Wrapf(secretDataFromServer.secretExportErr, "Unable to get Tanzu Mission Control secret export SWITCH entry, name : %s", secretName))
		}
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(secretDataFromServer.meta)); err != nil {
		return diag.FromErr(err)
	}

	var (
		flattenedSpec   []interface{}
		flattenedStatus interface{}
	)

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		flattenedSpec = spec.FlattenSpecForClusterScope(secretDataFromServer.atomicSpec, password, opaqueData)
		flattenedStatus = status.FlattenStatusForClusterScope(secretDataFromServer.clusterScopeStatus)
	case commonscope.ClusterGroupScope:
		clusterGroupScopeSpec := &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{
			AtomicSpec: secretDataFromServer.atomicSpec,
		}
		flattenedSpec = spec.FlattenSpecForClusterGroupScope(clusterGroupScopeSpec, password, opaqueData)
		flattenedStatus = status.FlattenStatusForClusterGroupScope(secretDataFromServer.clusterGroupScopeStatus)
	}

	if err := d.Set(spec.SpecKey, flattenedSpec); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(statusKey, flattenedStatus); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func retrieveSecretDataFromServer(config authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, d *schema.ResourceData) (*dataFromServer, error) {
	var secretDataFromServer = &dataFromServer{}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			resp, err := config.TMCConnection.SecretResourceService.SecretResourceServiceGet(scopedFullnameData.FullnameCluster)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return secretDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster secret entry, name : %s", scopedFullnameData.FullnameCluster.Name)
				}

				return secretDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster secret entry, name : %s", scopedFullnameData.FullnameCluster.Name)
			}

			scopedFullnameData.FullnameCluster = resp.Secret.FullName
			secretDataFromServer.UID = resp.Secret.Meta.UID
			secretDataFromServer.meta = resp.Secret.Meta
			secretDataFromServer.atomicSpec = resp.Secret.Spec
			secretDataFromServer.clusterScopeStatus = resp.Secret.Status

			secretExportResp, secretExportErr := config.TMCConnection.SecretExportResourceService.SecretExportResourceServiceGet(
				&secretexportclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName{
					Name:                  scopedFullnameData.FullnameCluster.Name,
					ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
					ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
					ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
					NamespaceName:         scopedFullnameData.FullnameCluster.NamespaceName,
					OrgID:                 scopedFullnameData.FullnameCluster.OrgID,
				},
			)

			secretDataFromServer.secretExportRespNil = (secretExportResp == nil)
			secretDataFromServer.secretExportErr = secretExportErr
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			resp, err := config.TMCConnection.ClusterGroupSecretResourceService.SecretResourceServiceGet(scopedFullnameData.FullnameClusterGroup)
			if err != nil {
				if clienterrors.IsNotFoundError(err) {
					d.SetId("")
					return secretDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group secret entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
				}

				return secretDataFromServer, errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster group secret entry, name : %s", scopedFullnameData.FullnameClusterGroup.Name)
			}

			scopedFullnameData.FullnameClusterGroup = resp.Secret.FullName
			secretDataFromServer.UID = resp.Secret.Meta.UID
			secretDataFromServer.meta = resp.Secret.Meta
			secretDataFromServer.atomicSpec = resp.Secret.Spec.AtomicSpec
			secretDataFromServer.clusterGroupScopeStatus = resp.Secret.Status

			secretExportResp, secretExporterr := config.TMCConnection.ClusterGroupSecretExportResourceService.SecretExportResourceServiceGet(
				&secretexportclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName{
					Name:             scopedFullnameData.FullnameClusterGroup.Name,
					ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
					NamespaceName:    scopedFullnameData.FullnameClusterGroup.NamespaceName,
					OrgID:            scopedFullnameData.FullnameClusterGroup.OrgID,
				},
			)

			secretDataFromServer.secretExportRespNil = (secretExportResp == nil)
			secretDataFromServer.secretExportErr = secretExporterr
		}
	case commonscope.UnknownScope:
		return secretDataFromServer, errors.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema.", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	fullName, name, namespace := scope.FlattenScope(scopedFullnameData)

	if err := d.Set(NameKey, name); err != nil {
		return secretDataFromServer, err
	}

	if err := d.Set(NamespaceNameKey, namespace); err != nil {
		return secretDataFromServer, err
	}

	if err := d.Set(commonscope.ScopeKey, fullName); err != nil {
		return secretDataFromServer, err
	}

	return secretDataFromServer, nil
}
