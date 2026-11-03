package marshal

import (
	"encoding/json"
	"errors"
	"reflect"

	"example.com/main-api-sdk/internal/utils"
)

// FromComplexObject marshals a discriminated union (oneOf) struct to JSON.
// Finds the first non-nil field (the active variant) and marshals only that field.
// Returns an error if all fields are nil.
func FromComplexObject(obj any) ([]byte, error) {
	types := utils.GetReflectType(reflect.TypeOf(obj))
	values := utils.GetReflectValue(reflect.ValueOf(obj))

	for i := 0; i < types.NumField(); i++ {
		if !values.Field(i).IsNil() {
			return json.Marshal(values.Field(i).Interface())
		}
	}

	return nil, errors.New("cannot marshal complex object, no non-nil fields found")
}
