package contenttypes

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
)

// ToFormData converts structs and primitives to multipart/form-data format for request bodies.
// Returns the encoded data, Content-Type header with boundary, and any error. Supports nested structs and byte arrays.
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

// encode recursively encodes a value into multipart form fields.
// Handles pointers, arrays/slices (byte arrays only), structs, and primitives.
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
			// Binary field: use CreateFormFile so Content-Disposition includes filename=
			fw, err := writer.CreateFormFile(key, key)
			if err != nil {
				return err
			}
			_, err = fw.Write(v.Bytes())
			return err
		}
		// Array of binary fields: each file becomes a separate named part.
		// An empty outer slice (v.Len() == 0) produces no parts — this is intentional
		// and differs from the []uint8 branch above, which always calls CreateFormFile
		// (producing a zero-byte part even for empty input).
		if v.Type().Elem().Kind() == reflect.Slice && v.Type().Elem().Elem().Kind() == reflect.Uint8 {
			for i := 0; i < v.Len(); i++ {
				partKey := fmt.Sprintf("%s[%d]", key, i)
				fw, err := writer.CreateFormFile(partKey, partKey)
				if err != nil {
					return err
				}
				_, err = fw.Write(v.Index(i).Bytes())
				if err != nil {
					return err
				}
			}
			return nil
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

// encodeStruct encodes all exported fields of a struct into multipart form fields.
// Uses 'form' tags for field names, supports nested bracket notation for nested structs.
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

// formatValue converts a primitive value to its string representation for form encoding.
// Supports bool, int, uint, float, and string types.
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
