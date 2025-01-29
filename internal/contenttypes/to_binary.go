package contenttypes

import (
	"bytes"
	"fmt"
)

func ToBinary(data any) (*bytes.Reader, error) {
	byteData, ok := data.([]byte)
	if !ok {
		return nil, fmt.Errorf("failed to serialize data to binary")
	}

	return bytes.NewReader(byteData), nil
}
