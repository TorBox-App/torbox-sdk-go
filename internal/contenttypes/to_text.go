package contenttypes

import (
	"bytes"
	"fmt"
)

func ToText(data any) (*bytes.Reader, error) {
	switch v := data.(type) {
	case string:
		return bytes.NewReader([]byte(v)), nil
	default:
		return nil, fmt.Errorf("failed to serialize text data to binary")
	}
}
