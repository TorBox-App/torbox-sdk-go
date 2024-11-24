package contenttypes

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func FromFormData(formData []byte, dest interface{}) error {
	// Parse the form data into url.Values
	data, err := url.ParseQuery(string(formData))
	if err != nil {
		return fmt.Errorf("failed to parse form data: %v", err)
	}

	// Check if dest is a pointer to a struct
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.IsNil() || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}
	val = val.Elem()

	// Iterate over each field in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)

		// Use the 'form' tag as the form field key, or fall back to the field name
		formKey := field.Tag.Get("form")
		if formKey == "" {
			formKey = field.Name
		}

		// Get form value by key
		if formValues, ok := data[formKey]; ok && len(formValues) > 0 {
			formValue := formValues[0]

			// Set the value on the struct field based on its type
			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(formValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intVal, err := strconv.ParseInt(formValue, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to parse int for field %s: %v", field.Name, err)
				}
				fieldValue.SetInt(intVal)
			case reflect.Float32, reflect.Float64:
				floatVal, err := strconv.ParseFloat(formValue, 64)
				if err != nil {
					return fmt.Errorf("failed to parse float for field %s: %v", field.Name, err)
				}
				fieldValue.SetFloat(floatVal)
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(formValue)
				if err != nil {
					return fmt.Errorf("failed to parse bool for field %s: %v", field.Name, err)
				}
				fieldValue.SetBool(boolVal)
			default:
				return fmt.Errorf("unsupported field type for field %s", field.Name)
			}
		}
	}
	return nil
}
