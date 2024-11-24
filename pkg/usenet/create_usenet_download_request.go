package usenet

type CreateUsenetDownloadRequest struct {
	// An NZB File. Optional.
	File *any `json:"file,omitempty"`
	// An accessible link to an NZB file. Cannot be a redirection. Optional.
	Link *string `json:"link,omitempty"`
}

func (c *CreateUsenetDownloadRequest) SetFile(file any) {
	c.File = &file
}

func (c *CreateUsenetDownloadRequest) GetFile() *any {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateUsenetDownloadRequest) SetLink(link string) {
	c.Link = &link
}

func (c *CreateUsenetDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}
