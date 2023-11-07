/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatetests

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	custompolicytemplateres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/custompolicytemplate"
)

const (
	CustomPolicyTemplateResourceName = "test_custom_policy_template"
)

var (
	CustomPolicyTemplateResourceFullName = fmt.Sprintf("%s.%s", custompolicytemplateres.ResourceName, CustomPolicyTemplateResourceName)
	CustomPolicyTemplateName             = acctest.RandomWithPrefix("acc-test-custom-policy-template")
)

type ResourceTFConfigBuilder struct {
	TemplateManifest string
}

func InitResourceTFConfigBuilder() *ResourceTFConfigBuilder {
	firstManifestPart := fmt.Sprintf(
		`<<YAML
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: %s
  annotations:
    description: Requires Pods to have readiness and/or liveness probes.
spec:
  crd:
    spec:
      names:
        kind: %[1]s
`,
		CustomPolicyTemplateName)

	secondManifestPart := `
      validation:
        openAPIV3Schema:
          properties:
            probes:
              type: array
              items:
                type: string
            probeTypes:
              type: array
              items:
                type: string
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8srequiredprobes
        probe_type_set = probe_types {
          probe_types := {type | type := input.parameters.probeTypes[_]}
        }
        violation[{"msg": msg}] {
          container := input.review.object.spec.containers[_]
          probe := input.parameters.probes[_]
          probe_is_missing(container, probe)
          msg := get_violation_message(container, input.review, probe)
        }
        probe_is_missing(ctr, probe) = true {
          not ctr[probe]
        }
        probe_is_missing(ctr, probe) = true {
          probe_field_empty(ctr, probe)
        }
        probe_field_empty(ctr, probe) = true {
          probe_fields := {field | ctr[probe][field]}
          diff_fields := probe_type_set - probe_fields
          count(diff_fields) == count(probe_type_set)
        }
        get_violation_message(container, review, probe) = msg {
          msg := sprintf("Container <%v> in your <%v> <%v> has no <%v>", [container.name, review.kind.kind, review.object.metadata.name, probe])
        }
YAML
`

	tfConfigBuilder := &ResourceTFConfigBuilder{
		TemplateManifest: fmt.Sprintf("%s\n%s", firstManifestPart, secondManifestPart),
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetFullCustomPolicyTemplateConfig() string {
	return fmt.Sprintf(`	
		resource "%s" "%s" {
		  name = "%s"
		
		  spec {
			object_type   = "ConstraintTemplate"
			template_type = "OPAGatekeeper"
		
			data_inventory {
			  kind    = "ConfigMap"
			  group   = "admissionregistration.k8s.io"
			  version = "v1"
			}
		
			data_inventory {
			  kind    = "Deployment"
			  group   = "extensions"
			  version = "v1"
			}
		
			template_manifest = %s
		  }
		}
		`,
		custompolicytemplateres.ResourceName,
		CustomPolicyTemplateResourceName,
		CustomPolicyTemplateName,
		builder.TemplateManifest,
	)
}

func (builder *ResourceTFConfigBuilder) GetSlimCustomPolicyTemplateConfig() string {
	return fmt.Sprintf(`	
		resource "%s" "%s" {
		  name = "%s"
		
		  spec {
			object_type   = "ConstraintTemplate"
			template_type = "OPAGatekeeper"

			template_manifest = %s
		  }
		}
		`,
		custompolicytemplateres.ResourceName,
		CustomPolicyTemplateResourceName,
		CustomPolicyTemplateName,
		builder.TemplateManifest,
	)
}
