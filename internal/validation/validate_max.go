package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"example.com/main-api-sdk/internal/utils"
)

// validateMax validates that a numeric field value is less than or equal to the 'max' tag value.
// Supports int and float types. Nil values are skipped.
func validateMax(field reflect.StructField, value reflect.Value) error {
	maxValue, found := field.Tag.Lookup("max")
	if !found || maxValue == "" {
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

	max, err := strconv.Atoi(maxValue)
	if err != nil {
		return err
	}

	val := utils.GetReflectValue(value)

	if val.CanInt() {
		if val.Int() > int64(max) {
			return fmt.Errorf("validation Error. Field %s is greater than max value", field.Name)
		}
	} else if val.CanFloat() {
		if val.Float() > float64(max) {
			return fmt.Errorf("validation Error. Field %s is greater than max value", field.Name)
		}
	} else {
		return fmt.Errorf("validation Error. Field %s is not a number", field.Name)
	}

	return nil
}
