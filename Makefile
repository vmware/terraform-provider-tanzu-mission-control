# Copyright © 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: MPL-2.0

default: build

# build binary
build: fmt
	go build -o bin/terraform-provider-tmc ./cmd/main.go

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Linter
lint:
	golangci-lint run -c ./.golangci.yaml ./cmd/... ./internal/...