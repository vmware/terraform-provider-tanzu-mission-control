# CHANGELOG

## [v1.4.9](https://github.com/vmware/terraform-provider-tanzu-mission-control/releases/tag/v1.4.9)

> Release Date: 2025-04-01

FEATURES:

- Added support to override variables in the control plane for VKS 3.2.0. [\#488](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/488)

DOCUMENTATION:

- Update links in product documentation. [\#524](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/524)
- Update naming style from imperative to declarative. [\#491](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/491)

BUG FIXES:

- Fixed datasource schema for tanzu kubernetes clusters. [\#479](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/479)

CHORES:

- Bump `codecov/codecov-action` from 5.3.1 to 5.4.0. [/#515](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/515)
- Bump `actions/checkout` from 4.2.1 to 4.2.2. [/#514](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/514)
- Remove `codeql`. [/#512](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/512)
- Bump `goreleaser/goreleaser-action` from 6.1.0 to 6.2.1. [/#510](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/510)
- Bump `golangci/golangci-lint-action` from 6.2.0 to 6.5.0. [/509](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/509)
- Bump `codecov/codecov-action` from 5.1.2 to 5.3.1. [/#502](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/502)
- Bump `actions/stale` from 9.0.0 to 9.1.0. [/#501](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/501)
- Bump `actions/setup-go` from 5.2.0 to 5.3.0. [/#500](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/500)

## [v1.4.8](https://github.com/vmware/terraform-provider-tanzu-mission-control/releases/tag/v1.4.8)

> Release Date: 2024-01-30

BUG FIXES:

- Fix for eks cluster and nodepool tag issue [\#463](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/463)
- Fix incorrect conversion between integer types. [\#467](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/467)

DEPRECATIONS:

- Remove TSM integration support. [\#475](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/475)

CHORES:

- Updated `crazy-max/ghaction-import-gpg` to 6.2.0. [\#459](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/459)
- Updated `goreleaser/goreleaser-action` to 6.1.0 [\#460](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/460)
- Updated `codecov/codecov-action` to 5 [\#461](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/461)
- Updated `go` to 1.23.2 [\#462](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/462)
- Updated `codecov/codecov-action` to 5.0.7 [\#464](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/464)
- Updated `github.com/stretchr/testify` to 1.10.0 [\#466](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/466)
- Updated `codecov/codecov-action` to 5.1.1 [\#473](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/473)
- Updated `golang.org/x/crypto` to 0.31.0 [\#474](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/474)
- Updated `actions/setup-go` to 5.2.0 [\#480](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/480)
- Updated `golang.org/x/net` to 0.33.0 [\#490](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/490)
- Updated `codecov/codecov-action` to 5.1.2 [\#496](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/496)
- Updated `golangci/golangci-lint-action` to 6.2.0 [\#497](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/497)

## [v1.4.7](https://github.com/vmware/terraform-provider-tanzu-mission-control/releases/tag/v1.4.7)

> Release Date: 2024-11-13

FEATURES:

- `data/tanzu-mission-control_tanzu_kubernetes_cluster`: Added Tanzu Kubernetes Cluster data source. [\#444](https://github.com/vmware/terraform-provider-vcf/pull/444)

DOCUMENTATION:

- Updated AKS Cluster guide. [\#326](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/326)

CHORES:

- Added CodeQL Analysis. [\#436](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/436)
- Updated `golang.org/x/net` to 0.23.0. [\#398](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/398)
- Updated `golang.org/x/oauth2` to 0.24.0. [\#438](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/438)
- Updated `sigs.k8s.io/yaml` to 1.4.0. [\#430](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/430)
- Updated `github.com/stretchr/testify` to 1.9.0. [\#432](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/432)
- Updated `github.com/go-openapi/strfmt` to 0.23.0. [\#433](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/433)
- Updated `github.com/go-test/deep` to 1.1.1. [\#439](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/439)
- Updated `github.com/jarcoal/httpmock` to 1.3.1. [\#441](https://github.com/vmware/terraform-provider-tanzu-mission-control/pull/441)
