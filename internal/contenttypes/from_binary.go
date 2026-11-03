package contenttypes

import (
	"encoding/base64"
	"example.com/main-api-sdk/internal/unmarshal"
	"fmt"
	"reflect"
)

// FromBinary deserializes binary response data into a byte slice target.
// The target must be a non-nil pointer to []byte. Returns an error if conversion fails.
func FromBinary(data any, target any) error {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()

	if b, ok := data.([]byte); ok {
		if targetValue.Kind() == reflect.Slice && targetValue.Type().Elem().Kind() == reflect.Uint8 {
			targetValue.Set(reflect.ValueOf(b))
			return nil
		}

		if targetValue.Kind() == reflect.Interface {
			// Target is `any`/`interface{}` — store as a []byte.
			targetValue.Set(reflect.ValueOf(b))
			return nil
		}

		if targetValue.Kind() == reflect.Struct {
			base64Str := base64.StdEncoding.EncodeToString(b)
			jsonBody := []byte(`"` + base64Str + `"`)
			return unmarshal.Unmarshal(jsonBody, target)
		}
	}

	return fmt.Errorf("Failed to convert response body to byte array")
}
