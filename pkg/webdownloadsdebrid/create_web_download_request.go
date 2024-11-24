package webdownloadsdebrid

type CreateWebDownloadRequest struct {
	// An accessible link to any file on the internet. Cannot be a redirection.
	Link *string `json:"link,omitempty"`
}

func (c *CreateWebDownloadRequest) SetLink(link string) {
	c.Link = &link
}

func (c *CreateWebDownloadRequest) GetLink() *string {
	if c == nil {
		return nil
	}
	return c.Link
}
