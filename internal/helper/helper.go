/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"fmt"
	"log"
	"os"
)

const (
	DEV     = "DEV"
	TmcMode = "TMC_MODE"
)

func GetFirstElementOf(parent string, children ...string) (key string) {
	if len(children) == 0 {
		return parent
	}

	key = parent
	for _, value := range children {
		key = fmt.Sprintf("%s.0.%s", key, value)
	}

	return key
}

// Terraform always reads float value as float64.
func readFloat(input interface{}, key string) float64 {
	data, ok := input.(float64)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to float64 for attribute %s \n", input, key)

		if os.Getenv(TmcMode) == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readInt(input interface{}, key string) int {
	data, ok := input.(int)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to int for attribute %s \n", input, key)

		if os.Getenv(TmcMode) == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readBool(input interface{}, key string) bool {
	data, ok := input.(bool)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to bool for attribute %s \n", input, key)

		if os.Getenv(TmcMode) == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readString(input interface{}, key string) string {
	data, ok := input.(string)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to string for attribute %s \n", input, key)

		if os.Getenv(TmcMode) == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func SetPrimitiveValue(input, model interface{}, key string) {
	switch m := model.(type) {
	case *float64:
		// Store the address of the model
		modelPtr, _ := model.(*float64)
		// Assign the type cast value to the model address pointed value
		*modelPtr = readFloat(input, key)
	case *float32:
		modelPtr, _ := model.(*float32)
		*modelPtr = float32(readFloat(input, key))
	case *int:
		modelPtr, _ := model.(*int)
		*modelPtr = readInt(input, key)
	case *int8:
		modelPtr, _ := model.(*int8)
		*modelPtr = int8(readInt(input, key))
	case *int16:
		modelPtr, _ := model.(*int16)
		*modelPtr = int16(readInt(input, key))
	case *int32:
		modelPtr, _ := model.(*int32)
		*modelPtr = int32(readInt(input, key))
	case *int64:
		modelPtr, _ := model.(*int64)
		*modelPtr = int64(readInt(input, key))
	case *bool:
		modelPtr, _ := model.(*bool)
		*modelPtr = readBool(input, key)
	case *string:
		modelPtr, _ := model.(*string)
		*modelPtr = readString(input, key)
	default:
		fmt.Printf("[ERROR]: Internal err, invalid use of SetPrimitive function for %s", key)

		if os.Getenv(TmcMode) == DEV {
			log.Fatalf("[ERROR}: SetPrimitive works on PassByReference and only for Primitive types. Got [%T]", m)
		}
	}
}
