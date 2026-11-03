package contenttypes

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// ToXML serializes data to XML format for HTTP request bodies.
// Returns a bytes.Reader containing the XML-encoded data, or nil if data is nil.
func ToXML(data any) (*bytes.Reader, error) {
	if data == nil {
		return nil, nil
	}

	marshalledBody, err := xml.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal XML request body: %w", err)
	}

	return bytes.NewReader(marshalledBody), nil
}
