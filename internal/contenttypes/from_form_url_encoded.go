package contenttypes

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// FromFormUrlEncoded deserializes URL-encoded form data into a struct.
// Uses struct field json tags to map form keys to struct fields. Only string fields are supported.
func FromFormUrlEncoded(data []byte, target any) error {
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() != reflect.Ptr || targetVal.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}
	targetVal = targetVal.Elem()

	if targetVal.Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	targetType := targetVal.Type()

	fieldsSet := 0
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldValue := targetVal.Field(i)

		key := getFieldName(field)

		if value, found := values[key]; found && len(value) > 0 {
			updateFieldWithValue(fieldValue, value)
			fieldsSet++
		}
	}

	if fieldsSet == 0 {
		return fmt.Errorf("no form fields were set - form data may be invalid or empty")
	}

	return nil
}

// updateFieldWithValue sets a struct field value from form data values.
// Handles both pointer and non-pointer string fields, using the first value if multiple are present.
func updateFieldWithValue(fieldValue reflect.Value, value []string) {
	firstValue := value[0]
	if fieldValue.Kind() == reflect.Ptr {
		if fieldValue.IsNil() {
			newStr := firstValue
			fieldValue.Set(reflect.ValueOf(&newStr))
		} else {
			fieldValue.Elem().SetString(firstValue)
		}
	} else if fieldValue.CanSet() {
		fieldValue.SetString(firstValue)
	}
}

// getFieldName extracts the field name from json tag or uses the struct field name as fallback.
// Returns the name to use as the form key for this field.
func getFieldName(field reflect.StructField) string {
	key := field.Tag.Get("json")
	if key != "" {
		key = strings.Split(key, ",")[0]
	} else {
		key = field.Name
	}
	return key
}
