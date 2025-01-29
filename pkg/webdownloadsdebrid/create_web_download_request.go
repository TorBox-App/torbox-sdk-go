package webdownloadsdebrid

import (
	"encoding/json"
)

type CreateWebDownloadRequest struct {
	// An accessible link to any file on the internet. Cannot be a redirection.
	Link    *string `json:"link,omitempty"`
	touched map[string]bool
}

func (c *CreateWebDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}

func (c *CreateWebDownloadRequest) SetLink(link string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Link"] = true
	c.Link = &link
}

func (c *CreateWebDownloadRequest) SetLinkNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Link"] = true
	c.Link = nil
}

func (c CreateWebDownloadRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Link"] && c.Link == nil {
		data["link"] = nil
	} else if c.Link != nil {
		data["link"] = c.Link
	}

	return json.Marshal(data)
}

func (c CreateWebDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadRequest to string"
	}
	return string(jsonData)
}
