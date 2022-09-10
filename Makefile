# Copyright © 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: MPL-2.0

default: build

build:
	go build -o bin/terraform-provider-tanzu-mission-control_v1.0.0
	mkdir -p ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/1.0.0/darwin_amd64/
	cp bin/terraform-provider-tanzu-mission-control_v1.0.0 ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/1.0.0/darwin_amd64/

build-arm:
	go build -o bin/terraform-provider-tanzu-mission-control_v1.0.0
	mkdir -p ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/1.0.0/darwin_arm64/
	cp bin/terraform-provider-tanzu-mission-control_v1.0.0 ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/1.0.0/darwin_arm64/

clean-up:
	rm -rf ~/.terraform.d/plugins/vmware/dev/tanzu-mission-control/*

test: | gofmt vet lint
	go mod tidy
	go test ./internal/... -cover

# Run go fmt against code
gofmt:
	@echo Checking code is gofmted
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Linter
lint: gofmt
	GO111MODULE=on golangci-lint run -c ./.golangci.yml ./internal/... .

set-tf-acc:
	@echo Setting Acceptance test env variable
	export TF_ACC=true;\
	echo $$TF_ACC

acc-test: set-tf-acc test

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/ || (echo; \
	    echo "Unexpected mispelling found in website files."; \
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