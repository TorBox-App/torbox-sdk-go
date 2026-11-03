package unmarshal

import (
	"encoding/json"
)

// ToObject unmarshals JSON bytes into a struct or array target using standard json.Unmarshal.
// Handles regular objects and arrays that don't require special unmarshaling logic.
func ToObject(source []byte, target any) error {
	err := json.Unmarshal(source, target)
	if err != nil {
		return err
	}
	return nil
}
