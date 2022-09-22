/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
)

var ErrMissingIntegrationClientService = errors.New("missing integration client service")

type defaultClient struct{}

func (d defaultClient) ManageV1alpha1ClusterIntegrationResourceServiceCreate(*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse, error) {
	panic(ErrMissingIntegrationClientService)
}

func (d defaultClient) ManageV1alpha1ClusterIntegrationResourceServiceRead(*integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse, error) {
	panic(ErrMissingIntegrationClientService)
}

func (d defaultClient) ManageV1alpha1ClusterIntegrationResourceServiceDelete(*integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) error {
	panic(ErrMissingIntegrationClientService)
}
