package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type MainAPISDKRequestDownloadLinksPostRequest2 struct {
	FileIds *param.Nullable[[]float64] `json:"file_ids,omitempty" xml:"file_ids,omitempty"`
}

func (m MainAPISDKRequestDownloadLinksPostRequest2) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MainAPISDKRequestDownloadLinksPostRequest2 to string"
	}
	return string(jsonData)
}

func (m *MainAPISDKRequestDownloadLinksPostRequest2) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}
