/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmrepository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	pakageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

const (
	https                = "https:/"
	clAPIVersionAndGroup = "v1alpha1/clusters"
	apiSubGroup          = "namespaces"
	apiKind              = "fluxcd/helm/repositories"
)

func bodyInspectingResponder(t *testing.T, expectedContent interface{}, successResponse int, successResponseBody interface{}) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		successFunc := func() (*http.Response, error) {
			return httpmock.NewJsonResponse(successResponse, successResponseBody)
		}

		if expectedContent == nil {
			return successFunc()
		}

		// Compare to expected content.
		expectedBytes, err := json.Marshal(expectedContent)
		if err != nil {
			t.Fail()
			return nil, err
		}

		if r.Body == nil {
			t.Fail()
			return nil, fmt.Errorf("expected body on request")
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fail()
			return nil, err
		}

		var bodyInterface map[string]interface{}
		if err = json.Unmarshal(bodyBytes, &bodyInterface); err == nil {
			var expectedInterface map[string]interface{}

			err = json.Unmarshal(expectedBytes, &expectedInterface)
			if err != nil {
				return nil, err
			}

			diff := deep.Equal(bodyInterface, expectedInterface)
			if diff == nil {
				return successFunc()
			}
		} else {
			return nil, err
		}

		return successFunc()
	}
}

func (testConfig *testAcceptanceConfig) setupHTTPMocks(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(httpmock.Deactivate)

	endpoint := os.Getenv("TMC_ENDPOINT")

	OrgID := os.Getenv("ORG_ID")

	reference := objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference{
		Rid: "test_rid",
		UID: "test_uid",
	}
	referenceArray := make([]*objectmetamodel.VmwareTanzuCoreV1alpha1ObjectReference, 0)
	referenceArray = append(referenceArray, &reference)

	// cluster level package resource.

	getResponse := &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryListResponse{
		TotalCount: "1",
		Repositories: []*pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository{
			{
				FullName: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName{
					ClusterName:           testConfig.ScopeHelperResources.Cluster.Name,
					ManagementClusterName: "attached",
					ProvisionerName:       "attacched",
					OrgID:                 OrgID,
					Name:                  "*",
					NamespaceName:         testConfig.Namespace,
				},
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					ParentReferences: referenceArray,
					Description:      "resource with description",
					Labels: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					UID:             "helmrepo1",
					ResourceVersion: "v1",
				},
				Spec: &pakageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySpec{
					URL: "https://charts.bitnami.com/bitnami",
				},
			},
		},
	}

	getPkgEndpoint := (helper.ConstructRequestURL(https, endpoint, clAPIVersionAndGroup, testConfig.ScopeHelperResources.Cluster.Name, apiSubGroup, testConfig.Namespace, apiKind)).String()

	httpmock.RegisterResponder("GET", getPkgEndpoint,
		bodyInspectingResponder(t, nil, 200, getResponse))
}
