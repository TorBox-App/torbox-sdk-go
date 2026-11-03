package validation

import (
	"reflect"

	"example.com/main-api-sdk/internal/utils"
)

// validatorFunc defines the signature for field validation functions.
type validatorFunc = func(fieldType reflect.StructField, fieldValue reflect.Value) error

// ValidateData validates data structures recursively according to struct field tags.
// Supports required, pattern, min, max, multipleOf, unique array, and array length validations.
// Handles Nullable wrappers and nested structs/arrays.
// Note: required validation only applies to nilable types (pointers, slices, maps).
// Value types (string, int, bool, struct) are always considered set regardless of their value.
func ValidateData(data any) error {
	if data == nil {
		return nil
	}

	dataValue := reflect.ValueOf(data)

	// Recursively dereference pointers
	for dataValue.Kind() == reflect.Ptr {
		if dataValue.IsNil() {
			return nil
		}
		dataValue = dataValue.Elem()
	}

	dataType := dataValue.Type()

	// Check if this is a Nullable wrapper
	if isNullableType(dataValue.Type()) {
		if dataValue.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		dataValue = dataValue.FieldByName("Value")
		dataType = dataValue.Type()
	}

	if utils.GetReflectKind(dataType) == reflect.Struct {
		return validateStructValue(dataValue)
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		return validateArray(dataValue)
	}

	return nil
}

// validateStruct validates a struct by interface value (for exported fields).
// Recursively validates nested structs and arrays according to their validation tags.
func validateStruct(data any) error {
	value := reflect.ValueOf(data)
	return validateStructValue(value)
}

// validateStructValue validates a struct by reflect.Value (works with both exported and unexported fields)
func validateStructValue(value reflect.Value) error {

	// Recursively dereference pointers
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}

	// Check if this is a Nullable wrapper
	if isNullableType(value.Type()) {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper and restart validation from scratch
		// This ensures the wrapped value is validated according to its actual type
		wrappedValue := value.FieldByName("Value")

		if wrappedValue.CanInterface() {
			return ValidateData(wrappedValue.Interface())
		} else {
			// For unexported values, route to the appropriate validator based on type
			innerKind := utils.GetReflectKind(wrappedValue.Type())
			if innerKind == reflect.Struct {
				return validateStructValue(wrappedValue)
			} else if innerKind == reflect.Array || innerKind == reflect.Slice {
				return validateArray(wrappedValue)
			}
			return nil
		}
	}

	structValue := utils.GetReflectValue(value)

	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		// For unexported fields, skip field-level validation but still recurse into nested structs
		isUnexported := fieldType.PkgPath != ""

		if !isUnexported {
			err := validateField(fieldValue, fieldType)
			if err != nil {
				return err
			}
		}

		// Only check IsNil for types that can be nil
		kind := fieldValue.Kind()
		if (kind == reflect.Ptr || kind == reflect.Interface ||
			kind == reflect.Map || kind == reflect.Slice ||
			kind == reflect.Chan) && fieldValue.IsNil() {
			continue
		}

		// Recurse into nested structs/arrays/slices, even if the field is unexported
		kind = utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct {
			// For pointer fields, dereference first
			valueToValidate := fieldValue
			if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
				valueToValidate = fieldValue.Elem()
			}

			// Use validateStructValue to handle both exported and unexported fields
			if valueToValidate.Kind() == reflect.Struct {
				err := validateStructValue(valueToValidate)
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Array || kind == reflect.Slice {
			// For arrays/slices, validate using the existing validateArray function
			err := validateArray(fieldValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// validateArray validates each element in an array or slice recursively.
// Handles Nullable wrappers around arrays.
func validateArray(value reflect.Value) error {
	// Check if this is a Nullable wrapper
	if isNullableType(value.Type()) {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		value = value.FieldByName("Value")
	}

	arrayValue := utils.GetReflectValue(value)
	for j := 0; j < arrayValue.Len(); j++ {
		elementValue := arrayValue.Index(j)

		// Try to validate the element - use Interface() if possible, otherwise use reflect.Value directly
		if elementValue.CanInterface() {
			err := ValidateData(elementValue.Interface())
			if err != nil {
				return err
			}
		} else {
			// For unexported elements, validate using reflect.Value directly
			elementKind := utils.GetReflectKind(elementValue.Type())
			if elementKind == reflect.Struct {
				err := validateStructValue(elementValue)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// isNullableType checks if a type is a Nullable wrapper struct.
// Returns true if the type has IsNull (bool) and Value fields.
func isNullableType(t reflect.Type) bool {
	// Dereference pointer types until we hit the base type
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Must be a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	// Should have at least the fields "IsNull" (bool) and "Value" (any type)
	isNullField, hasIsNull := t.FieldByName("IsNull")
	_, hasValue := t.FieldByName("Value")

	if !hasIsNull || !hasValue {
		return false
	}

	// Check that IsNull is a bool
	if isNullField.Type.Kind() != reflect.Bool {
		return false
	}

	return true
}

// validateField runs all applicable validators on a struct field.
// Returns the first validation error encountered, or nil if all validations pass.
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

// getValidators returns the list of all validator functions to apply to a field.
// Includes required, pattern, multipleOf, min, max, array unique, and array length validators.
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
