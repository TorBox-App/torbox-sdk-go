package contenttypes

import (
	"example.com/main-api-sdk/internal/unmarshal"
	"fmt"
)

// FromJSON deserializes JSON data from HTTP response bodies into the target struct.
// Uses the custom unmarshal package to handle complex types and validations.
func FromJSON(data []byte, target any) error {
	err := unmarshal.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body into struct: %v", err)
	}
	return nil
}
