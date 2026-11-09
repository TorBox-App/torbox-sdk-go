package mainapisdk1

import (
	"example.com/main-api-sdk/param"
)

// RequestDownloadLinkRequestParams holds the optional parameters for the API request.
type RequestDownloadLinkRequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	TorrentID  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	FileID     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"redirect"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksGetRequestParams holds the optional parameters for the API request.
type RequestDownloadLinksGetRequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	TorrentID  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	FileIds    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_ids"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksPostRequestParams holds the optional parameters for the API request.
type RequestDownloadLinksPostRequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	TorrentID  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// GetTorrentListRequestParams holds the optional parameters for the API request.
type GetTorrentListRequestParams struct {
	BypassCache *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	ID          *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	Active      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"active"`
	Files       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"files"`
}

// GetTorrentCachedAvailabilityRequestParams holds the optional parameters for the API request.
type GetTorrentCachedAvailabilityRequestParams struct {
	Hash      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"format"`
	ListFiles *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"list_files"`
}

// ExportTorrentDataRequestParams holds the optional parameters for the API request.
type ExportTorrentDataRequestParams struct {
	TorrentID *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	Type      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"type"`
}

// GetTorrentInfoRequestParams holds the optional parameters for the API request.
type GetTorrentInfoRequestParams struct {
	Hash           *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Timeout        *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"timeout"`
	UseCacheLookup *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"use_cache_lookup"`
}

// RequestDownloadLink1RequestParams holds the optional parameters for the API request.
type RequestDownloadLink1RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	UsenetID   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"usenet_id"`
	FileID     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"redirect"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksGet1RequestParams holds the optional parameters for the API request.
type RequestDownloadLinksGet1RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	UsenetID   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"usenet_id"`
	FileIds    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_ids"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksPost1RequestParams holds the optional parameters for the API request.
type RequestDownloadLinksPost1RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	UsenetID   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"usenet_id"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// GetUsenetListRequestParams holds the optional parameters for the API request.
type GetUsenetListRequestParams struct {
	BypassCache *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	ID          *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	Active      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"active"`
	Files       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"files"`
}

// GetUsenetCachedAvailabilityRequestParams holds the optional parameters for the API request.
type GetUsenetCachedAvailabilityRequestParams struct {
	Hash      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"format"`
	ListFiles *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"list_files"`
}

// RequestDownloadLink2RequestParams holds the optional parameters for the API request.
type RequestDownloadLink2RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	WebID      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"web_id"`
	FileID     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"redirect"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksGet2RequestParams holds the optional parameters for the API request.
type RequestDownloadLinksGet2RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	WebID      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"web_id"`
	FileIds    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_ids"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// RequestDownloadLinksPost2RequestParams holds the optional parameters for the API request.
type RequestDownloadLinksPost2RequestParams struct {
	Token      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	WebID      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"web_id"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	AppendName *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"append_name"`
}

// GetWebDownloadListRequestParams holds the optional parameters for the API request.
type GetWebDownloadListRequestParams struct {
	BypassCache *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	ID          *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	Active      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"active"`
	Files       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"files"`
}

// GetWebDownloadCachedAvailabilityRequestParams holds the optional parameters for the API request.
type GetWebDownloadCachedAvailabilityRequestParams struct {
	Hash      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"format"`
	ListFiles *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"list_files"`
}

// GetWebDownloadCachedAvailabilityBatchRequestParams holds the optional parameters for the API request.
type GetWebDownloadCachedAvailabilityBatchRequestParams struct {
	Format    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"format"`
	ListFiles *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"list_files"`
}

// GetChangelogsRssFeedRequestParams holds the optional parameters for the API request.
type GetChangelogsRssFeedRequestParams struct {
	Offset *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
}

// GetChangelogsJSONRequestParams holds the optional parameters for the API request.
type GetChangelogsJSONRequestParams struct {
	Offset *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
}

// GetSpeedtestFilesRequestParams holds the optional parameters for the API request.
type GetSpeedtestFilesRequestParams struct {
	TestLength *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"test_length"`
	Region     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"region"`
	UserIP     *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
}

// GetRssNotificationFeedRequestParams holds the optional parameters for the API request.
type GetRssNotificationFeedRequestParams struct {
	Token *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
}

// GetUserDataRequestParams holds the optional parameters for the API request.
type GetUserDataRequestParams struct {
	Settings *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"settings"`
}

// AddReferralToAccountRequestParams holds the optional parameters for the API request.
type AddReferralToAccountRequestParams struct {
	Referral *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"referral"`
}

// StartDeviceCodeAuthorizationProcessRequestParams holds the optional parameters for the API request.
type StartDeviceCodeAuthorizationProcessRequestParams struct {
	App *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"app"`
}

// GetTransactionPdfRequestParams holds the optional parameters for the API request.
type GetTransactionPdfRequestParams struct {
	TransactionID *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"transaction_id"`
}

// GetUserStatsRequestParams holds the optional parameters for the API request.
type GetUserStatsRequestParams struct {
	General           *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"general"`
	Bandwidth         *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bandwidth"`
	BandwidthGrouping *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bandwidth_grouping"`
}

// GetUserRssFeedsRequestParams holds the optional parameters for the API request.
type GetUserRssFeedsRequestParams struct {
	ID *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
}

// GetRssFeedItemsRequestParams holds the optional parameters for the API request.
type GetRssFeedItemsRequestParams struct {
	RssFeedID *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"rss_feed_id"`
}

// GetQueuedDownloadsRequestParams holds the optional parameters for the API request.
type GetQueuedDownloadsRequestParams struct {
	BypassCache *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	ID          *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	Type        *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"type"`
}

// CreateStreamRequestParams holds the optional parameters for the API request.
type CreateStreamRequestParams struct {
	ID                    *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"id"`
	FileID                *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	Type                  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"type"`
	ChosenSubtitleIndex   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_subtitle_index"`
	ChosenAudioIndex      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_audio_index"`
	ChosenResolutionIndex *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_resolution_index"`
}

// GetStreamDataRequestParams holds the optional parameters for the API request.
type GetStreamDataRequestParams struct {
	PresignedToken        *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"presigned_token"`
	Token                 *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"token"`
	ChosenSubtitleIndex   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_subtitle_index"`
	ChosenAudioIndex      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_audio_index"`
	ChosenResolutionIndex *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"chosen_resolution_index"`
}
