/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package openapiv3schemavalidator

import (
	"regexp"

	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

type OpenAPIV3Key string
type OpenAPIV3Types string

const (
	RequiredKey             OpenAPIV3Key = "required"
	DefaultKey              OpenAPIV3Key = "default"
	PropertiesKey           OpenAPIV3Key = "properties"
	TypeKey                 OpenAPIV3Key = "type"
	AdditionalPropertiesKey OpenAPIV3Key = "additionalProperties"
	PreserveUnknownFieldKey OpenAPIV3Key = "x-kubernetes-preserve-unknown-fields"
	ItemsKey                OpenAPIV3Key = "items"
	PatternKey              OpenAPIV3Key = "pattern"
	MinLengthKey            OpenAPIV3Key = "minLength"
	MaxLengthKey            OpenAPIV3Key = "maxLength"
	MinimumKey              OpenAPIV3Key = "minimum"
	MaximumKey              OpenAPIV3Key = "maximum"
)

const (
	ObjectType  OpenAPIV3Types = "object"
	NumberType  OpenAPIV3Types = "number"
	IntegerType OpenAPIV3Types = "integer"
	BooleanType OpenAPIV3Types = "boolean"
	ArrayType   OpenAPIV3Types = "array"
	StringType  OpenAPIV3Types = "string"
)

type OpenAPIV3SchemaValidator struct {
	Schema map[string]interface{}
}

func (validator *OpenAPIV3SchemaValidator) ValidateRequiredFields(objectValues map[string]interface{}) (errs []error) {
	errs = make([]error, 0)

	for k, v := range validator.Schema {
		objectValue := objectValues[k]
		errs = append(errs, validateRequiredFields(false, k, objectValue, v.(map[string]interface{}))...)
	}

	return errs
}

func (validator *OpenAPIV3SchemaValidator) ValidateFormat(objectValues map[string]interface{}) (errs []error) {
	errs = make([]error, 0)

	for k, v := range objectValues {
		fieldSchema, fieldExists := validator.Schema[k]

		if !fieldExists {
			errs = append(errs, errors.Errorf("Key '%s' is not expected in cluster class schema.", k))
		} else {
			vErrs := validateSchemaFormat(k, v, fieldSchema.(map[string]interface{}))

			for _, e := range vErrs {
				errs = append(errs, errors.Wrapf(e, "Value validation failed for key '%s'", k))
			}
		}
	}

	return errs
}

func validateRequiredFields(isParentRequired bool, parentKey string, variableValue interface{}, variableSchema map[string]interface{}) (errs []error) {
	errs = make([]error, 0)
	requiredValue := variableSchema[string(RequiredKey)]
	isRequired, isRequiredBool := requiredValue.(bool)
	isValueEmpty := helper.IsEmptyInterface(variableValue)

	if isValueEmpty && isRequired && variableSchema[string(TypeKey)] != string(ObjectType) {
		errs = append(errs, errors.Errorf("Key '%s' is required but not provided.", parentKey))

		return errs
	} else if variableSchema[string(TypeKey)] == string(ObjectType) {
		var (
			variableValueMap map[string]interface{}
		)

		if !isValueEmpty {
			variableValueMap, _ = variableValue.(map[string]interface{})

			if !isRequiredBool && requiredValue != nil {
				for _, requiredField := range requiredValue.([]interface{}) {
					_, requiredExists := variableValueMap[requiredField.(string)]

					if !requiredExists {
						errs = append(errs, errors.Errorf("Key '%s' is required in object '%s' but not provided!", requiredField, parentKey))
					}
				}
			}
		}

		mapFields, mapFieldsExist := variableSchema[string(PropertiesKey)]

		if mapFieldsExist {
			for k, v := range mapFields.(map[string]interface{}) {
				var subKeyValue interface{} = nil

				if variableValueMap != nil {
					subKeyValue = variableValueMap[k]
				}

				_, subKeyValueDefaultExist := v.(map[string]interface{})[string(DefaultKey)]

				if (isRequired || isParentRequired) && subKeyValue == nil && !subKeyValueDefaultExist {
					errs = append(errs, errors.Errorf("Key '%s' is required in object '%s' but not provided!", k, parentKey))
				}

				for _, e := range validateRequiredFields(isRequired || isParentRequired, k, subKeyValue, v.(map[string]interface{})) {
					errs = append(errs, errors.Wrapf(e, "Object '%s' field validation failed", parentKey))
				}
			}
		}
	}

	return errs
}

func validateSchemaFormat(parentKey string, variableValue interface{}, variableSchema map[string]interface{}) (errs []error) {
	varType := variableSchema[string(TypeKey)].(string)

	switch varType {
	case string(ObjectType):
		errs = validateObjectFormat(parentKey, variableValue, variableSchema)
	case string(ArrayType):
		errs = validateArrayFormat(parentKey, variableValue, variableSchema)
	case string(StringType):
		errs = validateStringFormat(parentKey, variableValue, variableSchema)
	case string(BooleanType):
		errs = validateBooleanFormat(parentKey, variableValue)
	case string(IntegerType), string(NumberType):
		errs = validateNumberFormat(parentKey, variableValue, variableSchema, varType)
	}

	return errs
}

func validateObjectFormat(parentKey string, variableValue interface{}, variableSchema map[string]interface{}) (errs []error) {
	errs = make([]error, 0)

	if variableValueMap, ok := variableValue.(map[string]interface{}); !ok {
		errs = append(errs, errors.Errorf("Key '%s' should be a map, type provided: %T", parentKey, variableValue))

		return errs
	} else {
		if objSchema, ok := variableSchema[string(PropertiesKey)]; ok {
			for k, v := range variableValueMap {
				kSchema, kSchemaExist := objSchema.(map[string]interface{})[k]

				if !kSchemaExist {
					errs = append(errs, errors.Errorf("Key '%s' is not expected in key %s.", k, parentKey))
				} else {
					errs = append(errs, validateSchemaFormat(k, v, kSchema.(map[string]interface{}))...)
				}
			}
		} else {
			for k, v := range variableValueMap {
				errs = append(errs, validateSchemaFormat(k, v, variableSchema[string(AdditionalPropertiesKey)].(map[string]interface{}))...)
			}
		}
	}

	return errs
}

func validateArrayFormat(parentKey string, variableValue interface{}, variableSchema map[string]interface{}) (errs []error) {
	errs = make([]error, 0)

	if variableValueArray, ok := variableValue.([]interface{}); !ok {
		errs = append(errs, errors.Errorf("Key '%s' should be an array, type provided: %T", parentKey, variableValue))

		return errs
	} else {
		for _, it := range variableValueArray {
			errs = append(errs, validateSchemaFormat(parentKey, it, variableSchema[string(ItemsKey)].(map[string]interface{}))...)
		}
	}

	return errs
}

func validateStringFormat(parentKey string, variableValue interface{}, variableSchema map[string]interface{}) (errs []error) {
	errs = make([]error, 0)

	if _, ok := variableValue.(string); !ok {
		errs = append(errs, errors.Errorf("Key '%s' should be a string, type provided: %T", parentKey, variableValue))

		return errs
	}

	if regexPattern, ok := variableSchema[string(PatternKey)]; ok {
		regex, err := regexp.Compile(regexPattern.(string))

		if err == nil {
			if !regex.MatchString(variableValue.(string)) {
				errs = append(errs, errors.Errorf("Key '%s' doesn't match regular expression '%s', value provided: '%s'", parentKey, regexPattern, variableValue))
			}
		}
	}

	if minLen, ok := variableSchema[string(MinLengthKey)]; ok {
		varLen := len(variableValue.(string))

		if varLen < int(minLen.(float64)) {
			errs = append(errs, errors.Errorf("Key '%s' should have a string longer than '%v', value provided: '%s' (%v)", parentKey, minLen, variableValue, varLen))
		}
	}

	if maxLen, ok := variableSchema[string(MaxLengthKey)]; ok {
		varLen := len(variableValue.(string))

		if varLen > int(maxLen.(float64)) {
			errs = append(errs, errors.Errorf("Key '%s' should have a string shorter than '%v', value provided: '%s' (%v)", parentKey, maxLen, variableValue, varLen))
		}
	}

	return errs
}

func validateBooleanFormat(parentKey string, variableValue interface{}) (errs []error) {
	errs = make([]error, 0)

	if _, ok := variableValue.(bool); !ok {
		errs = append(errs, errors.Errorf("Key '%s' should be a boolean, type provided: %T", parentKey, variableValue))

		return errs
	}

	return errs
}

func validateNumberFormat(parentKey string, variableValue interface{}, variableSchema map[string]interface{}, varType string) (errs []error) {
	errs = make([]error, 0)

	if varType == string(IntegerType) {
		// JSON UnMarshal always store numbers as Float64.
		if _, ok := variableValue.(float64); !ok || variableValue.(float64) != float64(int(variableValue.(float64))) {
			errs = append(errs, errors.Errorf("Key '%s' should be an integer, type provided: %T", parentKey, variableValue))

			return errs
		}
	} else {
		if _, ok := variableValue.(float64); !ok {
			errs = append(errs, errors.Errorf("Key '%s' should be a float, type provided: %T", parentKey, variableValue))

			return errs
		}
	}

	if minValue, ok := variableSchema[string(MinimumKey)]; ok {
		if variableValue.(float64) < minValue.(float64) {
			errs = append(errs, errors.Errorf("Key '%s' should be greater than '%v', value provided: '%s'", parentKey, minValue, variableValue))
		}
	}

	if maxValue, ok := variableSchema[string(MaximumKey)]; ok {
		if variableValue.(float64) > maxValue.(float64) {
			errs = append(errs, errors.Errorf("Key '%s' should be lower than '%v', value provided: '%s'", parentKey, maxValue, variableValue))
		}
	}

	return errs
}
