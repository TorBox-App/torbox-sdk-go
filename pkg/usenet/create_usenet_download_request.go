package usenet

import "encoding/json"

type CreateUsenetDownloadRequest struct {
	// An NZB File. Optional.
	File []byte `json:"file,omitempty"`
	// An accessible link to an NZB file. Cannot be a redirection. Optional.
	Link *string `json:"link,omitempty"`
}

func (c *CreateUsenetDownloadRequest) GetFile() []byte {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateUsenetDownloadRequest) SetFile(file []byte) {
	c.File = file
}

func (c *CreateUsenetDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}

func (c *CreateUsenetDownloadRequest) SetLink(link string) {
	c.Link = &link
}

func (c CreateUsenetDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadRequest to string"
	}
	return string(jsonData)
}
