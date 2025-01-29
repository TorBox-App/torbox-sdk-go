package torrents

import (
	"encoding/json"
)

type CreateTorrentRequest struct {
	// Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger.
	AllowZip *string `json:"allow_zip,omitempty"`
	// Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional.
	AsQueued *string `json:"as_queued,omitempty"`
	// The torrent's torrent file. Optional.
	File *any `json:"file,omitempty"`
	// The torrent's magnet link. Optional.
	Magnet *string `json:"magnet,omitempty"`
	// The name you want the torrent to be. Optional.
	Name *string `json:"name,omitempty"`
	// Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings.
	Seed    *string `json:"seed,omitempty"`
	touched map[string]bool
}

func (c *CreateTorrentRequest) GetAllowZip() *string {
	if c == nil {
		return nil
	}
	return c.AllowZip
}

func (c *CreateTorrentRequest) SetAllowZip(allowZip string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AllowZip"] = true
	c.AllowZip = &allowZip
}

func (c *CreateTorrentRequest) SetAllowZipNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AllowZip"] = true
	c.AllowZip = nil
}

func (c *CreateTorrentRequest) GetAsQueued() *string {
	if c == nil {
		return nil
	}
	return c.AsQueued
}

func (c *CreateTorrentRequest) SetAsQueued(asQueued string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AsQueued"] = true
	c.AsQueued = &asQueued
}

func (c *CreateTorrentRequest) SetAsQueuedNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AsQueued"] = true
	c.AsQueued = nil
}

func (c *CreateTorrentRequest) GetFile() *any {
	if c == nil {
		return nil
	}
	return c.File
}

func (c *CreateTorrentRequest) SetFile(file any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["File"] = true
	c.File = &file
}

func (c *CreateTorrentRequest) SetFileNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["File"] = true
	c.File = nil
}

func (c *CreateTorrentRequest) GetMagnet() *string {
	if c == nil {
		return nil
	}
	return c.Magnet
}

func (c *CreateTorrentRequest) SetMagnet(magnet string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Magnet"] = true
	c.Magnet = &magnet
}

func (c *CreateTorrentRequest) SetMagnetNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Magnet"] = true
	c.Magnet = nil
}

func (c *CreateTorrentRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTorrentRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateTorrentRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *CreateTorrentRequest) GetSeed() *string {
	if c == nil {
		return nil
	}
	return c.Seed
}

func (c *CreateTorrentRequest) SetSeed(seed string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Seed"] = true
	c.Seed = &seed
}

func (c *CreateTorrentRequest) SetSeedNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Seed"] = true
	c.Seed = nil
}

func (c CreateTorrentRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["AllowZip"] && c.AllowZip == nil {
		data["allow_zip"] = nil
	} else if c.AllowZip != nil {
		data["allow_zip"] = c.AllowZip
	}

	if c.touched["AsQueued"] && c.AsQueued == nil {
		data["as_queued"] = nil
	} else if c.AsQueued != nil {
		data["as_queued"] = c.AsQueued
	}

	if c.touched["File"] && c.File == nil {
		data["file"] = nil
	} else if c.File != nil {
		data["file"] = c.File
	}

	if c.touched["Magnet"] && c.Magnet == nil {
		data["magnet"] = nil
	} else if c.Magnet != nil {
		data["magnet"] = c.Magnet
	}

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["Seed"] && c.Seed == nil {
		data["seed"] = nil
	} else if c.Seed != nil {
		data["seed"] = c.Seed
	}

	return json.Marshal(data)
}

func (c CreateTorrentRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTorrentRequest to string"
	}
	return string(jsonData)
}
