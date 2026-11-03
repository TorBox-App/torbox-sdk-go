package utils

import "reflect"

// CloneMap creates a shallow copy of a map.
// Returns a new map with the same key-value pairs as the source.
func CloneMap[T any](sourceMap map[string]T) map[string]T {
	newMap := make(map[string]T)
	for key, value := range sourceMap {
		newMap[key] = value
	}

	return newMap
}

// GetReflectValueFromAny gets the reflect.Value from any type, dereferencing pointers.
// Wraps GetReflectValue for convenience when starting with an any type.
func GetReflectValueFromAny(input any) reflect.Value {
	val := reflect.ValueOf(input)
	return GetReflectValue(val)
}

// GetReflectValue dereferences pointer types to get the underlying value.
// Returns the element if the value is a pointer, otherwise returns the value unchanged.
func GetReflectValue(fieldValue reflect.Value) reflect.Value {
	if fieldValue.Kind() == reflect.Pointer {
		return fieldValue.Elem()
	} else {
		return fieldValue
	}
}

// GetReflectTypeFromAny gets the reflect.Type from any type, dereferencing pointer types.
// Wraps GetReflectType for convenience when starting with an any type.
func GetReflectTypeFromAny(input any) reflect.Type {
	dataType := reflect.TypeOf(input)
	return GetReflectType(dataType)
}

// GetReflectType dereferences pointer types to get the underlying type.
// Returns the element type if the type is a pointer, otherwise returns the type unchanged.
func GetReflectType(fieldType reflect.Type) reflect.Type {
	if fieldType.Kind() == reflect.Ptr {
		return fieldType.Elem()
	} else {
		return fieldType
	}
}

// GetReflectKindFromAny gets the reflect.Kind from any type, dereferencing pointer types.
// Wraps GetReflectKind for convenience when starting with an any type.
func GetReflectKindFromAny(input any) reflect.Kind {
	dataType := reflect.TypeOf(input)
	return GetReflectKind(dataType)
}

// GetReflectKind dereferences pointer types to get the underlying kind.
// Returns the element kind if the type is a pointer, otherwise returns the kind unchanged.
func GetReflectKind(fieldType reflect.Type) reflect.Kind {
	if fieldType.Kind() == reflect.Pointer {
		return fieldType.Elem().Kind()
	} else {
		return fieldType.Kind()
	}
}

// IsNilable reports whether v is a kind that can be nil (pointer, interface, map, slice, chan, or func).
// Use this before calling v.IsNil() to avoid a reflect panic on value types.
func IsNilable(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return true
	default:
		return false
	}
}
