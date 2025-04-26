package torrents

import "encoding/json"

type GetTorrentInfo1Request struct {
	// Hash of the torrent you want to get info for. This is required.
	Hash *string `json:"hash,omitempty"`
}

func (g *GetTorrentInfo1Request) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentInfo1Request) SetHash(hash string) {
	g.Hash = &hash
}

func (g GetTorrentInfo1Request) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfo1Request to string"
	}
	return string(jsonData)
}
