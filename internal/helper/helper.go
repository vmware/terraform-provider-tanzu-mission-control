// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helper

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DEV = "DEV"

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

		if os.Getenv("TMC_MODE") == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readInt(input interface{}, key string) int {
	data, ok := input.(int)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to int for attribute %s \n", input, key)

		if os.Getenv("TMC_MODE") == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readBool(input interface{}, key string) bool {
	data, ok := input.(bool)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to bool for attribute %s \n", input, key)

		if os.Getenv("TMC_MODE") == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

func readString(input interface{}, key string) string {
	data, ok := input.(string)
	if !ok {
		fmt.Printf("[ERROR]: Unable to covert %T to string for attribute %s \n", input, key)

		if os.Getenv("TMC_MODE") == DEV {
			log.Fatalf("[ERROR]: Invalid type conversion for the %s. Please check the schema", key)
		}
	}

	return data
}

// nolint:gosec
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

		if os.Getenv("TMC_MODE") == DEV {
			log.Fatalf("[ERROR}: SetPrimitive works on PassByReference and only for Primitive types. Got [%T]", m)
		}
	}
}

func SetPrimitiveList[T any](data any, key string) []T {
	list, ok := data.([]any)
	if !ok || len(list) < 1 || (len(list) == 1 && list[0] == nil) {
		return nil
	}

	out := make([]T, 0, len(list))

	for _, v := range list {
		var value T

		SetPrimitiveValue(v, &value, key)

		out = append(out, value)
	}

	return out
}

func BoolPointer(b bool) *bool {
	return &b
}

func Float64Pointer(f float64) *float64 {
	return &f
}

func StringPointer(s string) *string {
	return &s
}

func PtrString[T ~string](val *T) string {
	if val == nil {
		return ""
	}

	return (string)(*val)
}

func UpdateDataSourceSchema(d *schema.Schema) *schema.Schema {
	dv := &schema.Schema{
		Type:        d.Type,
		Description: d.Description,
		Computed:    true,
		Elem:        d.Elem,
	}

	return dv
}

func ConvertToString(value interface{}, sliceSep string) string {
	outputStr := ""

	switch v := value.(type) {
	case float64, float32:
		outputStr = fmt.Sprintf("%f", v)
	case []byte:
		outputStr = string(v)
	case int, int8, int16, int32, int64:
		outputStr = fmt.Sprintf("%d", v)
	case bool:
		outputStr = fmt.Sprintf("%t", v)
	case []interface{}:
		strSlice := make([]string, len(v))

		for i, elem := range v {
			strSlice[i] = ConvertToString(elem, sliceSep)
		}

		outputStr = strings.Join(strSlice, sliceSep)
	case string:
		outputStr = v
	}

	return outputStr
}

func IsEmptyInterface(value interface{}) bool {
	if value == nil {
		return true
	}

	switch value := value.(type) {
	case map[string]interface{}:
		return len(value) == 0
	case []interface{}:
		return len(value) == 0
	case string:
		return value == ""
	}

	return false
}

func GetAllMapsKeys(maps ...map[string]interface{}) map[string]bool {
	keys := make(map[string]bool)

	for _, m := range maps {
		for key := range m {
			keys[key] = true
		}
	}

	return keys
}

// nolint:godot
// DatasourceSchemaFromResourceSchema is a recursive func that
// converts an existing Resource schema to a Datasource schema.
// All schema elements are copied, but certain attributes are ignored or changed:
// - all attributes have Computed = true
// - all attributes have ForceNew, Required = false
// - Validation funcs and attributes (e.g. MaxItems) are not copied
func DatasourceSchemaFromResourceSchema(rs map[string]*schema.Schema) map[string]*schema.Schema {
	ds := make(map[string]*schema.Schema, len(rs))

	for k, v := range rs {
		dv := &schema.Schema{
			Computed:    true,
			ForceNew:    false,
			Required:    false,
			Description: v.Description,
			Type:        v.Type,
		}

		switch v.Type {
		case schema.TypeSet:
			dv.Set = v.Set
			fallthrough
		case schema.TypeList:
			// List & Set types are generally used for 2 cases:
			// - a list/set of simple primitive values (e.g. list of strings)
			// - a sub resource
			if elem, ok := v.Elem.(*schema.Resource); ok {
				// handle the case where the Element is a sub-resource
				dv.Elem = &schema.Resource{
					Schema: DatasourceSchemaFromResourceSchema(elem.Schema),
				}
			} else {
				// handle simple primitive case
				dv.Elem = v.Elem
			}

		default:
			// Elem of all other types are copied as-is
			dv.Elem = v.Elem
		}

		ds[k] = dv
	}

	return ds
}

// fixDatasourceSchemaFlags is a convenience func that toggles the Computed,
// Optional + Required flags on a schema element. This is useful when the schema
// has been generated (using `DatasourceSchemaFromResourceSchema` above for
// example) and therefore the attribute flags were not set appropriately when
// first added to the schema definition. Currently only supports top-level
// schema elements.
func FixDatasourceSchemaFlags(schema map[string]*schema.Schema, required bool, keys ...string) {
	for _, v := range keys {
		schema[v].Computed = false
		schema[v].Optional = !required
		schema[v].Required = required
	}
}

func AddRequiredFieldsToSchema(schema map[string]*schema.Schema, keys ...string) {
	FixDatasourceSchemaFlags(schema, true, keys...)
}

func AddOptionalFieldsToSchema(schema map[string]*schema.Schema, keys ...string) {
	FixDatasourceSchemaFlags(schema, false, keys...)
}

func DeleteFieldsFromSchema(schema map[string]*schema.Schema, keys ...string) {
	for _, key := range keys {
		delete(schema, key)
	}
}
