package general

import (
	"encoding/json"
)

type GetStatsOkResponse struct {
	Data    *GetStatsOkResponseData `json:"data,omitempty"`
	Detail  *string                 `json:"detail,omitempty"`
	Error   *bool                   `json:"error,omitempty"`
	Success *bool                   `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetStatsOkResponse) GetData() *GetStatsOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetStatsOkResponse) SetData(data GetStatsOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = &data
}

func (g *GetStatsOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetStatsOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetStatsOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetStatsOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetStatsOkResponse) GetError() *bool {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetStatsOkResponse) SetError(error bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = &error
}

func (g *GetStatsOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetStatsOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetStatsOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetStatsOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetStatsOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Data"] && g.Data == nil {
		data["data"] = nil
	} else if g.Data != nil {
		data["data"] = g.Data
	}

	if g.touched["Detail"] && g.Detail == nil {
		data["detail"] = nil
	} else if g.Detail != nil {
		data["detail"] = g.Detail
	}

	if g.touched["Error"] && g.Error == nil {
		data["error"] = nil
	} else if g.Error != nil {
		data["error"] = g.Error
	}

	if g.touched["Success"] && g.Success == nil {
		data["success"] = nil
	} else if g.Success != nil {
		data["success"] = g.Success
	}

	return json.Marshal(data)
}

func (g GetStatsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetStatsOkResponse to string"
	}
	return string(jsonData)
}

type GetStatsOkResponseData struct {
	ActiveTorrents        *float64 `json:"active_torrents,omitempty"`
	ActiveUsenetDownloads *float64 `json:"active_usenet_downloads,omitempty"`
	ActiveWebDownloads    *float64 `json:"active_web_downloads,omitempty"`
	TotalBytesDownloaded  *float64 `json:"total_bytes_downloaded,omitempty"`
	TotalBytesUploaded    *float64 `json:"total_bytes_uploaded,omitempty"`
	TotalDownloads        *float64 `json:"total_downloads,omitempty"`
	TotalServers          *float64 `json:"total_servers,omitempty"`
	TotalTorrentDownloads *float64 `json:"total_torrent_downloads,omitempty"`
	TotalUsenetDownloads  *float64 `json:"total_usenet_downloads,omitempty"`
	TotalUsers            *float64 `json:"total_users,omitempty"`
	TotalWebDownloads     *float64 `json:"total_web_downloads,omitempty"`
	touched               map[string]bool
}

func (g *GetStatsOkResponseData) GetActiveTorrents() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveTorrents
}

func (g *GetStatsOkResponseData) SetActiveTorrents(activeTorrents float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveTorrents"] = true
	g.ActiveTorrents = &activeTorrents
}

func (g *GetStatsOkResponseData) SetActiveTorrentsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveTorrents"] = true
	g.ActiveTorrents = nil
}

func (g *GetStatsOkResponseData) GetActiveUsenetDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveUsenetDownloads
}

func (g *GetStatsOkResponseData) SetActiveUsenetDownloads(activeUsenetDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveUsenetDownloads"] = true
	g.ActiveUsenetDownloads = &activeUsenetDownloads
}

func (g *GetStatsOkResponseData) SetActiveUsenetDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveUsenetDownloads"] = true
	g.ActiveUsenetDownloads = nil
}

func (g *GetStatsOkResponseData) GetActiveWebDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveWebDownloads
}

func (g *GetStatsOkResponseData) SetActiveWebDownloads(activeWebDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveWebDownloads"] = true
	g.ActiveWebDownloads = &activeWebDownloads
}

func (g *GetStatsOkResponseData) SetActiveWebDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ActiveWebDownloads"] = true
	g.ActiveWebDownloads = nil
}

func (g *GetStatsOkResponseData) GetTotalBytesDownloaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalBytesDownloaded
}

func (g *GetStatsOkResponseData) SetTotalBytesDownloaded(totalBytesDownloaded float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalBytesDownloaded"] = true
	g.TotalBytesDownloaded = &totalBytesDownloaded
}

func (g *GetStatsOkResponseData) SetTotalBytesDownloadedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalBytesDownloaded"] = true
	g.TotalBytesDownloaded = nil
}

func (g *GetStatsOkResponseData) GetTotalBytesUploaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalBytesUploaded
}

func (g *GetStatsOkResponseData) SetTotalBytesUploaded(totalBytesUploaded float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalBytesUploaded"] = true
	g.TotalBytesUploaded = &totalBytesUploaded
}

func (g *GetStatsOkResponseData) SetTotalBytesUploadedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalBytesUploaded"] = true
	g.TotalBytesUploaded = nil
}

func (g *GetStatsOkResponseData) GetTotalDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalDownloads
}

func (g *GetStatsOkResponseData) SetTotalDownloads(totalDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalDownloads"] = true
	g.TotalDownloads = &totalDownloads
}

func (g *GetStatsOkResponseData) SetTotalDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalDownloads"] = true
	g.TotalDownloads = nil
}

func (g *GetStatsOkResponseData) GetTotalServers() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalServers
}

func (g *GetStatsOkResponseData) SetTotalServers(totalServers float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalServers"] = true
	g.TotalServers = &totalServers
}

func (g *GetStatsOkResponseData) SetTotalServersNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalServers"] = true
	g.TotalServers = nil
}

func (g *GetStatsOkResponseData) GetTotalTorrentDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalTorrentDownloads
}

func (g *GetStatsOkResponseData) SetTotalTorrentDownloads(totalTorrentDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalTorrentDownloads"] = true
	g.TotalTorrentDownloads = &totalTorrentDownloads
}

func (g *GetStatsOkResponseData) SetTotalTorrentDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalTorrentDownloads"] = true
	g.TotalTorrentDownloads = nil
}

func (g *GetStatsOkResponseData) GetTotalUsenetDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalUsenetDownloads
}

func (g *GetStatsOkResponseData) SetTotalUsenetDownloads(totalUsenetDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalUsenetDownloads"] = true
	g.TotalUsenetDownloads = &totalUsenetDownloads
}

func (g *GetStatsOkResponseData) SetTotalUsenetDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalUsenetDownloads"] = true
	g.TotalUsenetDownloads = nil
}

func (g *GetStatsOkResponseData) GetTotalUsers() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalUsers
}

func (g *GetStatsOkResponseData) SetTotalUsers(totalUsers float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalUsers"] = true
	g.TotalUsers = &totalUsers
}

func (g *GetStatsOkResponseData) SetTotalUsersNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalUsers"] = true
	g.TotalUsers = nil
}

func (g *GetStatsOkResponseData) GetTotalWebDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalWebDownloads
}

func (g *GetStatsOkResponseData) SetTotalWebDownloads(totalWebDownloads float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalWebDownloads"] = true
	g.TotalWebDownloads = &totalWebDownloads
}

func (g *GetStatsOkResponseData) SetTotalWebDownloadsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalWebDownloads"] = true
	g.TotalWebDownloads = nil
}

func (g GetStatsOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["ActiveTorrents"] && g.ActiveTorrents == nil {
		data["active_torrents"] = nil
	} else if g.ActiveTorrents != nil {
		data["active_torrents"] = g.ActiveTorrents
	}

	if g.touched["ActiveUsenetDownloads"] && g.ActiveUsenetDownloads == nil {
		data["active_usenet_downloads"] = nil
	} else if g.ActiveUsenetDownloads != nil {
		data["active_usenet_downloads"] = g.ActiveUsenetDownloads
	}

	if g.touched["ActiveWebDownloads"] && g.ActiveWebDownloads == nil {
		data["active_web_downloads"] = nil
	} else if g.ActiveWebDownloads != nil {
		data["active_web_downloads"] = g.ActiveWebDownloads
	}

	if g.touched["TotalBytesDownloaded"] && g.TotalBytesDownloaded == nil {
		data["total_bytes_downloaded"] = nil
	} else if g.TotalBytesDownloaded != nil {
		data["total_bytes_downloaded"] = g.TotalBytesDownloaded
	}

	if g.touched["TotalBytesUploaded"] && g.TotalBytesUploaded == nil {
		data["total_bytes_uploaded"] = nil
	} else if g.TotalBytesUploaded != nil {
		data["total_bytes_uploaded"] = g.TotalBytesUploaded
	}

	if g.touched["TotalDownloads"] && g.TotalDownloads == nil {
		data["total_downloads"] = nil
	} else if g.TotalDownloads != nil {
		data["total_downloads"] = g.TotalDownloads
	}

	if g.touched["TotalServers"] && g.TotalServers == nil {
		data["total_servers"] = nil
	} else if g.TotalServers != nil {
		data["total_servers"] = g.TotalServers
	}

	if g.touched["TotalTorrentDownloads"] && g.TotalTorrentDownloads == nil {
		data["total_torrent_downloads"] = nil
	} else if g.TotalTorrentDownloads != nil {
		data["total_torrent_downloads"] = g.TotalTorrentDownloads
	}

	if g.touched["TotalUsenetDownloads"] && g.TotalUsenetDownloads == nil {
		data["total_usenet_downloads"] = nil
	} else if g.TotalUsenetDownloads != nil {
		data["total_usenet_downloads"] = g.TotalUsenetDownloads
	}

	if g.touched["TotalUsers"] && g.TotalUsers == nil {
		data["total_users"] = nil
	} else if g.TotalUsers != nil {
		data["total_users"] = g.TotalUsers
	}

	if g.touched["TotalWebDownloads"] && g.TotalWebDownloads == nil {
		data["total_web_downloads"] = nil
	} else if g.TotalWebDownloads != nil {
		data["total_web_downloads"] = g.TotalWebDownloads
	}

	return json.Marshal(data)
}

func (g GetStatsOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetStatsOkResponseData to string"
	}
	return string(jsonData)
}
