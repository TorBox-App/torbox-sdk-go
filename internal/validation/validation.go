package validation

import (
	"reflect"

	"torbox-sdk-go/internal/utils"
)

type validatorFunc = func(fieldType reflect.StructField, fieldValue reflect.Value) error

func ValidateData(data any) error {
	if data == nil {
		return nil
	}

	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	// Check if this is a Nullable wrapper
	if dataType.Name() == "Nullable" {
		if dataValue.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		dataValue = dataValue.FieldByName("Value")
		dataType = dataValue.Type()
	}

	if utils.GetReflectKind(dataType) == reflect.Struct {
		return validateStruct(dataValue.Interface())
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		return validateArray(dataValue)
	}

	return nil
}

func validateStruct(data any) error {
	value := reflect.ValueOf(data)

	// Check if this is a Nullable wrapper
	if value.Type().Name() == "Nullable" {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		value = value.FieldByName("Value")
	}

	structValue := utils.GetReflectValue(value)
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		err := validateField(fieldValue, fieldType)
		if err != nil {
			return err
		}

		// Only check IsNil for types that can be nil
		kind := fieldValue.Kind()
		if (kind == reflect.Ptr || kind == reflect.Interface ||
			kind == reflect.Map || kind == reflect.Slice ||
			kind == reflect.Chan) && fieldValue.IsNil() {
			continue
		}

		kind = utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice {
			err := ValidateData(fieldValue.Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateArray(value reflect.Value) error {
	// Check if this is a Nullable wrapper
	if value.Type().Name() == "Nullable" {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		value = value.FieldByName("Value")
	}

	arrayValue := utils.GetReflectValue(value)
	for j := 0; j < arrayValue.Len(); j++ {
		err := ValidateData(arrayValue.Index(j).Interface())
		if err != nil {
			return err
		}
	}

	return nil
}

func validateField(fieldValue reflect.Value, fieldType reflect.StructField) error {
	validators := getValidators(fieldType)
	for _, validator := range validators {
		err := validator(fieldType, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func getValidators(fieldType reflect.StructField) []validatorFunc {
	return []validatorFunc{
		validateRequired,
		validatePattern,
		validateMultipleOf,
		validateMin,
		validateMax,
		validateArrayIsUnique,
		validateArrayLength,
	}
}
