// Create/ Delete/ Update Tanzu Mission Control organization scoped custom namespace quota policy entry
resource "tanzu-mission-control_namespace_quota_policy" "organization_scoped_custom_quota_policy" {
  name = "<namespace-quota-policy-name>"

  scope {
    organization {
      organization = "<organization-id>" // Required
    }
  }

  spec {
    input {
      custom {
        limits_cpu               = "<CPU-limits>"
        limits_memory            = "<memory-limits>Mi"
        persistent_volume_claims = 0 // Default: 0
        persistent_volume_claims_per_class = {
          "<class-name-1>" : 0 // Default: 0
          "<class-name-2>" : 0 // Default: 0
        }
        requests_cpu     = "<CPU-requests>"
        requests_memory  = "<memory-requests>Mi"
        requests_storage = "<storage-requests>G"
        requests_storage_per_class = {
          "<class-name-1>" : "<count>G"
          "<class-name-2>" : "<count>G"
        }
        resource_counts = {
          "<object-1>" : 0 // Default: 0
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values = [
          "<label-selector-requirement-value-1>",
          "<label-selector-requirement-value-2>"
        ]
      }
      match_expressions {
        key      = "<label-selector-requirement-key-2>"
        operator = "<label-selector-requirement-operator>"
        values   = []
      }
    }
  }
}
