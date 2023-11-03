/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	arrIndexExceededMsg     = "array index exceeded"
	arrayCannotBeReachedMsg = "arrays can't be reached"
)

func (converter *TFSchemaModelConverter[T]) buildTFValue(modelJSONData *map[string]interface{}, mapValue interface{}, arrIndexer *ArrIndexer) (interface{}, error) {
	if modelJSONData == nil || mapValue == nil {
		return nil, nil
	}

	var (
		err           error
		tfSchemaValue interface{}
	)

	switch mapValue := mapValue.(type) {
	case *BlockToStruct, *Map:
		return converter.tfHandleBlockMap(modelJSONData, mapValue, arrIndexer)
	case *BlockToStructSlice:
		return converter.tfHandleBlockStructSlice(modelJSONData, mapValue, arrIndexer)
	case *BlockSliceToStructSlice:
		return converter.tfHandleBlockSliceStructSlice(modelJSONData, mapValue, arrIndexer)
	case *ListToStruct:
		return converter.tfHandleListStruct(modelJSONData, mapValue, arrIndexer)
	case *EvaluatedField:
		var modelValue interface{}

		modelField := mapValue.Field
		modelValue, err = converter.getModelValue(modelJSONData, modelField, arrIndexer)

		if err == nil {
			tfSchemaValue = mapValue.EvalFunc(ConstructTFSchema, modelValue)
		} else if strings.Contains(err.Error(), arrayCannotBeReachedMsg) && strings.HasSuffix(modelField, ArrayFieldMarker) {
			err = nil
		}
	case string:
		tfSchemaValue, err = converter.getModelValue(modelJSONData, mapValue, arrIndexer)
	}

	return tfSchemaValue, err
}

func (converter *TFSchemaModelConverter[T]) tfHandleBlockMap(modelJSONData *map[string]interface{}, mapValue interface{}, arrIndexer *ArrIndexer) (tfSchemaValue interface{}, err error) {
	_, isMap := mapValue.(*Map)

	if isMap {
		if allFlagsKeyValue, exists := (*mapValue.(*Map))[AllMapKeysFieldMarker]; exists {
			modelValue, _ := converter.buildTFValue(modelJSONData, allFlagsKeyValue, arrIndexer)

			if modelValue != nil {
				if tfSchemaValue == nil {
					tfSchemaValue = make(map[string]interface{})
				}

				for key, value := range modelValue.(map[string]interface{}) {
					tfSchemaValue.(map[string]interface{})[key] = value
				}
			}
		}

		newBlock := BlockToStruct(*(mapValue.(*Map).Copy([]string{AllMapKeysFieldMarker})))
		mapValue = &newBlock
	}

	for elemKey, elemValue := range *mapValue.(*BlockToStruct) {
		modelValue, err := converter.buildTFValue(modelJSONData, elemValue, arrIndexer)

		if modelValue != nil {
			if tfSchemaValue == nil {
				tfSchemaValue = make(map[string]interface{})
			}

			tfSchemaValue.(map[string]interface{})[elemKey] = modelValue
		} else if err != nil {
			return nil, err
		}
	}

	if tfSchemaValue != nil && !isMap {
		tfSchemaValue = []interface{}{tfSchemaValue}
	}

	return tfSchemaValue, err
}

func (converter *TFSchemaModelConverter[T]) tfHandleBlockStructSlice(modelJSONData *map[string]interface{}, mapValue *BlockToStructSlice, arrIndexer *ArrIndexer) (tfSchemaValue interface{}, err error) {
	var (
		modelValue  interface{}
		tfElemValue map[string]interface{}
	)

	for i, elemMap := range *mapValue {
		arrIndexer.New()

		for err == nil {
			modelValue, err = converter.buildTFValue(modelJSONData, elemMap, arrIndexer)

			if modelValue != nil {
				if tfElemValue == nil {
					tfElemValue = make(map[string]interface{})
				}

				for key, value := range modelValue.([]interface{})[0].(map[string]interface{}) {
					if vArr, ok := value.([]interface{}); ok {
						if _, exists := tfElemValue[key]; !exists {
							tfElemValue[key] = make([]interface{}, 0)
						}

						for _, v := range vArr {
							tfElemValue[key] = append(tfElemValue[key].([]interface{}), v)
						}
					} else {
						tfElemValue[key] = value
					}
				}
			}

			isEvaluatedField := false

			for _, v := range *elemMap {
				_, isEvaluatedField = v.(*EvaluatedField)
			}

			if isEvaluatedField {
				break
			}

			if err == nil {
				arrIndexer.IncrementLastIndex()
			} else if err.Error() == arrIndexExceededMsg && i+1 < len(*mapValue) {
				err = nil

				break
			}
		}

		arrIndexer.RemoveLastIndex()
	}

	if err != nil {
		errMsg := err.Error()

		if errMsg == arrIndexExceededMsg && !arrIndexer.IsActive() {
			err = nil
		} else if strings.Contains(errMsg, arrayCannotBeReachedMsg) {
			splitErrMsg := strings.Split(errMsg, ":")
			numOfUnreachableArrays, _ := strconv.Atoi(splitErrMsg[1])
			numOfUnreachableArrays--

			if numOfUnreachableArrays > 0 {
				err = errors.New(fmt.Sprintf("%s:%v", arrayCannotBeReachedMsg, numOfUnreachableArrays))
			} else {
				err = nil
			}

			return nil, err
		}
	}

	if tfElemValue != nil {
		tfSchemaValue = make([]interface{}, 0)
		tfSchemaValue = append(tfSchemaValue.([]interface{}), tfElemValue)
	}

	return tfSchemaValue, err
}

func (converter *TFSchemaModelConverter[T]) tfHandleBlockSliceStructSlice(modelJSONData *map[string]interface{}, mapValue *BlockSliceToStructSlice, arrIndexer *ArrIndexer) (tfSchemaValue interface{}, err error) {
	var modelValue interface{}

	for i, elemMap := range *mapValue {
		arrIndexer.New()

		for err == nil {
			modelValue, err = converter.buildTFValue(modelJSONData, elemMap, arrIndexer)

			if modelValue != nil {
				if tfSchemaValue == nil {
					tfSchemaValue = make([]interface{}, 0)
				}

				tfSchemaValue = append(tfSchemaValue.([]interface{}), modelValue.([]interface{})[0])
			}

			if err == nil {
				arrIndexer.IncrementLastIndex()
			} else if err.Error() == arrIndexExceededMsg && i+1 < len(*mapValue) {
				err = nil

				break
			}
		}

		arrIndexer.RemoveLastIndex()
	}

	if err != nil {
		errMsg := err.Error()

		if errMsg == arrIndexExceededMsg && !arrIndexer.IsActive() {
			err = nil
		} else if strings.Contains(errMsg, arrayCannotBeReachedMsg) {
			splitErrMsg := strings.Split(errMsg, ":")
			numOfUnreachableArrays, _ := strconv.Atoi(splitErrMsg[1])
			numOfUnreachableArrays--

			if numOfUnreachableArrays > 0 {
				err = errors.New(fmt.Sprintf("%s:%v", arrayCannotBeReachedMsg, numOfUnreachableArrays))
			} else {
				err = nil
			}

			return nil, err
		}
	}

	return tfSchemaValue, err
}

func (converter *TFSchemaModelConverter[T]) tfHandleListStruct(modelJSONData *map[string]interface{}, mapValue *ListToStruct, arrIndexer *ArrIndexer) (tfSchemaValue interface{}, err error) {
	var (
		arr        []interface{}
		modelValue interface{}
	)

	for err == nil {
		modelValue, err = converter.getModelValue(modelJSONData, (*mapValue)[0], arrIndexer)
		arrIndexer.IncrementLastIndex()

		if modelValue != nil {
			arr = append(arr, modelValue)
		}
	}

	if len(arr) > 0 {
		tfSchemaValue = arr
	}

	return tfSchemaValue, err
}

func (converter *TFSchemaModelConverter[T]) getModelValue(modelJSONData *map[string]interface{}, mapValue string, arrIndexer *ArrIndexer) (interface{}, error) {
	var (
		err                error
		lastIndex          int
		arrIndexerPosition int

		modelRootValue  interface{} = *modelJSONData
		modelValuePaths             = strings.Split(mapValue, converter.getModelPathSeparator())
	)

	for i := 0; i < len(modelValuePaths); i++ {
		nextModelPath := strings.ReplaceAll(modelValuePaths[i], ArrayFieldMarker, "")

		if nextModelPath == AllMapKeysFieldMarker {
			newMap := make(map[string]interface{})

			for k, v := range modelRootValue.(map[string]interface{}) {
				newMap[k] = v
			}

			modelRootValue = newMap

			break
		} else {
			switch currentRootValue := modelRootValue.(type) {
			case map[string]interface{}:
				modelRootValue = currentRootValue[nextModelPath]
			case []interface{}:
				switch {
				case len(currentRootValue) == 0:
					err = errors.New(arrIndexExceededMsg)

					return nil, err
				case len(currentRootValue) > 0 && i < len(modelValuePaths):
					lastIndex = arrIndexer.IndicesInts[arrIndexerPosition]

					if lastIndex >= len(currentRootValue) {
						err = errors.New(arrIndexExceededMsg)

						return nil, err
					}

					modelRootValue = currentRootValue[lastIndex].(map[string]interface{})[nextModelPath]
					arrIndexerPosition++
				default:
					modelRootValue = currentRootValue[0].(map[string]interface{})[nextModelPath]
				}
			}
		}

		if modelRootValue == nil {
			arraysCount := strings.Count(mapValue, ArrayFieldMarker)

			switch {
			case err != nil:
				return nil, err
			case arrIndexerPosition < arraysCount:
				// This is in case a root to array is nil
				err = errors.New(fmt.Sprintf("%s:%v", arrayCannotBeReachedMsg, arraysCount-arrIndexerPosition))

				return nil, err
			default:
				return nil, nil
			}
		}
	}

	return modelRootValue, err
}
