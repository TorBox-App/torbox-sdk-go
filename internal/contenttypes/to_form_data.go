package contenttypes

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
)

// ToForm converts any value (struct, array, map, or primitive) to form data
func ToFormData(data interface{}) (*bytes.Reader, string, error) {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	defer writer.Close()

	err := encode("", reflect.ValueOf(data), writer)
	if err != nil {
		return nil, "", err
	}

	contentTypeHeader := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return bytes.NewReader(buffer.Bytes()), contentTypeHeader, nil
}

func encode(key string, v reflect.Value, writer *multipart.Writer) error {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		return encode(key, v.Elem(), writer)
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return writer.WriteField(key, string(v.Bytes()))
		}
		return fmt.Errorf("encoding error: only byte arrays/slices are supported")

	case reflect.Map:
		return fmt.Errorf("encoding error: maps are not supported")

	case reflect.Struct:
		return encodeStruct(key, v, writer)

	default:
		if key == "" {
			key = "value"
		}
		return writer.WriteField(key, formatValue(v))
	}
}

func encodeStruct(prefix string, v reflect.Value, writer *multipart.Writer) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		fieldValue := v.Field(i)
		fieldName := field.Name

		if tag, ok := field.Tag.Lookup("form"); ok {
			if tag == "-" {
				continue
			}
			fieldName = tag
		}

		if prefix != "" {
			fieldName = fmt.Sprintf("%s[%s]", prefix, fieldName)
		}

		if err := encode(fieldName, fieldValue, writer); err != nil {
			return err
		}
	}
	return nil
}

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		return v.String()
	default:
		return fmt.Sprint(v.Interface())
	}
}
