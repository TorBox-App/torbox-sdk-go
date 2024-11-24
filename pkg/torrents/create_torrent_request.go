package torrents

type CreateTorrentRequest struct {
	// The torrent's torrent file. Optional.
	File *any `json:"file,omitempty"`
	// The torrent's magnet link. Optional.
	Magnet *string `json:"magnet,omitempty"`
}

func (c *CreateTorrentRequest) SetFile(file any) {
	c.File = &file
}

func (c *CreateTorrentRequest) GetFile() *any {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateTorrentRequest) SetMagnet(magnet string) {
	c.Magnet = &magnet
}

func (c *CreateTorrentRequest) GetMagnet() *string {
	if c == nil {
		return nil
	}
	return c.Magnet
}
