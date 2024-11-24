package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"torbox-sdk-go/internal/utils"
)

func validateMax(field reflect.StructField, value reflect.Value) error {
	maxValue, found := field.Tag.Lookup("max")
	if !found || maxValue == "" || value.IsNil() {
		return nil
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
