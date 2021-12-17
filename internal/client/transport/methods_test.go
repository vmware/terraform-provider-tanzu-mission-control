/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package transport

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	workspacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

type Invoke struct {
	HTTPMethodType string
	URL            string
	Request        Request
	Response       Response
}

func TestConcurrentAccessOfInvokeAction(t *testing.T) {
	var input Invoke

	var waitGroup sync.WaitGroup

	input.HTTPMethodType = "POST"
	input.URL = "xyz.com"
	input.Request = &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceRequest{
		Workspace: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceWorkspace{
			FullName: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: "tf-workspace-test",
			},
		},
	}
	input.Response = &workspacemodel.VmwareTanzuManageV1alphaWorkspaceResponse{
		Workspace: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceWorkspace{
			FullName: &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{
				Name: "tf-workspace-test",
			},
		},
	}

	c := NewClient()
	c.AddHeaders(map[string][]string{
		"header1": {"one", "two"},
		"header2": {"three", "four"},
	})

	for i := 1; i <= 100; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			actual := c.invokeAction(input.HTTPMethodType, input.URL, input.Request, input.Response)
			require.Error(t, actual)
		}()
	}
	waitGroup.Wait()
}
