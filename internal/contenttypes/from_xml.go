package contenttypes

import (
	"encoding/xml"
	"fmt"
)

// FromXML deserializes XML data from HTTP response bodies into the target struct.
// Uses the standard encoding/xml package for unmarshaling.
func FromXML(data []byte, target any) error {
	err := xml.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal XML response body into struct: %w", err)
	}
	return nil
}
