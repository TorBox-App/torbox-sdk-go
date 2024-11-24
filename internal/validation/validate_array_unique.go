package validation

import (
	"fmt"
	"reflect"
)

func validateArrayIsUnique(field reflect.StructField, value reflect.Value) error {
	unique, found := field.Tag.Lookup("uniqueItems")
	if !found || unique != "true" || value.IsNil() {
		return nil
	}

	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return nil
	}

	seen := make(map[any]bool)
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i).Interface()
		if seen[item] {
			return fmt.Errorf("the elements of this array must be unique, but this element appeared more than once: %v", item)
		}
		seen[item] = true
	}

	return nil
}
