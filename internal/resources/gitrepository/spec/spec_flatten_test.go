// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"testing"

	"github.com/stretchr/testify/require"

	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
)

const (
	test123prereleasebuild                       = "1.2.3-prerelease+build"
	testCeb15bcd23d4bb76751064534e3c8d2e09104da6 = "ceb15bcd23d4bb76751064534e3c8d2e09104da6"
	testMain                                     = "main"
	testNameOfTheSecret                          = "name-of-the-secret"
	testURL                                      = "https://github.com/dineshtripathi30/tmc-cd"
	testV100                                     = "v1.0.0"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster git repository spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster git repository spec",
			input: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
				URL:               testURL,
				SecretRef:         testNameOfTheSecret,
				Interval:          "5m",
				GitImplementation: gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT.Pointer(),
				Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
					Branch: testMain,
					Tag:    testV100,
					Semver: test123prereleasebuild,
					Commit: testCeb15bcd23d4bb76751064534e3c8d2e09104da6,
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					gitImplementationKey: "GO_GIT",
					intervalKey:          "5m",
					refKey: []interface{}{
						map[string]interface{}{
							branchKey: testMain,
							commitKey: testCeb15bcd23d4bb76751064534e3c8d2e09104da6,
							semverKey: test123prereleasebuild,
							tagKey:    testV100,
						},
					},
					secretRefKey: testNameOfTheSecret,
					URLKey:       testURL,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestFlattenSpecForClusterGroupScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec
		expected    []interface{}
	}{
		{
			description: "check for nil cluster group git repository spec",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete cluster group git repository spec",
			input: &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositorySpec{
				AtomicSpec: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{
					URL:               testURL,
					SecretRef:         testNameOfTheSecret,
					Interval:          "5m",
					GitImplementation: gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT.Pointer(),
					Ref: &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryReference{
						Branch: testMain,
						Tag:    testV100,
						Semver: test123prereleasebuild,
						Commit: testCeb15bcd23d4bb76751064534e3c8d2e09104da6,
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					gitImplementationKey: "GO_GIT",
					intervalKey:          "5m",
					refKey: []interface{}{
						map[string]interface{}{
							branchKey: testMain,
							commitKey: testCeb15bcd23d4bb76751064534e3c8d2e09104da6,
							semverKey: test123prereleasebuild,
							tagKey:    testV100,
						},
					},
					secretRefKey: testNameOfTheSecret,
					URLKey:       testURL,
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterGroupScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
