package unmarshal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"example.com/main-api-sdk/internal/utils"
	"example.com/main-api-sdk/internal/validation"
)

// candidate represents a potential type match when deserializing a oneOf discriminated union.
// Tracks validation status and field counts to determine the best match.
type candidate struct {
	obj           any
	valid         bool
	requiredCount int
	optionalCount int
	kind          reflect.Kind
}

// ToComplexObject unmarshals JSON data into a discriminated union (oneOf) struct.
// Attempts to unmarshal into all possible types, validates each, and selects the best match
// based on required/optional field counts. Removes non-matching candidates from the result.
func ToComplexObject[T any](data []byte, result *T) error {
	err := unmarshalIntoProps(data, result)
	if err != nil {
		return err
	}

	candidates := createCandidatesFromProps(data, result)
	chosenCandidateIndex := chooseCandidateIndex(candidates)
	if chosenCandidateIndex == -1 {
		return errors.New("cannot unmarshal response, no valid candidate found")
	}

	removeOtherCandidates(result, chosenCandidateIndex)

	return nil
}

// unmarshalIntoProps attempts to unmarshal JSON data into all fields of a oneOf struct.
// Tries to parse the data as each possible type (struct, array, primitive) and sets matching fields.
func unmarshalIntoProps(data []byte, obj any) error {
	types := reflect.TypeOf(obj).Elem()
	values := reflect.ValueOf(obj).Elem()

	for i := 0; i < types.NumField(); i++ {
		fieldType := types.Field(i)
		kind := utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map {
			unmarshalledValue := reflect.New(fieldType.Type)
			err := json.Unmarshal(data, unmarshalledValue.Interface())
			if err != nil {
				continue
			}

			value := unmarshalledValue.Elem()
			values.Field(i).Set(value)
		} else if kind == reflect.String {
			strValue := string(data)
			values.Field(i).Set(reflect.ValueOf(&strValue))
		} else if kind == reflect.Float32 || kind == reflect.Float64 {
			value, err := strconv.ParseFloat(string(data), 64)
			if err == nil {
				values.Field(i).Set(reflect.ValueOf(&value))
			}
		} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
			value, err := strconv.ParseInt(string(data), 10, 64)
			if err == nil {
				values.Field(i).Set(reflect.ValueOf(&value))
			}
		} else if kind == reflect.Bool {
			value, err := strconv.ParseBool(string(data))
			if err == nil {
				values.Field(i).Set(reflect.ValueOf(&value))
			}
		} else if kind == reflect.Interface {
			values.Field(i).Set(reflect.ValueOf(string(data)))
		} else {
			return fmt.Errorf("cannot unmarshal response, unsupported type: %s", kind)
		}
	}

	return nil
}

// createCandidatesFromProps creates a candidate entry for each non-nil field in the oneOf struct.
// Validates each candidate and counts required/optional fields for ranking.
// Struct candidates that are missing required JSON keys are marked invalid and excluded from selection.
func createCandidatesFromProps(data []byte, obj any) []candidate {
	values := utils.GetReflectValue(reflect.ValueOf(obj))
	types := utils.GetReflectType(reflect.TypeOf(obj))

	candidates := make([]candidate, 0)
	for i := 0; i < types.NumField(); i++ {
		fieldValue := values.Field(i)
		kind := utils.GetReflectKind(types.Field(i).Type)

		var c candidate
		isNilable := utils.IsNilable(fieldValue)
		if isNilable && fieldValue.IsNil() {
			c = candidate{
				obj:           nil,
				valid:         false,
				requiredCount: 0,
				optionalCount: 0,
				kind:          kind,
			}
		} else if kind == reflect.Struct {
			value := fieldValue.Interface()
			c = candidate{
				obj:           value,
				valid:         isValid(value) && ValidateRequiredJSONKeys(data, value) == nil,
				requiredCount: countFields(value, validation.IsRequiredField),
				optionalCount: countFields(value, validation.IsOptionalField),
				kind:          kind,
			}
		} else if kind == reflect.Array || kind == reflect.Slice {
			value := fieldValue.Interface()
			c = candidate{
				obj:           value,
				valid:         isValid(value),
				requiredCount: countArrayFields(value, validation.IsRequiredField),
				optionalCount: countArrayFields(value, validation.IsOptionalField),
				kind:          kind,
			}
		} else {
			value := fieldValue.Interface()
			c = candidate{
				obj:           value,
				valid:         true,
				requiredCount: 0,
				optionalCount: 0,
				kind:          kind,
			}
		}

		candidates = append(candidates, c)
	}

	return candidates
}

// countFields recursively counts fields matching the predicate (required or optional) in a struct.
// Used to rank candidates by how many required/optional fields were successfully unmarshaled.
func countFields(c any, isFieldRequiredOrOptional func(reflect.StructField) bool) int {
	values := utils.GetReflectValue(reflect.ValueOf(c))
	types := utils.GetReflectType(reflect.TypeOf(c))

	if isPrimitive(utils.GetReflectKind(types)) {
		return 0
	}

	count := 0
	for i := 0; i < types.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldType := types.Field(i)

		isNilable := utils.IsNilable(fieldValue)
		if isNilable && fieldValue.IsNil() {
			continue
		}

		if isFieldRequiredOrOptional(fieldType) {
			count++
		}

		kind := utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice {
			count += countFields(fieldValue.Interface(), isFieldRequiredOrOptional)
		}
	}

	return count
}

// countArrayFields counts matching fields across all elements in an array.
func countArrayFields(candidates any, isFieldRequiredOrOptional func(reflect.StructField) bool) int {
	count := 0
	values := utils.GetReflectValue(reflect.ValueOf(candidates))
	for i := 0; i < values.Len(); i++ {
		candidate := values.Index(i).Interface()
		count += countFields(candidate, isFieldRequiredOrOptional)
	}

	return count
}

// isValid checks if a candidate passes all validation constraints.
func isValid(candidate any) bool {
	err := validation.ValidateData(candidate)
	return err == nil
}

// chooseCandidateIndex selects the best matching candidate from the list.
// Prefers non-primitive types first, then falls back to primitives in order: bool, int, float, string.
func chooseCandidateIndex(candidates []candidate) int {
	chosenCandidateIndex := chooseNonPrimitiveCandidate(candidates)

	if chosenCandidateIndex == -1 {
		chosenCandidateIndex = choosePrimitiveCandidate(candidates)
	}

	return chosenCandidateIndex
}

// chooseNonPrimitiveCandidate selects the best non-primitive (struct/array) candidate.
// Prioritizes candidates with more required fields, then more optional fields.
func chooseNonPrimitiveCandidate(candidates []candidate) int {
	chosenCandidateIndex := -1
	chosenCandidateRequiredCount := -1
	chosenCandidateOptionalCount := -1

	for i, candidate := range candidates {
		if isBetterCandidate(candidate, chosenCandidateRequiredCount, chosenCandidateOptionalCount) {
			chosenCandidateIndex = i
			chosenCandidateRequiredCount = candidate.requiredCount
			chosenCandidateOptionalCount = candidate.optionalCount
		}
	}

	return chosenCandidateIndex
}

// isBetterCandidate determines if a candidate is better than the currently chosen one.
// A candidate is better if it has more required fields, or equal required but more optional fields.
func isBetterCandidate(c candidate, chosenCandidateRequiredCount int, chosenCandidateOptionalCount int) bool {
	if !c.valid || isPrimitive(c.kind) {
		return false
	}

	if c.requiredCount > chosenCandidateRequiredCount {
		return true
	}

	if c.requiredCount == chosenCandidateRequiredCount && c.optionalCount > chosenCandidateOptionalCount {
		return true
	}

	return false
}

// choosePrimitiveCandidate selects the first non-nil primitive candidate in priority order.
// Priority: bool > int > float > string.
func choosePrimitiveCandidate(candidates []candidate) int {
	predicates := []func(kind reflect.Kind) bool{isBool, isInteger, isFloat, isString}

	for _, predicate := range predicates {
		chosenCandidateIndex := findFirstNonNil(candidates, predicate)
		if chosenCandidateIndex != -1 {
			return chosenCandidateIndex
		}
	}

	return -1
}

// removeOtherCandidates zeroes out all fields except the chosen candidate.
// Ensures only the matched oneOf variant remains in the result struct.
func removeOtherCandidates(obj any, chosenCandidateIndex int) {
	values := utils.GetReflectValue(reflect.ValueOf(obj))
	types := utils.GetReflectType(reflect.TypeOf(obj))

	for i := 0; i < types.NumField(); i++ {
		if i != chosenCandidateIndex {
			fieldValue := values.Field(i)
			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		}
	}
}

// findFirstNonNil finds the first candidate matching the kind predicate.
func findFirstNonNil(candidates []candidate, predicate func(kind reflect.Kind) bool) int {
	for i, c := range candidates {
		if c.obj != nil && predicate(c.kind) {
			return i
		}
	}
	return -1
}

func isPrimitive(kind reflect.Kind) bool {
	return isInteger(kind) || isString(kind) || isBool(kind) || isFloat(kind)
}

func isInteger(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

func isFloat(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

func isBool(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

func isString(kind reflect.Kind) bool {
	return kind == reflect.String
}
