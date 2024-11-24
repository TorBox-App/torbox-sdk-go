package contenttypes

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"torbox-sdk-go/internal/utils"
)

func ToFormUrlEncoded(data any) (*bytes.Reader, error) {
	val := utils.GetReflectValueFromAny(data)

	dataType := utils.GetReflectTypeFromAny(data)
	if utils.GetReflectKindFromAny(dataType) != reflect.Struct {
		return nil, fmt.Errorf("FormUrlEncodingError: input must be a struct")
	}

	values := url.Values{}

	for i := 0; i < val.NumField(); i++ {
		field := dataType.Field(i)
		fieldValue := utils.GetReflectValue(val.Field(i))

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
