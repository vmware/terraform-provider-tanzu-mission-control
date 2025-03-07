/*
  NOTE: Creation of custom policy depends on cluster group and custom policy template.
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group.
resource "tanzu-mission-control_cluster_group" "cluster_group" {
  name = "tf-demo-cluster-group"
}

# Create custom policy template.
resource "tanzu-mission-control_custom_policy_template" "sample_template" {
  name = "tf-custom-template-test"

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

    template_manifest = <<YAML
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: tf-custom-template-test
  annotations:
    description: Requires Pods to have readiness and/or liveness probes.
spec:
  crd:
    spec:
      names:
        kind: tf-custom-template-test
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
  }
}


# Cluster group scoped custom template assigned Custom Policy
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_custom_template_assigned_custom_policy" {
  name = "tf-custom-template-policy-test"

  scope {
    cluster_group {
      cluster_group = tanzu-mission-control_cluster_group.cluster_group.name
    }
  }

  spec {
    input {
      custom {
        template_name = tanzu-mission-control_custom_policy_template.sample_template.name
        audit = false

        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Deployment"
          ]
        }

        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "StatefulSet",
          ]
        }
      }
    }
  }
}
