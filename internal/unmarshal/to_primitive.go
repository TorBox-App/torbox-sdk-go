package unmarshal

import (
	"reflect"
	"strconv"
)

// ToString unmarshals bytes into a string target.
// Sets the string value directly on the target pointer element.
func ToString(source []byte, target reflect.Value) error {
	target.Elem().SetString(string(source))
	return nil
}

// ToInt unmarshals bytes into an int64 target by parsing the string representation.
// Returns an error if parsing fails.
func ToInt(source []byte, target reflect.Value) error {
	intBody, err := strconv.ParseInt(string(source), 10, 64)
	if err != nil {
		return err
	}

	target.Elem().SetInt(intBody)

	return nil
}

// ToFloat unmarshals bytes into a float64 target by parsing the string representation.
// Returns an error if parsing fails.
func ToFloat(source []byte, target reflect.Value) error {
	floatBody, err := strconv.ParseFloat(string(source), 64)
	if err != nil {
		return err
	}

	target.Elem().SetFloat(floatBody)

	return nil
}

// ToBool unmarshals bytes into a bool target by parsing the string representation.
// Returns an error if parsing fails.
func ToBool(source []byte, target reflect.Value) error {
	boolBody, err := strconv.ParseBool(string(source))
	if err != nil {
		return err
	}

	target.Elem().SetBool(boolBody)

	return nil
}
