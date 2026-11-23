package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type ControlTagRequest struct {
	ID        *param.Nullable[float64] `json:"id,omitempty" xml:"id,omitempty"`
	Operation *param.Nullable[string]  `json:"operation,omitempty" xml:"operation,omitempty"`
}

func (c ControlTagRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ControlTagRequest to string"
	}
	return string(jsonData)
}

func (c *ControlTagRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
