/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
	secretexportmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster/secretexport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/kubernetessecret/scope"
)

func constructFullname(d *schema.ResourceData) (fullname *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) {
	name, _ := d.Get(NameKey).(string)

	fullname = (scope.ConstructScope(d, name)).FullnameCluster

	fullname.NamespaceName, _ = d.Get(NamespaceNameKey).(string)

	fullname.OrgID, _ = d.Get(OrgIDKey).(string)

	return fullname
}

func constructFullnameSecetExport(d *schema.ResourceData) (fullname *secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName) {
	fullname = &secretexportmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportFullName{}

	secretFullname := constructFullname(d)

	fullname.ClusterName = secretFullname.ClusterName

	fullname.ManagementClusterName = secretFullname.ManagementClusterName

	fullname.Name = secretFullname.Name

	fullname.NamespaceName = secretFullname.NamespaceName

	fullname.OrgID = secretFullname.OrgID

	fullname.ProvisionerName = secretFullname.ProvisionerName

	return fullname
}
