/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type APIModel interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}

type TFSchemaModelConverter[T APIModel] struct {
	TFModelMap         *BlockToStruct
	ModelPathSeparator string
}

const (
	DefaultModelPathSeparator = "|"
	ArrayFieldMarker          = "[]"
	AllMapKeysFieldMarker     = "*"
)

// ### Public Funcs ###.

// ConvertTFSchemaToAPIModel This function converts a given TF Schema data (aka schema.ResourceData) to Swagger API Model.
// Arguments for the function are:
// data - a schema.ResourceData object.
// buildOnlyKeys - a []string for specifying specific keys to be built in the model.
// keys should be supplied relatively to the Terraform schema structure, for nested keys it should be supplied with a separator.
// If no keys supplied in the slice, the function will build the entire model.
func (converter *TFSchemaModelConverter[T]) ConvertTFSchemaToAPIModel(data *schema.ResourceData, buildOnlyKeys []string) (modelPtr T, err error) {
	var (
		arrIndexer ArrIndexer

		typ       = reflect.TypeOf(modelPtr)
		elem      = typ.Elem()
		modelJSON = make(BlockToStruct)
	)

	modelPtr = reflect.New(elem).Interface().(T)

	if len(buildOnlyKeys) > 0 {
		for _, subKey := range buildOnlyKeys {
			converter.handleOffsetBuildModelField(&modelJSON, data, strings.Split(subKey, converter.getModelPathSeparator()), &arrIndexer)
		}
	} else {
		for mapKey, mapValue := range *converter.TFModelMap {
			schemaData := data.Get(mapKey)
			converter.buildModelField(&modelJSON, schemaData, mapValue, &arrIndexer)
		}
	}

	jsonBytes, _ := json.Marshal(modelJSON)
	err = modelPtr.UnmarshalBinary(jsonBytes)

	return modelPtr, err
}

// FillTFSchema This function converts a given Swagger API Model to TF Schema data (aka schema.ResourceData) structure and fills the schema.
// Arguments for the function are:
// modelPtr - a Swagger API Model object.
// data - a schema.ResourceData object.
func (converter *TFSchemaModelConverter[T]) FillTFSchema(modelPtr T, data *schema.ResourceData) error {
	var (
		modelJSONData map[string]interface{}
		arrIndexer    ArrIndexer
		tfValue       interface{}
	)

	jsonBytes, err := modelPtr.MarshalBinary()

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, &modelJSONData)

	if err != nil {
		return err
	}

	for mapKey, mapValue := range *converter.TFModelMap {
		tfValue, err = converter.buildTFValue(&modelJSONData, mapValue, &arrIndexer)

		if err != nil {
			return err
		}

		if tfValue != nil {
			err = data.Set(mapKey, tfValue)

			if err != nil {
				return err
			}
		}
	}

	return err
}

// UnpackSchema This function helps you create a higher level Schema.
// It utilizes the implementation of schema unpacking implemented for BlockToStruct types.
// As an example, if you have converter or converter mapping of an object like Target Location and you need to create another schema for a slice of Target Locations,
// then you can use the converter or the mapping to unpack the schema for the higher level struct containing the list of Target Locations.
func (converter *TFSchemaModelConverter[T]) UnpackSchema(prefix string) *BlockToStruct {
	newTFModelMap := converter.TFModelMap.UnpackSchema(converter.getModelPathSeparator(), nil, prefix).(*BlockToStruct)

	return newTFModelMap
}

// ### Private Funcs ###.
func (converter *TFSchemaModelConverter[T]) handleOffsetBuildModelField(rootModelJSON *BlockToStruct, data *schema.ResourceData, offsetPaths []string, arrIndexer *ArrIndexer) {
	rootPath := offsetPaths[0]
	rootSchemaData := data.Get(rootPath)
	rootMapValue := (*converter.TFModelMap)[rootPath]

	if len(offsetPaths) > 1 {
		rootSchemaData = rootSchemaData.([]interface{})[0]

		for i := 1; i < len(offsetPaths); i++ {
			s, ok := rootSchemaData.([]interface{})

			if !ok {
				rootSchemaData = rootSchemaData.(map[string]interface{})[offsetPaths[i]]
			} else {
				if len(s) > 1 {
					rootSchemaData = s
				} else if len(s) == 1 {
					rootSchemaData = s[0].(map[string]interface{})[offsetPaths[i]]
				}
			}

			rootMapValue = rootMapValue.(BlockToStruct)[offsetPaths[i]]
		}
	}

	converter.buildModelField(rootModelJSON, rootSchemaData, rootMapValue, arrIndexer)
}

func (converter *TFSchemaModelConverter[T]) getModelPathSeparator() string {
	if converter.ModelPathSeparator == "" {
		return DefaultModelPathSeparator
	}

	return converter.ModelPathSeparator
}
