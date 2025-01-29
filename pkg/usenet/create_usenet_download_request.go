package usenet

import (
	"encoding/json"
)

type CreateUsenetDownloadRequest struct {
	// An NZB File. Optional.
	File *any `json:"file,omitempty"`
	// An accessible link to an NZB file. Cannot be a redirection. Optional.
	Link    *string `json:"link,omitempty"`
	touched map[string]bool
}

func (c *CreateUsenetDownloadRequest) GetFile() *any {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateUsenetDownloadRequest) SetFile(file any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["File"] = true
	c.File = &file
}

func (c *CreateUsenetDownloadRequest) SetFileNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["File"] = true
	c.File = nil
}

func (c *CreateUsenetDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}

func (c *CreateUsenetDownloadRequest) SetLink(link string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Link"] = true
	c.Link = &link
}

func (c *CreateUsenetDownloadRequest) SetLinkNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Link"] = true
	c.Link = nil
}

func (c CreateUsenetDownloadRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["File"] && c.File == nil {
		data["file"] = nil
	} else if c.File != nil {
		data["file"] = c.File
	}

	if c.touched["Link"] && c.Link == nil {
		data["link"] = nil
	} else if c.Link != nil {
		data["link"] = c.Link
	}

	return json.Marshal(data)
}

func (c CreateUsenetDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadRequest to string"
	}
	return string(jsonData)
}
