package contenttypes

import (
	"fmt"
)

func FromBinary(data any, target any) error {
	if b, ok := data.([]byte); ok {
		target = &b
		return nil
	}

	return fmt.Errorf("Failed to convert response body to byte array")
}
