#Copyright © 2021 VMware, Inc. All Rights Reserved.
#SPDX-License-Identifier: MPL-2.0

FROM 498533941640.dkr.ecr.us-west-2.amazonaws.com/base-images/golang-build-1.16 as build
WORKDIR /go/src/gitlab.eng.vmware/olympus/terraform-provider-tanzu

COPY .golangci.yaml go.mod go.sum ./
COPY cmd cmd
COPY internal internal

VOLUME /test-results

ENTRYPOINT  go test -coverprofile=/test-results/coverage.out ./internal/... ./cmd/... && \
            go tool cover -func=/test-results/coverage.out > /test-results/function-wise-coverage.txt && \
            cat /test-results/function-wise-coverage.txt | awk 'END{print $NF}' > /test-results/totalcoverage.txt && \
            go tool cover -html=/test-results/coverage.out -o=/test-results/coverage.html && \
            GO111MODULE=on CGO_ENABLED=0 golangci-lint run -c ./.golangci.yaml ./cmd/... ./internal/...
