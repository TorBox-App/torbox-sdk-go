package contenttypes

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"fmt"
	"reflect"
	"unicode/utf8"
)

// FromText deserializes plain text response data into a string target.
// The target must be a non-nil pointer to a string. Returns an error for unsupported types.
func FromText[T any](data []byte, target any) error {
	if !utf8.Valid(data) {
		return fmt.Errorf("data is not valid UTF-8 text")
	}

	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()

	switch targetValue.Kind() {
	case reflect.String:
		targetValue.SetString(string(data))
		return nil
	case reflect.Interface:
		// Target is `any`/`interface{}` — store as a plain string.
		targetValue.Set(reflect.ValueOf(string(data)))
		return nil
	case reflect.Struct:
		jsonBody, err := json.Marshal(string(data))
		if err != nil {
			return fmt.Errorf("failed to marshal text as JSON string: %v", err)
		}
		return unmarshal.Unmarshal(jsonBody, target)
	default:
		return fmt.Errorf("unsupported target type: %s", targetValue.Kind())
	}
}
