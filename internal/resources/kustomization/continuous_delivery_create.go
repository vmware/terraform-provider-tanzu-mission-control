// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomization

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	continuousdeliveryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/cluster"
	continuousdeliveryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/continuousdelivery/clustergroup"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kustomization/scope"
)

func enableContinuousDelivery(config *authctx.TanzuContext, scopedFullnameData *scope.ScopedFullname, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta) error {
	if config == nil || scopedFullnameData == nil || meta == nil {
		return errors.New("missing variables: error while enabling Tanzu Mission Control cluster continuous delivery feature")
	}

	switch scopedFullnameData.Scope {
	case commonscope.ClusterScope:
		if scopedFullnameData.FullnameCluster != nil {
			continuousDeliveryReq := &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest{
				ContinuousDelivery: &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery{
					FullName: &continuousdeliveryclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName{
						ClusterName:           scopedFullnameData.FullnameCluster.ClusterName,
						ManagementClusterName: scopedFullnameData.FullnameCluster.ManagementClusterName,
						ProvisionerName:       scopedFullnameData.FullnameCluster.ProvisionerName,
					},
					Meta: meta,
				},
			}

			_, err := config.TMCConnection.ClusterContinuousDeliveryResourceService.VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryResourceServiceCreate(continuousDeliveryReq)
			if err != nil && status.Code(err) != codes.AlreadyExists {
				return err
			}
		}
	case commonscope.ClusterGroupScope:
		if scopedFullnameData.FullnameClusterGroup != nil {
			continuousDeliveryReq := &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest{
				ContinuousDelivery: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery{
					FullName: &continuousdeliveryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName{
						ClusterGroupName: scopedFullnameData.FullnameClusterGroup.ClusterGroupName,
					},
					Meta: meta,
				},
			}

			_, err := config.TMCConnection.ClusterGroupContinuousDeliveryResourceService.VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryResourceServiceCreate(continuousDeliveryReq)
			if err != nil && status.Code(err) != codes.AlreadyExists {
				return err
			}
		}
	case commonscope.UnknownScope:
		return fmt.Errorf("no valid scope type block found: minimum one valid scope type block is required among: %v. Please check the schema", strings.Join(scope.ScopesAllowed[:], `, `))
	}

	return nil
}
