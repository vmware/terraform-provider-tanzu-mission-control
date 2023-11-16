package clusterclass

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_schema_validator"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

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

func generateClusterVariablesTemplate(clusterVariables map[string]interface{}) (schemaTemplate map[string]interface{}) {
	schemaTemplate = make(map[string]interface{})

	for k, v := range clusterVariables {
		schemaTemplate[k] = buildOpenAPIV3Template(v.(map[string]interface{}))
	}

	return schemaTemplate
}

func buildOpenAPIV3Template(openAPIV3Schema map[string]interface{}) (templateValue interface{}) {
	switch openAPIV3Schema[string(openapiv3.TypeKey)].(string) {
	case string(openapiv3.ObjectType):
		templateValue = map[string]interface{}{}

		if objSchema, ok := openAPIV3Schema[string(openapiv3.PropertiesKey)]; ok {
			for k, v := range objSchema.(map[string]interface{}) {
				templateValue.(map[string]interface{})[k] = buildOpenAPIV3Template(v.(map[string]interface{}))
			}
		} else {
			templateValue.(map[string]interface{})["custom_key"] = buildOpenAPIV3Template(openAPIV3Schema[string(openapiv3.AdditionalPropertiesKey)].(map[string]interface{}))
		}
	case string(openapiv3.ArrayType):
		templateValue = []interface{}{}
		templateValue = append(templateValue.([]interface{}), buildOpenAPIV3Template(openAPIV3Schema[string(openapiv3.ItemsKey)].(map[string]interface{})))
	case string(openapiv3.StringType):
		templateValue = "String"

		if regexPattern, ok := openAPIV3Schema[string(openapiv3.PatternKey)]; ok {
			templateValue = fmt.Sprintf("%s (regex: %s)", templateValue, regexPattern)
		}

		if minLen, ok := openAPIV3Schema[string(openapiv3.MinLengthKey)]; ok {
			templateValue = fmt.Sprintf("%s (minLen: %v)", templateValue, minLen)
		}

		if maxLen, ok := openAPIV3Schema[string(openapiv3.MaxLengthKey)]; ok {
			templateValue = fmt.Sprintf("%s (maxLen: %v)", templateValue, maxLen)
		}
	case string(openapiv3.BooleanType):
		templateValue = false
	case string(openapiv3.IntegerType):
		templateValue = 1

		if minValue, ok := openAPIV3Schema[string(openapiv3.MinimumKey)]; ok {
			templateValue = minValue
		} else if maxValue, ok := openAPIV3Schema[string(openapiv3.MaximumKey)]; ok {
			templateValue = maxValue
		}
	case string(openapiv3.NumberType):
		templateValue = 0.5
		templateValue = fmt.Sprintf("%.2f", templateValue)
		templateValue, _ = strconv.ParseFloat(templateValue.(string), 64)
	}

	return templateValue
}
