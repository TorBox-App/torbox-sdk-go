package utils

import "reflect"

func CloneMap[T any](sourceMap map[string]T) map[string]T {
	newMap := make(map[string]T)
	for key, value := range sourceMap {
		newMap[key] = value
	}

	return newMap
}

func GetReflectValueFromAny(input any) reflect.Value {
	val := reflect.ValueOf(input)
	return GetReflectValue(val)
}

func GetReflectValue(fieldValue reflect.Value) reflect.Value {
	if fieldValue.Kind() == reflect.Pointer {
		return fieldValue.Elem()
	} else {
		return fieldValue
	}
}

func GetReflectTypeFromAny(input any) reflect.Type {
	dataType := reflect.TypeOf(input)
	return GetReflectType(dataType)
}

func GetReflectType(fieldType reflect.Type) reflect.Type {
	if fieldType.Kind() == reflect.Ptr {
		return fieldType.Elem()
	} else {
		return fieldType
	}
}

func GetReflectKindFromAny(input any) reflect.Kind {
	dataType := reflect.TypeOf(input)
	return GetReflectKind(dataType)
}

func GetReflectKind(fieldType reflect.Type) reflect.Kind {
	if fieldType.Kind() == reflect.Pointer {
		return fieldType.Elem().Kind()
	} else {
		return fieldType.Kind()
	}
}
