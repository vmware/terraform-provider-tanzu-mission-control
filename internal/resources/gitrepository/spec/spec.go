/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Repository.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			URLKey: {
				Type:         schema.TypeString,
				Description:  "URL of the git repository. Repository URL should begin with http, https, or ssh",
				Required:     true,
				ValidateFunc: validation.IsURLWithScheme([]string{"http", "https", "ssh"}),
			},
			secretRefKey: {
				Type:        schema.TypeString,
				Description: "Reference to the secret. Repository credential.",
				Optional:    true,
				Default:     "",
			},
			intervalKey: {
				Type:        schema.TypeString,
				Description: "Interval at which to check gitrepository for updates. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the repository to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.",
				Optional:    true,
				Default:     "5m",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					durationInScript, err := time.ParseDuration(old)
					if err != nil {
						return false
					}

					durationInState, err := time.ParseDuration(new)
					if err != nil {
						return false
					}

					return durationInScript.Seconds() == durationInState.Seconds()
				},
			},
			gitImplementationKey: {
				Type:        schema.TypeString,
				Description: "GitImplementation specifies which client library implementation to use. go-git is the default git implementation.",
				Optional:    true,
				Default:     fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT),
				ValidateFunc: validation.StringInSlice([]string{
					fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT),
					fmt.Sprint(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationLIBGIT2),
				}, false),
			},
			refKey: refSchema,
		},
	},
}

var refSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Reference specifies git reference to resolve.",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			branchKey: {
				Type:        schema.TypeString,
				Description: "Branch from git to checkout. When branch is given, then that branch from the git repository will be checked out. If the given branch doesn’t exist in the git repository, then adding the git repository will fail. If no branch is given, the `master` branch will be used.",
				Optional:    true,
				Default:     "",
			},
			tagKey: {
				Type:        schema.TypeString,
				Description: "Tag from git to checkout. Takes precedence over branch. When a tag is given, that tag from the git repository will be checked out. If the given tag doesn’t exist in the git repository, then adding the git repository will fail. If both tag and branch are given, tag overrides branch and the branch value will be ignored.",
				Optional:    true,
				Default:     "",
			},
			semverKey: {
				Type:        schema.TypeString,
				Description: "SemVer expression to checkout from git tags. Takes precedence over tag. When semver is given, then the latest tag matching that semver will be checked out from the git repository. If no tag in the git repository matches semver, then adding the git repository will fail. If semver is given, tag and branch will be ignored if they are populated.",
				Optional:    true,
				Default:     "",
			},
			commitKey: {
				Type:        schema.TypeString,
				Description: "Commit SHA to checkout. Takes precedence over all other reference fields. When git_implementation is `GO_GIT`, this can be combined with branch to shallow clone branch in which the commit is expected to exist.",
				Optional:    true,
				Default:     "",
			},
		},
	},
}

func expandRef(data []interface{}) (ref *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference) {
	if len(data) == 0 || data[0] == nil {
		return ref
	}

	refData, _ := data[0].(map[string]interface{})

	ref = &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{}

	if branchValue, ok := refData[branchKey]; ok {
		helper.SetPrimitiveValue(branchValue, &ref.Branch, branchKey)
	}

	if tagValue, ok := refData[tagKey]; ok {
		helper.SetPrimitiveValue(tagValue, &ref.Tag, tagKey)
	}

	if semverValue, ok := refData[semverKey]; ok {
		helper.SetPrimitiveValue(semverValue, &ref.Semver, semverKey)
	}

	if commitValue, ok := refData[commitKey]; ok {
		helper.SetPrimitiveValue(commitValue, &ref.Commit, commitKey)
	}

	return ref
}

func flattenRef(ref *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference) (data []interface{}) {
	if ref == nil {
		return data
	}

	flattenRefData := make(map[string]interface{})

	flattenRefData[branchKey] = ref.Branch
	flattenRefData[tagKey] = ref.Tag
	flattenRefData[semverKey] = ref.Semver
	flattenRefData[commitKey] = ref.Commit

	return []interface{}{flattenRefData}
}

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, URLKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, secretRefKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, intervalKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, gitImplementationKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, refKey)):
		updateRequired = true
	}

	return updateRequired
}
