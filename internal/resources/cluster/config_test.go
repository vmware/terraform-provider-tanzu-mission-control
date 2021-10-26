/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

const testDefaultAttachClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

	  {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = false
	}
`

const testAttachClusterWithKubeConfigScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

	  attach_k8s_cluster {
		kubeconfig_file = "{{.KubeConfigPath}}"
		description     = "optional description about the kube-config provided"
	  }

	  {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = true
	}
`

const testDataSourceAttachClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

	  {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = false
	}

	data {{.ResourceName}} {{.DataSourceNameVar}} {
		management_cluster_name = {{.ResourceName}}.{{.ResourceNameVar}}.management_cluster_name
		provisioner_name        = {{.ResourceName}}.{{.ResourceNameVar}}.provisioner_name
		name                    = {{.ResourceName}}.{{.ResourceNameVar}}.name
	}
`
