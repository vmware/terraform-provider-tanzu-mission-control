/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/integration"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

const (
	enableNamespaceExclusionsSpecKey = "enableNamespaceExclusions"
	namespaceExclusionsSpecKey       = "namespaceExclusions"
)

func TestGenerate(t *testing.T) {
	for expect, given := range map[string]*integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{
		"Cg==": nil,
		"cmVzb3VyY2UgInRhbnp1LW1pc3Npb24tY29udHJvbF9pbnRlZ3JhdGlvbiIgImRlZmF1bHQiIHsKfQo=": {
			FullName: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{},
		},
		"cmVzb3VyY2UgInRhbnp1LW1pc3Npb24tY29udHJvbF9pbnRlZ3JhdGlvbiIgImRlZmF1bHQiIHsKICBpbnRlZ3JhdGlvbl9uYW1lICAgICAgICA9ICJ0YW56dS1zZXJ2aWNlLW1lc2giCiAgY2x1c3Rlcl9uYW1lICAgICAgICAgICAgPSAidGVzdC1jbHVzdGVyLW5hbWUiCiAgcHJvdmlzaW9uZXJfbmFtZSAgICAgICAgPSAiYXR0YWNoZWQiCiAgbWFuYWdlbWVudF9jbHVzdGVyX25hbWUgPSAiYXR0YWNoZWQiCiAgbWV0YSB7CiAgICBkZXNjcmlwdGlvbiA9ICJtZXRhLCBtZXRhLCBtZXRhLCBtb2RlbCIKICAgIGxhYmVscyA9IHsKICAgICAgImtleSI6ICJ2YWx1ZSIKICAgIH0KICB9CiAgc3BlYyB7IAogICAgY29uZmlndXJhdGlvbnMgPSAie1wiZW5hYmxlTmFtZXNwYWNlRXhjbHVzaW9uc1wiOnRydWUsXCJuYW1lc3BhY2VFeGNsdXNpb25zXCI6W119IgogIH0KfQo=": {
			FullName: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName{
				ClusterName:           "test-cluster-name",
				ManagementClusterName: "attached",
				Name:                  "tanzu-service-mesh",
				ProvisionerName:       "attached",
			},
			Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
				Description: "meta, meta, meta, model",
				Labels:      map[string]string{"key": "value"},
			},
			Spec: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationSpec{
				Configurations: map[string]interface{}{
					enableNamespaceExclusionsSpecKey: true,
					namespaceExclusionsSpecKey:       []map[string]interface{}{},
				},
			},
		},
	} {
		expectStr := decode(expect)

		got, err := generateResourceManifest(given)
		if err != nil {
			t.Fatalf("got error: %v\nexpected %q", err, expectStr)
		}

		encodedGot := encode(got)

		if encodedGot != expect {
			t.Errorf("expected: %q\n%s\ngot: %q\n%s", expect, expectStr, encodedGot, got)
		}

		t.Logf("got:\n%s", got)
	}
}

func generateResourceManifest(model *integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration) (string, error) {
	var dst strings.Builder

	tmpl := template.Must(template.New("manifest").
		Funcs(map[string]any{
			"toJSON": toJSON,
		}).
		Parse(integrationResourceManifestTemplate),
	)

	if err := tmpl.Execute(&dst, model); err != nil {
		return "", err
	}

	return dst.String(), nil
}

func generateID(name *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) string {
	if name == nil {
		return ""
	}

	return fmt.Sprintf(
		"ID:%s:%s:%s:%s:%s",
		name.ManagementClusterName,
		name.ProvisionerName,
		name.ClusterName,
		name.Name,
		name.OrgID)
}

func initTestProvider(t *testing.T) *schema.Provider {
	provider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		DataSourcesMap: map[string]*schema.Resource{
			ResourceName: DataSourceIntegration(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := provider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return provider
}

func configureTestProvider(client *testClient) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		return client, diags
	}
}

func newTestClient() *testClient {
	return &testClient{
		resources: map[string]*integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{},
	}
}

type testClient struct {
	resources map[string]*integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration
}

func (t *testClient) ManageV1alpha1ClusterIntegrationResourceServiceCreate(request *integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse, error) {
	id := generateID(request.Integration.FullName)
	if _, ok := t.resources[id]; ok {
		return nil, fmt.Errorf("conflict, %q already created", id)
	}

	resource := &integration.VmwareTanzuManageV1alpha1ClusterIntegrationIntegration{
		FullName: request.Integration.FullName,
		Meta:     &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{UID: id},
		Spec:     request.Integration.Spec,
		Status: &integration.VmwareTanzuManageV1alpha1ClusterIntegrationStatus{
			ClusterViewURL: "",
			Version:        "",
		},
	}

	t.resources[id] = resource

	return &integration.VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse{Integration: resource}, nil
}

func (t *testClient) ManageV1alpha1ClusterIntegrationResourceServiceRead(name *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) (*integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse, error) {
	id := generateID(name)

	r, ok := t.resources[id]
	if !ok {
		return nil, fmt.Errorf("%q not found (have: %+v)", id, t.resources)
	}

	return &integration.VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse{Integration: r}, nil
}

func (t *testClient) ManageV1alpha1ClusterIntegrationResourceServiceDelete(name *integration.VmwareTanzuManageV1alpha1ClusterIntegrationFullName) error {
	id := generateID(name)
	if _, ok := t.resources[id]; !ok {
		return fmt.Errorf("cannot delete %q: not found", id)
	}

	delete(t.resources, id)

	return nil
}

func must(s string, err error) string {
	if err != nil {
		return err.Error()
	}

	return s
}

func encode(in string) string { return base64.StdEncoding.EncodeToString([]byte(in)) }

func decode(in string) string {
	v, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return string(v)
}

const (
	integrationResourceManifestTemplate = `{{ with . }}{{ $fn := .FullName -}}
resource "tanzu_mission_control_integration" "default" {
  {{- with $fn }}{{ with .Name }}
  integration_name        = "{{ . }}"{{ end }}{{ with .ClusterName }}
  cluster_name            = "{{ . }}"{{ end }}{{ with .ProvisionerName }}
  provisioner_name        = "{{ . }}"{{ end }}{{ with .ManagementClusterName }}
  management_cluster_name = "{{ . }}"{{ end }}{{ end }}{{ with .Meta }}
  meta { {{- with .Description }}
    description = "{{ . }}"{{ end }}{{ with .Labels }}
    labels = { {{- range $k, $v := . }}
      "{{ $k }}": "{{ $v }}"{{ end }}
    } {{- end }}
  } {{- end }}{{ with .Spec }}
  spec { {{ with .Configurations }}
    configurations = {{ toJSON . | printf "%q" }}{{ end }}
  } {{- end }}
} {{- end }}
`
)
