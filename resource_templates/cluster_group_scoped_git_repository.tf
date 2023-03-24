// Tanzu Mission Control GitRepository
// Operations supported : Read, Create, Update & Delete

# Read Tanzu Mission Control cluster group scoped gitrepository : fetch cluster group scoped gitrepository details
data "tanzu-mission-control_git_repository" "test_repository" {
  name = "<existing-name>" // Required

  namespace_name = "<existing-namespace_name>" // Required

  org_id = "<existing-org-id>" // Default: taken from the authentication context

  scope {
    cluster_group {
      name = "<existing-cluster_group_name>" // Required
    }
  }
}

# Create Tanzu Mission Control cluster group scoped gitrepository entry
resource "tanzu-mission-control_git_repository" "test_repository" {
  name = "test"

  namespace_name = "test"

  org_id = "test" // optional

  scope {
    cluster_group {
      name = "test"
    }
  }

  spec {
    url = "https://github.com/dineshtripathi30/tmc-cd"

    secret_ref = "name-of-the-secret" // can be referenced in Tf - to explicitly call delete on this resource if the dependency is deleted

    interval = "5m" // can be "5s" etc.

    git_implementation = "GO_GIT" // enum - can be LIB_GIT2

    ref { // specifies git reference to resolve and checkout
      branch = "main"

      tag = "v1.0.0"

      semver = "1.2.3-prerelease+build"

      commit = "ceb15bcd23d4bb76751064534e3c8d2e09104da6"
    }
  }
}
