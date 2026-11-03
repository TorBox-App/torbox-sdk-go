package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"example.com/main-api-sdk/internal/utils"
)

// validateMin validates that a numeric field value is greater than or equal to the 'min' tag value.
// Supports int and float types. Nil values are skipped.
func validateMin(field reflect.StructField, value reflect.Value) error {
	minValue, found := field.Tag.Lookup("min")
	if !found || minValue == "" {
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

	min, err := strconv.Atoi(minValue)
	if err != nil {
		return err
	}

	val := utils.GetReflectValue(value)

	if val.CanInt() {
		if val.Int() < int64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else if val.CanFloat() {
		if val.Float() < float64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else {
		return fmt.Errorf("field %s is not a number", field.Name)
	}

	return nil
}
