package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type ChatWithAssistantRequest struct {
	Message   *param.Nullable[string]   `json:"message,omitempty" xml:"message,omitempty"`
	SessionID *param.Nullable[string]   `json:"session_id,omitempty" xml:"session_id,omitempty"`
	Files     *param.Nullable[[]string] `json:"files,omitempty" xml:"files,omitempty"`
	Blocks    *param.Nullable[[]string] `json:"blocks,omitempty" xml:"blocks,omitempty"`
}

func (c ChatWithAssistantRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChatWithAssistantRequest to string"
	}
	return string(jsonData)
}

func (c *ChatWithAssistantRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
