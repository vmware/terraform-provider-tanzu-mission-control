// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocationtests

import (
	"fmt"
	"strings"

	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	clustergroupres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
	targetlocationres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/targetlocation"
)

type ResourceBuildMode string

const (
	RsFullBuild                  ResourceBuildMode = "FULL"
	RsClusterOnly                ResourceBuildMode = "CLUSTER"
	RsClusterGroupOnly           ResourceBuildMode = "CLUSTER_GROUP"
	RsClusterOnlyNoParentRs      ResourceBuildMode = "CLUSTER_NO_PARENT_RESOURCE"
	RsClusterGroupOnlyNoParentRs ResourceBuildMode = "CLUSTER_GROUP_NO_PARENT_RESOURCE"
)

const (
	TmcManagedResourceName       = "test_tmc_managed"
	TargetLocationTMCManagedName = "tmc-managed-test"

	AwsSelfManagedResourceName       = "test_aws_self_managed"
	TargetLocationAWSSelfManagedName = "aws-self-managed-test"

	AzureSelfManagedResourceName       = "test_azure_self_managed"
	TargetLocationAzureSelfManagedName = "azure-self-managed-test"
)

var (
	TmcManagedResourceFullName       = fmt.Sprintf("%s.%s", targetlocationres.ResourceName, TmcManagedResourceName)
	AwsSelfManagedResourceFullName   = fmt.Sprintf("%s.%s", targetlocationres.ResourceName, AwsSelfManagedResourceName)
	AzureSelfManagedResourceFullName = fmt.Sprintf("%s.%s", targetlocationres.ResourceName, AzureSelfManagedResourceName)
)

type ResourceTFConfigBuilder struct {
	RequiredResources   string
	AssignedGroupsBlock string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, bMode ResourceBuildMode) *ResourceTFConfigBuilder {
	var (
		requiredResources         string
		clusterAssignedGroup      string
		clusterGroupAssignedGroup string
	)

	if bMode == RsFullBuild || bMode == RsClusterOnly || bMode == RsClusterOnlyNoParentRs {
		if bMode != RsClusterOnlyNoParentRs {
			requiredResources, _ = scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterScope, []string{commonscope.ClusterKey})
		}

		mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
		provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
		clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
		clusterAssignedGroup = fmt.Sprintf(`
		cluster {
			%s = %s
			%s = %s
			%s = %s
		}
		`,
			targetlocationres.ClusterScopeManagementClusterNameKey, mgmtClusterName,
			targetlocationres.ClusterScopeProvisionerNameKey, provisionerName,
			targetlocationres.ClusterScopeNameKey, clusterName)
	}

	if bMode == RsFullBuild || bMode == RsClusterGroupOnly || bMode == RsClusterGroupOnlyNoParentRs {
		if bMode != RsClusterGroupOnlyNoParentRs {
			clusterGroupRes, _ := scopeHelper.GetTestResourceHelperAndScope(commonscope.ClusterGroupScope, []string{commonscope.ClusterGroupKey})
			requiredResources = fmt.Sprintf("%s\n%s", requiredResources, clusterGroupRes)
		}

		clusterGroupAssignedGroup = fmt.Sprintf("cluster_groups = [%s.%s]", scopeHelper.ClusterGroup.ResourceName, clustergroupres.NameKey)
	}

	assignedGroupsBlock := fmt.Sprintf(`
		assigned_groups {
			%s

			%s
		}
		`,
		clusterAssignedGroup,
		clusterGroupAssignedGroup)

	tfConfigBuilder := &ResourceTFConfigBuilder{
		RequiredResources:   strings.Trim(requiredResources, " "),
		AssignedGroupsBlock: strings.Trim(assignedGroupsBlock, " "),
	}

	return tfConfigBuilder
}

// TODO: Figure out how to build credentials for all 3 resource permutations

func (builder *ResourceTFConfigBuilder) GetTMCManagedTargetLocationConfig(tmcManageCredentials string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
		  name          = "%s"

		  spec {
			target_provider = "AWS"

			credential = {
			  name = "%s"
			}

			%s
		  }
		}
		`,
		builder.RequiredResources,
		targetlocationres.ResourceName,
		TmcManagedResourceName,
		TargetLocationTMCManagedName,
		tmcManageCredentials,
		builder.AssignedGroupsBlock,
	)
}

func (builder *ResourceTFConfigBuilder) GetAWSSelfManagedTargetLocationConfig() string {
	return fmt.Sprintf(`
		resource "tanzu-mission-control_credential" "aws_self_provisioned" {
		  name = "aws-self-provisioned-test"

		  meta {
			description = "Minio storage"
			labels = {
			  "key1" : "value1",
			}
		  }

		  spec {
			capability = "DATA_PROTECTION"
			provider   = "GENERIC_S3"
			data {
			  key_value {
				data = {
				  "aws_secret_key_id" = "abcd"
				  "aws_secret_access_key" = "abcd"
				}
			  }
			}
		  }
		}

		%s

		resource "%s" "%s" {
		  name          = "%s"

		  spec {
			target_provider = "AWS"

			config {
			  aws {
				s3_force_path_style = false
				s3_bucket_url       = "https://minio.vrabbi.cloud"
				s3_public_url       = "https://minio.vrabbi.cloud"
			  }
			}

			bucket = "test-backups"
			region = "us-east-1"
			credential = {
			  name = tanzu-mission-control_credential.aws_self_provisioned.name
			}

			%s

			ca_cert = <<EOF
		-----BEGIN CERTIFICATE-----
		MIIDHzCCAgegAwIBAgIRANzEvNJ7NUMGlLRiJ+yPUVowDQYJKoZIhvcNAQELBQAw
		KTEnMCUGA1UEAxMedGFwLWluZ3Jlc3Mtc2VsZnNpZ25lZC1yb290LWNhMB4XDTIz
		MDMyMTE2MDUyN1oXDTIzMDYxOTE2MDUyN1owKTEnMCUGA1UEAxMedGFwLWluZ3Jl
		c3Mtc2VsZnNpZ25lZC1yb290LWNhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
		CgKCAQEA1lD/uulQCUqgpYY0KKHPMMmvO/oZhLAWuf1XPgI7HeafDXa4bAXPgWyU
		93bkMm0Ir1Owe5+6/GS/rbXT2LwuluIpWfbEUAFREYlZdNmr2bsfmGAJytL45Euz
		W00KAagp3Hxac4xoGu+spRT0KiWspoPcnNktrHSqqT+ApxojNlVVlDC9Bcd+hMdL
		jzpnngS3pOjw8c5bOfhYMr1Pos+b2fyqohpPlqpO2ZXkusiFO6X+Y6FNXxlNy+zB
		SXi+n9CBgJQ20/bvDOHizz31N7JYfGR9QDlPJAlcMNXXMi7kF24MI/938QLOb5Bp
		m/s2Tmlh1Q/ofPr+gMTuiFefyS1VPQIDAQABo0IwQDAOBgNVHQ8BAf8EBAMCAqQw
		DwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUjKyP5mBeHyXE4+b6mVKwD+1DgnYw
		DQYJKoZIhvcNAQELBQADggEBAEBiu9oxbdl27W9rECyyY5L80fS9Kbov+Zo4O6UY
		Mj9eb7a15OUMs5Xy2bJQ94FYpvYoYM4Bx5TPslu25yDWUSQ+ExZBr7+DjhLnBofe
		mV0jzq4Wk2j3vdhs02BztsL8XfwB8csppr21pZzHMZD29wANVnoHoDrNVjT+v9Sa
		vvxeSFFuq0w8mTXX/gM9WbRyo9XSDpTqFbur8lWoCO1AK2Bbk027xF7iDB41xYto
		BEt6yd/ULPf2X1yZU7YBjf2HSCWWpBApE9RX4qk4uD985Y7OqnWYLf9epPQx7r3E
		KMKKMqDkaQn5+um5glbZJXluXDkGO8jEi4bZDsL0irq5uU8=
		-----END CERTIFICATE-----
		EOF
		  }
		}
		`,
		builder.RequiredResources,
		targetlocationres.ResourceName,
		AwsSelfManagedResourceName,
		TargetLocationAWSSelfManagedName,
		builder.AssignedGroupsBlock)
}

func (builder *ResourceTFConfigBuilder) GetAzureSelfManagedTargetLocationConfig(azureCredentialsName string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
		  name          = "%s"

		  spec {
			target_provider = "AZURE"

			config {
			  azure {
				resource_group  = "demo"
				storage_account = "demo"
				subscription_id = "123e4567-e89b-12d3-a456-426614174000"
			  }
			}

			bucket = "test-backups"

			credential = {
			  name = "%s"
			}

			%s
		  }
		}
		`,
		builder.RequiredResources,
		targetlocationres.ResourceName,
		AzureSelfManagedResourceName,
		TargetLocationAzureSelfManagedName,
		azureCredentialsName,
		builder.AssignedGroupsBlock)
}
