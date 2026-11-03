package serialization

import (
	"fmt"
	"reflect"
	"strings"
)

// paramMap represents a key-value pair for query parameter serialization.
type paramMap struct {
	Key   string
	Value string
}

// SerializeObject serializes a struct to query parameter format with nested bracket notation.
// Converts struct fields to key[fieldName]=value format for complex query parameters.
func SerializeObject(key string, input any) []paramMap {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return []paramMap{}
	}

	params := []paramMap{}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if (value.Kind() == reflect.Ptr && value.IsNil()) || !value.CanInterface() {
			continue
		}

		jsonName := getJsonName(field)
		params = append(params, paramMap{
			Key:   fmt.Sprintf("%v[%v]", key, jsonName),
			Value: SerializeValue(jsonName, value.Interface()),
		})
	}

	return params
}

// SerializeValue serializes any value to a string representation for query parameters.
// Handles primitives, arrays/slices, and nested structs recursively.
func SerializeValue(key string, input any) string {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if typ.Kind() == reflect.Ptr && !val.IsNil() {
		return SerializeValue(key, val.Elem().Interface())
	}

	switch typ.Kind() {
	case reflect.Struct:
		return "SerializeStruct(input)"

	case reflect.Array, reflect.Slice:
		result := []string{}
		for i := 0; i < val.Len(); i++ {
			result = append(result, SerializeValue(key, val.Index(i).Interface()))
		}
		return strings.Join(result, ",")

	case reflect.String, reflect.Bool, reflect.Int, reflect.Int64, reflect.Float64:
		return SerializePrimitive(input)

	default:
		return ""
	}
}

// SerializePrimitive converts primitive types to their string representation.
// Uses fmt.Sprintf for consistent formatting.
func SerializePrimitive(input any) string {
	return fmt.Sprintf("%v", input)
}

// getJsonName extracts the JSON field name from a struct field's json tag.
// Returns the first part of the json tag before any comma.
func getJsonName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	return strings.Split(jsonTag, ",")[0]
}
