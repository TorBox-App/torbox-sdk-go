package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"torbox-sdk-go/internal/utils"
)

func validateArrayLength(field reflect.StructField, value reflect.Value) error {
	if value.IsNil() {
		return nil
	}

	kind := utils.GetReflectKind(value.Type())
	if kind != reflect.Array && kind != reflect.Slice {
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

func validateMinLength(field reflect.StructField, value reflect.Value) error {
	if value.IsNil() {
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

func validateMaxLength(field reflect.StructField, value reflect.Value) error {
	if value.IsNil() {
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
		return fmt.Errorf("too long")
	}

	return nil
}
