package contenttypes

import (
	"bytes"
	"example.com/main-api-sdk/internal/utils"
	"fmt"
	"net/url"
	"reflect"
)

// ToFormUrlEncoded serializes a struct to URL-encoded form format for request bodies.
// Uses struct field json tags as form keys. Supports string, int, float, and bool types.
func ToFormUrlEncoded(data any) (*bytes.Reader, error) {
	val := utils.GetReflectValueFromAny(data)

	dataType := utils.GetReflectTypeFromAny(data)
	if utils.GetReflectKindFromAny(dataType) != reflect.Struct {
		return nil, fmt.Errorf("FormUrlEncodingError: input must be a struct")
	}

	values := url.Values{}

	for i := 0; i < val.NumField(); i++ {
		field := dataType.Field(i)
		rawField := val.Field(i)

		if utils.IsNilable(rawField) && rawField.IsNil() {
			continue
		}

		fieldValue := utils.GetReflectValue(rawField)

		if !fieldValue.CanInterface() {
			continue
		}

		key := getFieldName(field)

		switch fieldValue.Kind() {
		case reflect.String:
			values.Set(key, fieldValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Set(key, fmt.Sprintf("%d", fieldValue.Int()))
		case reflect.Float32, reflect.Float64:
			values.Set(key, fmt.Sprintf("%f", fieldValue.Float()))
		case reflect.Bool:
			values.Set(key, fmt.Sprintf("%t", fieldValue.Bool()))
		default:
			values.Set(key, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}

	formString := values.Encode()
	return bytes.NewReader([]byte(formString)), nil
}
