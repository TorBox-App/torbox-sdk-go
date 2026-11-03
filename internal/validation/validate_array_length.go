package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"example.com/main-api-sdk/internal/utils"
)

// validateArrayLength validates that an array or slice length is within 'minLength' and 'maxLength' tag constraints.
// Only validates array and slice types. Nil slices are skipped.
func validateArrayLength(field reflect.StructField, value reflect.Value) error {
	kind := utils.GetReflectKind(value.Type())
	if kind != reflect.Array && kind != reflect.Slice {
		return nil
	}

	if kind == reflect.Slice && value.IsNil() {
		return nil
	}

	err := validateMinLength(field, value)
	if err != nil {
		return err
	}

	err = validateMaxLength(field, value)
	if err != nil {
		return err
	}

	return nil
}

// validateMinLength validates that an array or slice meets the minimum length requirement from 'minLength' tag.
func validateMinLength(field reflect.StructField, value reflect.Value) error {
	kind := utils.GetReflectKind(value.Type())
	if kind == reflect.Slice && value.IsNil() {
		return nil
	}

	minLength, found := field.Tag.Lookup("minLength")
	if !found {
		return nil
	}

	minLengthInteger, err := strconv.Atoi(minLength)
	if err != nil {
		return err
	}

	actualLength := value.Len()
	if actualLength < minLengthInteger {
		return fmt.Errorf("the field myArray needs a minimum length of %v, but it currently has %v", minLengthInteger, actualLength)
	}

	return nil
}

// validateMaxLength validates that an array or slice doesn't exceed the maximum length from 'maxLength' tag.
func validateMaxLength(field reflect.StructField, value reflect.Value) error {
	kind := utils.GetReflectKind(value.Type())
	if kind == reflect.Slice && value.IsNil() {
		return nil
	}

	maxLength, found := field.Tag.Lookup("maxLength")
	if !found {
		return nil
	}

	maxLengthInteger, err := strconv.Atoi(maxLength)
	if err != nil {
		return err
	}

	if value.Len() > maxLengthInteger {
		return fmt.Errorf("the field myArray needs a maximum length of %v, but it currently has %v", maxLengthInteger, value.Len())
	}

	return nil
}
