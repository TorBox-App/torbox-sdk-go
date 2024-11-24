package contenttypes

import (
	"bytes"
	"encoding/json"
)

func ToJson(data any) (*bytes.Reader, error) {
	if data == nil {
		return nil, nil
	}

	marshalledBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(marshalledBody), nil
}
