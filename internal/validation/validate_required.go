package validation

import (
	"fmt"
	"reflect"

	"example.com/main-api-sdk/internal/utils"
)

// validateRequired checks if a required nilable field has been set.
// For nilable types (pointer, interface, map, slice, etc.) it checks for nil.
// Value types (string, int, bool, struct) are always considered set — Go initialises
// them to their zero value, and 0 / false / "" are all legitimate required values.
func validateRequired(fieldType reflect.StructField, fieldValue reflect.Value) error {
	if !IsRequiredField(fieldType) {
		return nil
	}
	if utils.IsNilable(fieldValue) && fieldValue.IsNil() {
		return fmt.Errorf("field %s is required", fieldType.Name)
	}
	return nil
}

// IsRequiredField checks if a struct field has the 'required:"true"' tag.
func IsRequiredField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return found && required == "true"
}

// IsOptionalField checks if a struct field is optional (not required or no tag).
func IsOptionalField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return !found || required == "" || required == "false"
}
