package validation

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"torbox-sdk-go/internal/utils"
)

func validateMultipleOf(field reflect.StructField, value reflect.Value) error {
	multipleOfValue, found := field.Tag.Lookup("multipleOf")
	if !found || multipleOfValue == "" || value.IsNil() {
		return nil
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
		if math.Mod(value.Float(), float64(multipleOf)) != 0 {
			return fmt.Errorf("validation Error: Field %s must be a multiple of %v. Value: %v", field.Name, multipleOf, val)
		}
	} else {
		return fmt.Errorf("validation Error: Field %s must a number. Value: %v", field.Name, val)
	}

	return nil
}
