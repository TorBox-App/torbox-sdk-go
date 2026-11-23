package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type EditTagRequest struct {
	ID      *param.Nullable[float64] `json:"id,omitempty" xml:"id,omitempty"`
	NewName *param.Nullable[string]  `json:"new_name,omitempty" xml:"new_name,omitempty"`
}

func (e EditTagRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditTagRequest to string"
	}
	return string(jsonData)
}

func (e *EditTagRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}
