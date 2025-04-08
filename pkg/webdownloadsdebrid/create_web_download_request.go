package webdownloadsdebrid

import "encoding/json"

type CreateWebDownloadRequest struct {
	// An accessible link to any file on the internet. Cannot be a redirection.
	Link *string `json:"link,omitempty"`
}

func (c *CreateWebDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}

func (c *CreateWebDownloadRequest) SetLink(link string) {
	c.Link = &link
}

func (c CreateWebDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadRequest to string"
	}
	return string(jsonData)
}
