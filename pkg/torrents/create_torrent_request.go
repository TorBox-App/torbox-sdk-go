package torrents

import "encoding/json"

type CreateTorrentRequest struct {
	// Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger.
	AllowZip *string `json:"allow_zip,omitempty"`
	// Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional.
	AsQueued *string `json:"as_queued,omitempty"`
	// The torrent's torrent file. Optional.
	File []byte `json:"file,omitempty"`
	// The torrent's magnet link. Optional.
	Magnet *string `json:"magnet,omitempty"`
	// The name you want the torrent to be. Optional.
	Name *string `json:"name,omitempty"`
	// Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings.
	Seed *string `json:"seed,omitempty"`
}

func (c *CreateTorrentRequest) GetAllowZip() *string {
	if c == nil {
		return nil
	}
	return c.AllowZip
}

func (c *CreateTorrentRequest) SetAllowZip(allowZip string) {
	c.AllowZip = &allowZip
}

func (c *CreateTorrentRequest) GetAsQueued() *string {
	if c == nil {
		return nil
	}
	return c.AsQueued
}

func (c *CreateTorrentRequest) SetAsQueued(asQueued string) {
	c.AsQueued = &asQueued
}

func (c *CreateTorrentRequest) GetFile() []byte {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateTorrentRequest) SetFile(file []byte) {
	c.File = file
}

func (c *CreateTorrentRequest) GetMagnet() *string {
	if c == nil {
		return nil
	}
	return c.Magnet
}

func (c *CreateTorrentRequest) SetMagnet(magnet string) {
	c.Magnet = &magnet
}

func (c *CreateTorrentRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTorrentRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateTorrentRequest) GetSeed() *string {
	if c == nil {
		return nil
	}
	return c.Seed
}

func (c *CreateTorrentRequest) SetSeed(seed string) {
	c.Seed = &seed
}

func (c CreateTorrentRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTorrentRequest to string"
	}
	return string(jsonData)
}
