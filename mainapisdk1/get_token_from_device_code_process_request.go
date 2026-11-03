package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type GetTokenFromDeviceCodeProcessRequest struct {
	DeviceCode *param.Nullable[string] `json:"device_code,omitempty" xml:"device_code,omitempty"`
}

func (g GetTokenFromDeviceCodeProcessRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTokenFromDeviceCodeProcessRequest to string"
	}
	return string(jsonData)
}

func (g *GetTokenFromDeviceCodeProcessRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
