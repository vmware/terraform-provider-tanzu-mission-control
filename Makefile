# Copyright © 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: MPL-2.0

ifeq ($(VERSION_TAG),)
	VERSION_TAG := 1.0.0
endif

ifeq ($(GOARCH),)
	GOARCH := $(shell go env GOARCH)
endif

ifeq ($(GOOS),)
	GOOS := $(shell go env GOOS)
endif

ifeq ($(TEST_PKGS),)
	TEST_PKGS := ./...
endif

ifeq ($(TEST_FLAGS),)
	TEST_FLAGS := -cover
endif

ifeq ($(BUILD_TAGS),)
	BUILD_TAGS := 'akscluster cluster clustergroup credential ekscluster gitrepository iampolicy kustomization namespace custompolicy imagepolicy networkpolicy quotapolicy securitypolicy sourcesecret workspace tanzupackage tanzupackages packagerepository packageinstall clustersecret integration mutationpolicy managementclusterregistration'
endif

.PHONY: build clean-up test gofmt vet lint acc-test website-lint website-lint-fix

default: build

build:
	go build -o bin/terraform-provider-tanzu-mission-control_$(VERSION_TAG)
	mkdir -p ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/$(VERSION_TAG:v%=%)/$(GOOS)_$(GOARCH)/
	cp bin/terraform-provider-tanzu-mission-control_$(VERSION_TAG) ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/$(VERSION_TAG:v%=%)/$(GOOS)_$(GOARCH)/

clean-up:
	rm -rf ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/*

test: gofmt terraform-fmt vet lint
	go mod tidy
	go test $(TEST_PKGS) $(TEST_FLAGS)

# Run go fmt against code
gofmt:
	@echo Checking code is gofmted
	go fmt $(TEST_PKGS)

terraform-fmt:
	terraform fmt -recursive

# Run go vet against code
vet:
	go vet $(TEST_PKGS)

# Linter
lint: gofmt
	golangci-lint run -c ./.golangci.yml ./internal/... .

acc-test:
	go test $(TEST_PKGS) -tags $(BUILD_TAGS)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/ || (echo; \
	    echo "Unexpected misspelling found in website files."; \
	    echo "To automatically fix the misspelling, run 'make website-lint-fix' and commit the changes."; \
	    exit 1)
	@terrafmt diff ./docs --check --pattern '*.markdown' --quiet || (echo; \
	    echo "Unexpected differences in website HCL formatting."; \
	    echo "To see the full differences, run: terrafmt diff ./website --pattern '*.markdown'"; \
	    echo "To automatically fix the formatting, run 'make website-lint-fix' and commit the changes."; \
	    exit 1)

website-lint-fix:
	@echo "==> Applying automatic website linter fixes..."
	@misspell -w -source=text website/
	@terrafmt fmt ./docs --pattern '*.markdown'
