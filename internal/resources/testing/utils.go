/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type ignoreFunc func(string) bool

func TagsIgnoreFunc(diffItem string) bool {
	return strings.Contains(diffItem, "map[tags]") && strings.Contains(diffItem, "tmc.cloud.vmware.com")
}

func MetaUIDIgnoreFunc(diffItem string) bool {
	return strings.Contains(diffItem, "map[uid]")
}

func BodyInspectingResponder(t *testing.T, expectedContent interface{}, successResponse int, successResponseBody interface{}, ignoreFuncs ...ignoreFunc) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		successFunc := func() (*http.Response, error) {
			return httpmock.NewJsonResponse(successResponse, successResponseBody)
		}

		if expectedContent == nil {
			return successFunc()
		}

		// Compare to expected content.
		expectedBytes, err := json.Marshal(expectedContent)
		if err != nil {
			t.Fail()
			return nil, err
		}

		if r.Body == nil {
			t.Fail()
			return nil, fmt.Errorf("expected body on request")
		}

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fail()
			return nil, err
		}

		// Map of map of strings for comparing subnet equality
		subnetMap := make(map[string]map[string][]string, 0)

		var bodyInterface map[string]interface{}

		err = json.Unmarshal(bodyBytes, &bodyInterface)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal body")
		}

		var expectedInterface map[string]interface{}

		err = json.Unmarshal(expectedBytes, &expectedInterface)
		if err != nil {
			return nil, err
		}

		diff := deep.Equal(bodyInterface, expectedInterface)
		if diff == nil {
			return successFunc()
		}
		// special check for subnets
		// First, populate all the diffs pertaining to subnets into maps
		for _, diffItem := range diff {
			if slices.IndexFunc(ignoreFuncs, func(f ignoreFunc) bool { return f(diffItem) }) != -1 {
				continue
			}

			if !strings.Contains(diffItem, "map[subnetIds]") || !strings.Contains(diffItem, ".slice") {
				t.Fail()
				return nil, errors.Errorf("diff identified outside of subnet order or passed ignoreDiff: %s", diffItem)
			}

			segments := strings.Split(diffItem, ":")
			key := strings.Split(segments[0], ".slice")[0]

			// Create map if not present
			if subnetMap[key] == nil {
				subnetMap[key] = make(map[string][]string, 0)
			}

			// Add vals to map
			vals := strings.Split(segments[1], "!=")
			subnetMap[key]["left"] = append(subnetMap[key]["left"], strings.TrimSpace(vals[0]))
			subnetMap[key]["right"] = append(subnetMap[key]["right"], strings.TrimSpace(vals[1]))
		}

		// Then, sort slices and compare
		for _, set := range subnetMap {
			left := set["left"]
			right := set["right"]

			// sort
			sort.Strings(left)
			sort.Strings(right)

			subnetDiff := deep.Equal(left, right)
			if subnetDiff != nil {
				t.Fail()
				return nil, errors.New("subnets did not match")
			}
		}

		return successFunc()
	}
}
