package contenttypes

import (
	"fmt"
	"reflect"
)

func FromText[T any](data []byte, target any) error {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()

	switch targetValue.Kind() {
	case reflect.String:
		targetValue.SetString(string(data))
	default:
		return fmt.Errorf("unsupported target type: %s", targetValue.Kind())
	}

	return nil
}
