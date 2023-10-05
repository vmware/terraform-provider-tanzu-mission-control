//go:build ignore
// +build ignore

package help

/*
This Module defines an object to help convert Terraform resource data to Swagger API models and vice versa

In order to use that, one should simply create a BlockToStruct type object (See ./map_types.go) and define the tfModelMap as follows:

######### Rules #########
- tfModelMap should always be a BlockToStruct type
- The key-value concept of the map should follow the format TF_SCHEMA_KEY : SWAGGER_MODEL_MAPPING
	- Keys are always strings and should always be set to the name of the scoped Terraform resource field. (e.g: "target_provider")
	- The following values are valid:
	  - For every Primitive data type in a Swagger model (int, string, []string/int.., ...) the value (String) should always be the full json name along the entire hierarchy
			of the Root Swagger Model (struct) concatenated by "."
        	e.g: "spec.bucket", "fullName.name"
      - For every Struct data type in a swagger model there are these are the valid (Mapping) values:
        - BlockToStruct - Most common use-case, this definition maps a terraform block (Type=TypeList, MaxItems=1) to a Swagger API Model (Struct)
		- Map - This should be used to map a terraform field (Type=TypeMap) to a Swagger API Model (Struct) or map[string]interface{} field in a struct
        - ListToStruct - Can be used if desired to map a terraform resource field (Type=TypeList) to a
			Swagger API Model (Struct) which only contains a single field to populate.
			e.g: {"cluster_groups": []string{"spec.assignedGroups[].clustergroup.name"}}
		- BlockToStructSlice - This is useful when mapping a terraform block (Type=TypeList, MaxItems=1) to a (Slice) of Swagger API Model (Struct).
			BlockToStructSlice is itself a slice and each value of the slice is essentially a mapping to a valid item in the destination (Slice) field of a Swagger API Model (Struct).
			The hierarchical path name of the Swagger API Model (Struct) field mapped to the in the inner definitions of the scoped terraform resource (Mapping)
			should be included with "[]" next to the Swagger API Model (Struct) name that corresponds to the (Slice) definition.
			e.g: "spec.assignedGroups[].cluster.managementClusterName"

			The level of the map value for this case should be set for the corresponding scoped terraform resource field matching the
			definition of a (Slice) of Swagger API Models (Struct).
			e.g: {"assigned_groups": BlockToStructSlice{}}
		 - BlockSliceToStructSlice - This is useful when mapping a list of terraform blocks (Type=TypeList) to a (Slice) of Swagger API Model (Struct).
			It follows the same rules of BlockToStructSlice but instead of creating a single block it will create a list of blocks.

*/
