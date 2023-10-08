/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"reflect"
	"strings"
)

func buildModelField(modelJSON *BlockToStruct, schemaData interface{}, mapValue interface{}, arrIndexer *ArrIndexer) {
	if schemaData == nil || mapValue == nil {
		return
	}

	switch mapValue := mapValue.(type) {
	case *BlockToStruct:
		modelHandleBlockStruct(modelJSON, schemaData, mapValue, arrIndexer)
	case *Map:
		modelHandleBlockMap(modelJSON, schemaData, mapValue, arrIndexer)
	case *BlockToStructSlice:
		modelHandleBlockStructSlice(modelJSON, schemaData, mapValue, arrIndexer)
	case *BlockSliceToStructSlice:
		modelHandleBlockSliceStructSlice(modelJSON, schemaData, mapValue, arrIndexer)
	case *ListToStruct:
		modelHandleListStruct(modelJSON, schemaData, mapValue, arrIndexer)
	case *EvaluatedField:
		modelField := mapValue.Field
		modelValue := mapValue.EvalFunc(ConstructModel, schemaData)
		setModelValue(modelJSON, modelField, modelValue, arrIndexer)
	case string:
		modelField := mapValue
		modelValue := schemaData
		setModelValue(modelJSON, modelField, modelValue, arrIndexer)
	}
}

func modelHandleBlockStruct(modelJSON *BlockToStruct, schemaData interface{}, mapValue *BlockToStruct, arrIndexer *ArrIndexer) {
	if schemaDataSlice, ok := schemaData.([]interface{}); ok && len(schemaDataSlice) > 0 {
		rootSchemaDict, _ := schemaDataSlice[0].(map[string]interface{})

		for key, value := range *mapValue {
			buildModelField(modelJSON, rootSchemaDict[key], value, arrIndexer)
		}
	}
}

func modelHandleBlockMap(modelJSON *BlockToStruct, schemaData interface{}, mapValue *Map, arrIndexer *ArrIndexer) {
	if rootSchemaDict, ok := schemaData.(map[string]interface{}); ok {
		definedKeysMapValue := mapValue.Copy([]string{"*"})

		if allKeysFlagMapValue, exists := (*mapValue)["*"]; exists {
			for key, value := range rootSchemaDict {
				var dynamicMapValue interface{}

				if allKeysFlagMapStr, ok := allKeysFlagMapValue.(string); ok {
					dynamicMapValue = strings.ReplaceAll(allKeysFlagMapStr, "*", key)
				} else {
					dynamicMapValue = allKeysFlagMapValue
				}

				buildModelField(modelJSON, value, dynamicMapValue, arrIndexer)
			}
		}

		for key, value := range definedKeysMapValue {
			buildModelField(modelJSON, rootSchemaDict[key], value, arrIndexer)
		}
	}
}

func modelHandleBlockStructSlice(modelJSON *BlockToStruct, schemaData interface{}, mapValue *BlockToStructSlice, arrIndexer *ArrIndexer) {
	if len(schemaData.([]interface{})) > 0 {
		arrIndexer.New()

		for _, elemTypeMap := range *mapValue {
			for elemMapKey, elemMapValue := range *elemTypeMap {
				var schemaValue, _ = (schemaData.([]interface{}))[0].(map[string]interface{})[elemMapKey]

				if schemaValue != nil {
					if _, ok := elemMapValue.(*ListToStruct); ok {
						buildModelField(modelJSON, schemaValue, elemMapValue, arrIndexer)
					} else {
						for _, item := range schemaValue.([]interface{}) {
							buildModelField(modelJSON, []interface{}{item}, elemMapValue, arrIndexer)
							arrIndexer.IncrementLastIndex()
						}
					}
				}
			}
		}

		arrIndexer.RemoveLastIndex()
	}
}

func modelHandleBlockSliceStructSlice(modelJSON *BlockToStruct, schemaData interface{}, mapValue *BlockSliceToStructSlice, arrIndexer *ArrIndexer) {
	if len(schemaData.([]interface{})) > 0 {
		for _, elemTypeMap := range *mapValue {
			arrIndexer.New()

			for _, item := range schemaData.([]interface{}) {
				var _, ok = item.(map[string]interface{})

				if ok {
					buildModelField(modelJSON, []interface{}{item}, elemTypeMap, arrIndexer)
					arrIndexer.IncrementLastIndex()
				}
			}

			arrIndexer.RemoveLastIndex()
		}
	}
}

func modelHandleListStruct(modelJSON *BlockToStruct, schemaData interface{}, mapValue *ListToStruct, arrIndexer *ArrIndexer) {
	if reflect.TypeOf(schemaData).Kind() == reflect.Slice {
		sliceValue := reflect.ValueOf(schemaData)

		for i := 0; i < sliceValue.Len(); i++ {
			val := sliceValue.Index(i).Interface()
			setModelValue(modelJSON, (*mapValue)[0], val, arrIndexer)
			arrIndexer.IncrementLastIndex()
		}
	}
}

func setModelValue(model *BlockToStruct, field string, value interface{}, arrIndexer *ArrIndexer) {
	if !strings.Contains(field, ".") {
		(*model)[field] = value
	} else {
		fieldPaths := strings.Split(field, ".")
		arrIndices := arrIndexer.GetAllIndexes()
		leafField := strings.ReplaceAll(fieldPaths[len(fieldPaths)-1], "[]", "")
		arrayFields := 0
		parentField := *model

		for i := 0; i < len(fieldPaths)-1; i++ {
			fieldName := fieldPaths[i]

			if !strings.Contains(fieldName, "[]") {
				if _, ok := parentField[fieldName]; !ok {
					parentField[fieldName] = make(map[string]interface{})
				}

				parentField = parentField[fieldName].(map[string]interface{})
			} else {
				var object map[string]interface{}

				fieldName = strings.ReplaceAll(fieldName, "[]", "")
				arrayIndex := arrIndices[arrayFields]

				arrayFields++

				if _, ok := parentField[fieldName]; !ok {
					parentField[fieldName] = make([]map[string]interface{}, 0)
				}

				parentFieldArray := parentField[fieldName].([]map[string]interface{})

				if arrayIndex < len(parentFieldArray) {
					object = parentFieldArray[arrayIndex]
				} else {
					object = make(map[string]interface{})
					parentField[fieldName] = append(parentFieldArray, object)
				}

				parentField = object
			}
		}

		parentField[leafField] = value
	}
}
