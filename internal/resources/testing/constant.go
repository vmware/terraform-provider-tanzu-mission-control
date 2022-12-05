/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// testing is a helper package created for testing purpose.
// go linker would not include this package in the binary, as it is not imported anywhere else other than for testing

package testing

const (
	providerName = "tanzu-mission-control"
	value1       = "value1"
	value2       = "value2"
	description  = "resource with description"
)

const TestEksClusterPostValue = "{\"eksCluster\":{\"fullName\":{\"credentialName\":\"TEST_CREDENTIAL\",\"name\":\"terraform-eks-test\",\"region\":\"us-west-2\"},\"meta\":{\"creationTime\":\"0001-01-01T00:00:00.000Z\",\"description\":\"resource with description\",\"labels\":{\"key1\":\"value1\",\"key2\":\"value2\"},\"parentReferences\":null,\"updateTime\":\"0001-01-01T00:00:00.000Z\"},\"spec\":{\"clusterGroupName\":\"default\",\"config\":{\"kubernetesNetworkConfig\":{\"serviceCidr\":\"10.100.0.0/16\"},\"logging\":{\"audit\":true,\"authenticator\":true,\"controllerManager\":true,\"scheduler\":true},\"roleArn\":\"arn:aws:iam::919197287370:role/control-plane.TEST_CLOUD_FORMATION_TEMPLATE_ID.eks.tmc.cloud.vmware.com\",\"version\":\"1.23\",\"vpc\":{\"enablePrivateAccess\":true,\"enablePublicAccess\":true,\"publicAccessCidrs\":[\"0.0.0.0/0\"],\"securityGroups\":[\"sg-0a6768722e9716768\"],\"subnetIds\":[\"subnet-0ed95d5c212ac62a1\",\"subnet-06897e1063cc0cf4e\",\"subnet-0526ecaecde5b1bf7\",\"subnet-0a184f6302af32a86\"]}},\"nodePools\":[{\"info\":{\"description\":\"tf nodepool description\",\"name\":\"first-np\"},\"spec\":{\"amiType\":\"AL2_x86_64\",\"capacityType\":\"ON_DEMAND\",\"instanceTypes\":[\"t3.medium\",\"m3.large\"],\"launchTemplate\":{},\"nodeLabels\":{\"testnplabelkey\":\"testnplabelvalue\"},\"remoteAccess\":{\"securityGroups\":[\"sg-0a6768722e9716768\"],\"sshKey\":\"anshulc\"},\"roleArn\":\"arn:aws:iam::919197287370:role/worker.TEST_CLOUD_FORMATION_TEMPLATE_ID.eks.tmc.cloud.vmware.com\",\"rootDiskSize\":40,\"scalingConfig\":{\"desiredSize\":4,\"maxSize\":8,\"minSize\":1},\"subnetIds\":[\"subnet-0ed95d5c212ac62a1\",\"subnet-06897e1063cc0cf4e\",\"subnet-0526ecaecde5b1bf7\",\"subnet-0a184f6302af32a86\"],\"tags\":{\"testnptag\":\"testnptagvalue\"},\"taints\":[],\"updateConfig\":{\"maxUnavailableNodes\":\"2\"}}},{\"info\":{\"description\":\"tf nodepool 2 description\",\"name\":\"second-np\"},\"spec\":{\"instanceTypes\":[],\"launchTemplate\":{\"name\":\"vivek\",\"version\":\"7\"},\"nodeLabels\":{\"testnplabelkey\":\"testnplabelvalue\"},\"remoteAccess\":{\"securityGroups\":null},\"roleArn\":\"arn:aws:iam::919197287370:role/worker.TEST_CLOUD_FORMATION_TEMPLATE_ID.eks.tmc.cloud.vmware.com\",\"scalingConfig\":{\"desiredSize\":4,\"maxSize\":8,\"minSize\":1},\"subnetIds\":[\"subnet-0ed95d5c212ac62a1\",\"subnet-06897e1063cc0cf4e\",\"subnet-0526ecaecde5b1bf7\",\"subnet-0a184f6302af32a86\"],\"tags\":{\"testnptag\":\"testnptagvalue\"},\"taints\":[{\"effect\":\"PREFER_NO_SCHEDULE\",\"key\":\"randomkey\",\"value\":\"randomvalue\"}],\"updateConfig\":{\"maxUnavailablePercentage\":\"12\"}}}]}}}"
