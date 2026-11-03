package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type MagnetToFileRequest struct {
	Magnet *param.Nullable[string] `json:"magnet,omitempty" xml:"magnet,omitempty"`
}

func (m MagnetToFileRequest) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MagnetToFileRequest to string"
	}
	return string(jsonData)
}

func (m *MagnetToFileRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}
