package contenttypes

import (
	"bytes"
	"fmt"
)

// ToBinary serializes byte slice data for binary request bodies.
// Returns an error if the data is not a []byte.
func ToBinary(data any) (*bytes.Reader, error) {
	byteData, ok := data.([]byte)
	if !ok {
		return nil, fmt.Errorf("failed to serialize data to binary")
	}

	return bytes.NewReader(byteData), nil
}
