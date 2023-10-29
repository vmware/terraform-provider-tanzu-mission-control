/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"fmt"
	"strings"
)

// BlockToStruct Converts a Terraform Block to a Swagger Struct (1:1).
type BlockToStruct map[string]interface{}

// BlockToStructSliceField Converts a Terraform Block to a Slice of Swagger Structs (1:N).
type BlockToStructSliceField struct {
	Field   string
	Mapping BlockToStructSlice
}

type BlockToStructSlice []*BlockToStruct

// BlockSliceToStructSliceField Converts a Terraform Block to a Slice of Swagger Structs (N:N).
type BlockSliceToStructSliceField struct {
	Field   string
	Mapping BlockToStructSlice
}

type BlockSliceToStructSlice []*BlockToStruct

// Map Converts a Terraform Block/Map to Swagger Struct/Map.
type Map map[string]interface{}

// ListToStruct Converts a Terraform primitive Slice field to a Slice of Swagger Structs (1:N).
type ListToStruct []string

// EvaluatedField If needed to intervene in specific fields, this helps you connect a function to the converter.
type EvaluatedField struct {
	Field    string
	EvalFunc EvaluationFunc
}

type EvaluationFunc func(EvaluationMode, interface{}) interface{}

type EvaluationMode string

const (
	ConstructModel    EvaluationMode = "CONSTRUCT_MODEL"
	ConstructTFSchema EvaluationMode = "CONSTRUCT_TF_SCHEMA"
)

// Copy Creates a copy of a Map object.
// excludeKeys argument can be used to exclude certain keys to be copied.
func (curMap *Map) Copy(excludedKeys []string) Map {
	nMap := make(Map)

	for k, v := range *curMap {
		if len(excludedKeys) > 0 {
			isExcluded := false

			for _, excludedKey := range excludedKeys {
				if excludedKey == k {
					isExcluded = true
					break
				}
			}

			if !isExcluded {
				nMap[k] = v
			}
		} else {
			nMap[k] = v
		}
	}

	return nMap
}

// UnpackSchema Unpacks a schema to a higher level schema, useful for data sources which list an individual Swagger API Model.
func (b *BlockToStruct) UnpackSchema(modelPathSeparator string, mapValue interface{}, prefix string) interface{} {
	var elem interface{}

	if mapValue == nil {
		mapValue = b
	}

	switch mapValue := mapValue.(type) {
	case *BlockToStruct, *Map:
		if _, ok := mapValue.(*BlockToStruct); ok {
			elem = &BlockToStruct{}

			for key, value := range *mapValue.(*BlockToStruct) {
				(*elem.(*BlockToStruct))[key] = b.UnpackSchema(modelPathSeparator, value, prefix)
			}
		} else {
			elem = &Map{}

			for key, value := range *mapValue.(*Map) {
				(*elem.(*Map))[key] = b.UnpackSchema(modelPathSeparator, value, prefix)
			}
		}
	case *BlockToStructSlice:
		elem = &BlockToStructSlice{}

		for _, elemMap := range *mapValue {
			elemValue := b.UnpackSchema(modelPathSeparator, elemMap, prefix)
			*elem.(*BlockToStructSlice) = append(*elem.(*BlockToStructSlice), elemValue.(*BlockToStruct))
		}
	case *BlockSliceToStructSlice:
		elem = &BlockSliceToStructSlice{}

		for _, elemMap := range *mapValue {
			elemValue := b.UnpackSchema(modelPathSeparator, elemMap, prefix)
			*elem.(*BlockSliceToStructSlice) = append(*elem.(*BlockSliceToStructSlice), elemValue.(*BlockToStruct))
		}
	case *ListToStruct:
		elem = &ListToStruct{}
		elemValue := b.UnpackSchema(modelPathSeparator, (*mapValue)[0], prefix)
		*elem.(*ListToStruct) = append(*elem.(*ListToStruct), elemValue.(string))
	case *EvaluatedField:
		elem = &EvaluatedField{
			Field:    strings.Join([]string{prefix, (mapValue).Field}, modelPathSeparator),
			EvalFunc: (mapValue).EvalFunc,
		}
	case string:
		return fmt.Sprintf("%s%s%s", prefix, modelPathSeparator, mapValue)
	}

	return elem
}
