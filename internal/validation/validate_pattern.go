package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"example.com/main-api-sdk/internal/utils"
)

// validatePattern validates that a string field matches the regex pattern specified in its 'pattern' tag.
// Returns an error if the field value doesn't match the pattern. Nil values are skipped.
func validatePattern(field reflect.StructField, value reflect.Value) error {
	pattern, found := field.Tag.Lookup("pattern")
	if !found {
		return nil
	}

	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("regex failed to compile")
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

	kind := utils.GetReflectKind(value.Type())
	if kind != reflect.String {
		return fmt.Errorf("field %s with value %v cannot match pattern %s because it is not a string", field.Name, value, pattern)
	}

	if !compiledRegex.MatchString(value.String()) {
		return fmt.Errorf("field %s with value %v does not match pattern %s", field.Name, value.String(), pattern)
	}

	return nil
}
