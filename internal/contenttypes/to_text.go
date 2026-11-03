package contenttypes

import (
	"bytes"
	"fmt"
)

// ToText serializes string data to bytes for text/plain request bodies.
// Returns an error if the data is not a string.
func ToText(data any) (*bytes.Reader, error) {
	switch v := data.(type) {
	case string:
		return bytes.NewReader([]byte(v)), nil
	default:
		return nil, fmt.Errorf("failed to serialize text data to binary")
	}
}
