package unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ValidateRequiredJSONKeys checks that every struct field tagged required:"true" has a
// corresponding key present in the JSON object. Call this before unmarshaling to catch
// missing required fields early, before zero values are created.
func ValidateRequiredJSONKeys(source []byte, target any) error {
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(source, &rawMap); err != nil {
		// Not a JSON object — skip key validation (e.g. primitive or array).
		return nil
	}
	t := reflect.TypeOf(target)
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	return checkRequiredKeys(rawMap, t)
}

// UnmarshalNullable deserializes JSON into a struct with Nullable fields, correctly handling null values.
// Sets IsNull=true for fields that are explicitly null in the JSON, rather than unmarshaling them as zero values.
func UnmarshalNullable(source []byte, target any) error {
	// Use a temporary map to decode the raw JSON
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(source, &rawMap); err != nil {
		return err
	}

	if err := validateRequiredKeys(rawMap, target); err != nil {
		return err
	}

	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			// An unexported field (a privatized literal or stream flag) is set by the model's own code.
			continue
		}
		jsonKey := field.Tag.Get("json")

		split := strings.Split(jsonKey, ",")
		if len(split) > 0 {
			jsonKey = split[0]
		}

		if jsonKey == "" {
			jsonKey = field.Name
		}

		rawVal, ok := rawMap[jsonKey]
		if !ok {
			continue
		}

		fieldVal := val.Field(i)

		// Handle null explicitly
		if string(rawVal) == "null" && fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() != reflect.Struct {
			// A plain optional pointer stays nil on null.
			continue
		}
		if string(rawVal) == "null" && fieldVal.Kind() == reflect.Ptr {
			// Create a new instance of the pointer's element (e.g. Nullable[T])
			newNullable := reflect.New(fieldVal.Type().Elem()).Elem()

			// Set the IsNull field to true
			isNullField := newNullable.FieldByName("IsNull")
			if isNullField.IsValid() && isNullField.CanSet() {
				isNullField.SetBool(true)
			}

			// Assign back to the pointer field
			fieldVal.Set(newNullable.Addr())
		} else if fieldVal.CanAddr() {
			if err := json.Unmarshal(rawVal, fieldVal.Addr().Interface()); err != nil {
				return fmt.Errorf("failed to unmarshal field %s: %w", field.Name, err)
			}
		}
	}

	return nil
}

// validateRequiredKeys checks that every required:"true" struct field has its JSON key present.
// rawMap is the already-parsed JSON object map so we avoid double-parsing.
func validateRequiredKeys(rawMap map[string]json.RawMessage, target any) error {
	t := reflect.TypeOf(target)
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	return checkRequiredKeys(rawMap, t)
}

// checkRequiredKeys iterates the fields of t and returns an error for the first
// required:"true" field whose JSON key is absent from rawMap.
func checkRequiredKeys(rawMap map[string]json.RawMessage, t reflect.Type) error {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if r, ok := field.Tag.Lookup("required"); !ok || r != "true" {
			continue
		}
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			// Not part of the JSON (a folded path, query or header parameter), so it cannot be
			// required in it.
			continue
		}
		jsonKey := strings.SplitN(jsonTag, ",", 2)[0]
		if jsonKey == "" {
			jsonKey = field.Name
		}
		if _, ok := rawMap[jsonKey]; !ok {
			return fmt.Errorf("field %s is required", field.Name)
		}
	}
	return nil
}

// hasNullableFields checks if a struct contains any Nullable[T] pointer fields.
// Used to determine if special nullable unmarshaling logic should be applied.
func hasNullableFields(obj any) bool {
	t := reflect.TypeOf(obj)

	// Dereference pointer if needed
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Ensure it's a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// For pointer to Nullable[T]
		if field.Type.Kind() == reflect.Ptr {
			elem := field.Type.Elem()
			// check if the name starts with "Nullable"
			if strings.HasPrefix(elem.Name(), "Nullable") {
				return true
			}
		}
	}

	return false
}
