package validation

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"example.com/main-api-sdk/internal/utils"
)

// validateMultipleOf validates that a numeric field value is a multiple of the 'multipleOf' tag value.
// Supports int and float types. Nil values are skipped.
func validateMultipleOf(field reflect.StructField, value reflect.Value) error {
	multipleOfValue, found := field.Tag.Lookup("multipleOf")
	if !found || multipleOfValue == "" {
		return nil
	}

	// Unwrap pointers and nullable wrappers; skip null values.
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}
	if isNullableType(value.Type()) {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		value = value.FieldByName("Value")
	}

	multipleOf, err := strconv.Atoi(multipleOfValue)
	if err != nil {
		return err
	}

	val := utils.GetReflectValue(value)

	if val.CanInt() {
		if val.Int()%int64(multipleOf) != 0 {
			return fmt.Errorf("validation Error: Field %s must be a multiple of %v. Value: %v", field.Name, multipleOf, val)
		}
	} else if val.CanFloat() {
		if math.Mod(val.Float(), float64(multipleOf)) != 0 {
			return fmt.Errorf("validation Error: Field %s must be a multiple of %v. Value: %v", field.Name, multipleOf, val)
		}
	} else {
		return fmt.Errorf("validation Error: Field %s must a number. Value: %v", field.Name, val)
	}

	return nil
}
