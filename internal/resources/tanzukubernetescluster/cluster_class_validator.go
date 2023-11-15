/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"

	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_schema_validator"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

type ClusterClassValidator struct {
	WorkerClasses      []string
	OpenAPIV3Validator *openapiv3.OpenAPIV3SchemaValidator
}

func NewClusterClassValidator(spec *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec) *ClusterClassValidator {
	validator := &ClusterClassValidator{
		OpenAPIV3Validator: &openapiv3.OpenAPIV3SchemaValidator{
			Schema: BuildClusterClassMap(spec),
		},
		WorkerClasses: spec.WorkersClasses,
	}

	return validator
}

func BuildClusterClassMap(spec *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec) map[string]interface{} {
	openAPIV3Schema := make(map[string]interface{})

	for _, v := range spec.Variables {
		decodedTemplate, _ := base64.StdEncoding.DecodeString(v.Schema.Template.Raw.String())
		templateJSON := make(map[string]interface{})
		_ = json.Unmarshal(decodedTemplate, &templateJSON)
		templateSchema := templateJSON["openAPIV3Schema"].(map[string]interface{})

		_, defaultExist := templateSchema[string(openapiv3.DefaultKey)]
		_, requiredExist := templateSchema[string(openapiv3.RequiredKey)]

		if !requiredExist && !defaultExist && v.Required {
			templateSchema[string(openapiv3.RequiredKey)] = true
		}

		openAPIV3Schema[v.Name] = templateSchema
	}

	return openAPIV3Schema
}

func (validator *ClusterClassValidator) ValidateClusterVariables(clusterVariables string, checkRequired bool) (errs []error) {
	errs = make([]error, 0)
	clusterVariablesJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(clusterVariables), &clusterVariablesJSON)

	if checkRequired {
		errs = append(errs, validator.OpenAPIV3Validator.ValidateRequiredFields(clusterVariablesJSON)...)
	}

	errs = append(errs, validator.OpenAPIV3Validator.ValidateFormat(clusterVariablesJSON)...)

	return errs
}

func (validator *ClusterClassValidator) ValidateNodePools(nodePools []interface{}) (errs []error) {
	errs = make([]error, 0)

	for _, np := range nodePools {
		npName := np.(map[string]interface{})[NameKey].(string)
		npSpec := np.(map[string]interface{})[SpecKey].([]interface{})[0].(map[string]interface{})
		npWorkerClass := npSpec[WorkerClassKey].(string)
		npOverrides := npSpec[OverridesKey].(string)
		npWorkerClassFound := false

		for _, wc := range validator.WorkerClasses {
			if wc == npWorkerClass {
				npWorkerClassFound = true
				break
			}
		}

		if !npWorkerClassFound {
			errs = append(errs, errors.Errorf("Worker class for node pool '%s' is invalid. Valid Worker Classes: %s, Worker Class Provided: %s", npName, validator.WorkerClasses, npWorkerClass))
		}

		errs = append(errs, validator.ValidateClusterVariables(npOverrides, false)...)
	}

	return errs
}
