package contenttypes

import (
	"net/url"
	"reflect"
	"strings"
)

func FromFormUrlEncoded(data []byte, target any) error {
	values, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	targetVal := reflect.ValueOf(target).Elem()
	targetType := targetVal.Type()

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldValue := targetVal.Field(i)

		key := getFieldName(field)

		if value, found := values[key]; found && len(value) > 0 {
			updateFieldWithValue(fieldValue, value)
		}
	}
	return nil
}

func updateFieldWithValue(fieldValue reflect.Value, value []string) {
	firstValue := value[0]
	if fieldValue.Kind() == reflect.Ptr {
		if fieldValue.IsNil() {
			newStr := firstValue
			fieldValue.Set(reflect.ValueOf(&newStr))
		} else {
			fieldValue.Elem().SetString(firstValue)
		}
	} else if fieldValue.CanSet() {
		fieldValue.SetString(firstValue) // Set only the first value if multiple are present
	}
}

func getFieldName(field reflect.StructField) string {
	key := field.Tag.Get("json")
	if key != "" {
		key = strings.Split(key, ",")[0]
	} else {
		key = field.Name
	}
	return key
}
