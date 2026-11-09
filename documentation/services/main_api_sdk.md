# MainAPISDK1

A list of all methods in the `MainAPISDK1` service. Click on the method name to view detailed information about that method.

| Methods                                                                         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| :------------------------------------------------------------------------------ | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreateTorrent](#createtorrent)                                                 | ### Overview Creates a torrent under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. ### Rate Limits This endpoint has two limits that apply at the same time: - **300 requests per minute:** a shared ceiling that _every_ request counts against, whether the torrent is cached or uncached. - **60 uncached torrents per hour:** an additional, stricter limit that only uncached torrents count against. The important thing is that cached and uncached requests both draw from the same 300/min pool, but uncached requests _also_ draw from their own separate 60/hour pool. An uncached request therefore counts twice: once against the per-minute limit, and once against the per-hour uncached limit. **Example:** In a single minute you add 250 cached + 50 uncached torrents. That's 300 total, so you've hit the per-minute limit for that minute. But those 50 uncached torrents also came out of your 60/hour uncached budget, leaving only **10 uncached torrents** for the remaining 59 minutes of that hour. (You could still add plenty more _cached_ torrents during that hour, since cached requests aren't capped hourly: only by the 300/min ceiling.) ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| [AsyncCreateTorrent](#asynccreatetorrent)                                       | ### Overview Creates a torrent asynchrously under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. The request will return instantly, not returning any data other than a success message. This is the same for failures, as they will return the error in the background, rather than immediately. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [ControlTorrent](#controltorrent)                                               | ### Overview Controls a torrent. By sending the torrent's ID and the type of operation you want to perform, it will send that request to the torrent client. Operations are either: - **Reannounce** `reannounces the torrent to get new peers` - **Delete** `deletes the torrent from the client and your account permanently` - **Resume** `resumes a paused torrent` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [RequestDownloadLink](#requestdownloadlink)                                     | ### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [RequestDownloadLinksGet](#requestdownloadlinksget)                             | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [RequestDownloadLinksPost](#requestdownloadlinkspost)                           | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetTorrentList](#gettorrentlist)                                               | ### Overview Gets the user's torrent list. This gives you the needed information to perform other torrent actions. This information only gets updated every 600 seconds, or when the _Request Update On Torrent_ request is sent to the _relay API_. #### Download States: - "downloading" -\> The torrent is currently downloading. - "uploading" -\> The torrent is currently seeding. - "stalled (no seeds)" -\> The torrent is trying to download, but there are no seeds connected to download from. - "paused" -\> The torrent is paused. - "completed" -\> The torrent is completely downloaded. Do not use this for download completion status. - "cached" -\> The torrent is cached from the server. - "metaDL" -\> The torrent is downloading metadata from the hoard. - "checkingResumeData" -\> The torrent is checking resumable data. All other statuses are basic qBittorrent states. [Check out here for the full list](<https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#torrent-management>). ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [GetTorrentCachedAvailability](#gettorrentcachedavailability)                   | ### Overview Takes in a list of comma separated torrent hashes and checks if the torrent is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [GetTorrentCachedAvailabilityBatch](#gettorrentcachedavailabilitybatch)         | ### Overview Takes in a list of comma separated torrent hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [ExportTorrentData](#exporttorrentdata)                                         | ### Overview Exports the magnet or torrent file. Requires a type to be passed. If type is **magnet**, it will return a JSON response with the magnet as a string in the _data_ key. If type is **file**, it will return a bittorrent file as a download. Not compatible with cached downloads. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetTorrentInfo](#gettorrentinfo)                                               | ### Overview A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant. **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.** ### Authorization None required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [GetTorrentInfo1](#gettorrentinfo1)                                             | ### Overview Same as the GET route, but allows posting data such as magnet, and torrent files. Hashes will have precedence over magnets, and magnets will have precedence over torrent files. Only proper torrent files are accepted. At least one of hash, magnet, or torrent file is required. A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant. **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.** ### Authorization None required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| [EditTorrentItem](#edittorrentitem)                                             | ### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [MagnetToFile](#magnettofile)                                                   | ### Overview Converts any magnet to a torrent file. Does not require the magnet to be on your TorBox account or in the TorBox cache. Free endpoint for public use. ### Authorization None required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [CreateUsenetDownload](#createusenetdownload)                                   | ### Overview Creates a usenet download under your account. Simply send **either** a link, or an nzb file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. #### Post Processing Options: All post processing options that the Usenet client will perform before TorBox's own processing to make the files available. It is recommended you either don't send this parameter, or keep it at `-1` for default, which will give only the wanted files. - `-1`, Default. This runs repairs, and extractions as well as deletes the source files leaving only the wanted downloaded files. - `0`, None. No post-processing at all. Just download all the files (including all PAR2). TorBox will still do its normal processing to make the download available, but no repairs, or extraction will take place. - `1`, Repair. Download files and do a PAR2 verification. If the verification fails, download more PAR2 files and attempt to repair the files. - `2`, Repair and Unpack. Download all files, do a PAR2 verification and unpack the files. The final folder will also include the RAR and ZIP files. - `3`, Repair, Unpack and Delete. Download all files, do a PAR2 verification, unpack the files to the final folder and delete the source files. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| [ControlUsenetDownload](#controlusenetdownload)                                 | ### Overview Controls a usenet download. By sending the usenet download's ID and the type of operation you want to perform, it will send that request to the usenet client. Operations are either: - **Delete** `deletes the nzb from the client and your account permanently` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [RequestDownloadLink1](#requestdownloadlink1)                                   | ### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [RequestDownloadLinksGet1](#requestdownloadlinksget1)                           | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [RequestDownloadLinksPost1](#requestdownloadlinkspost1)                         | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetUsenetList](#getusenetlist)                                                 | ### Overview Gets the user's usenet download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live usenet downloads. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [GetUsenetCachedAvailability](#getusenetcachedavailability)                     | ### Overview Takes in a list of comma separated usenet hashes and checks if the usenet download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a usenet download, you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetUsenetCachedAvailabilityBatch](#getusenetcachedavailabilitybatch)           | ### Overview Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [EditUsenetItem](#editusenetitem)                                               | ### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [GetNewsServerConnectionInformation](#getnewsserverconnectioninformation)       | ### Overview Gets the TorBox News Server Connection information. This is a static path, and is not dynamic. ### Authorization No authentication required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [GetNewsServerAccount](#getnewsserveraccount)                                   | ### Overview Gets the user's TorBox News Server Account with authentication details. Does not show password if not creating the account. Account creation happens on the first request. If you need to get the password, or reset the password, use the `/usenet/provider/account/resetpw` `POST` endpoint. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| [ResetNewsServerAccountPassword](#resetnewsserveraccountpassword)               | ### Overview Resets the user's TorBox News Server account password. This password is automatically generated, and is not saved, it is only shown once. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [CreateWebDownload](#createwebdownload)                                         | ### Overview Creates a web download under your account. Simply send a link to any file on the internet. Once it has been checked, it will begin downloading assuming your account has available active download slots, and they aren't too large. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [ControlWebDownload](#controlwebdownload)                                       | ### Overview Controls a web download. By sending the web download's ID and the type of operation you want to perform, it will send that request to the debrid client. Operations are either: - **Delete** `deletes the download from the client and your account permanently` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [RequestDownloadLink2](#requestdownloadlink2)                                   | ### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [RequestDownloadLinksGet2](#requestdownloadlinksget2)                           | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [RequestDownloadLinksPost2](#requestdownloadlinkspost2)                         | ### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetWebDownloadList](#getwebdownloadlist)                                       | ### Overview Gets the user's web download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live web downloads. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [GetWebDownloadCachedAvailability](#getwebdownloadcachedavailability)           | ### Overview Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a web download you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [GetWebDownloadCachedAvailabilityBatch](#getwebdownloadcachedavailabilitybatch) | ### Overview Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [GetHosterList](#gethosterlist)                                                 | ### Overview A dynamic list of hosters that TorBox is capable of downloading through its paid service. - **Name** - a clean name for display use, the well known name of the service, should be recognizable to users. - **Domains** - an array of known domains that the hoster uses. While each may serve a different purpose it is still included. - **URL** - the main url of the service. This should take you to the home page or a service page of the hoster. - **Icon** - a square image, usually a favicon or logo, that represents the service, should be recognizable as the hoster's icon. - **Status** - whether this hoster can be used on TorBox or not at the current time. It is usually a good idea to check this value before submitting a download to TorBox's servers for download. - **Type** - values are either "hoster" or "stream". Both do the same thing, but is good to differentiate services used for different things. - **Note** - a string value (or null) that may give helpful information about the current status or state of a hoster. This can and should be shown to end users. - **Daily Link Limit** - the number of downloads a user can use per day. As a user submits links, once they hit this number, the API will deny them from adding anymore of this type of link. A zero value means that it is unlimited. - **Daily Link Used** - the number of downloads a user has already used. Usually zero unless you send authentication to the endpoint. This will return accurate values. - **Daily Bandwidth Limit** - the value in bytes that a user is allowed to download from this hoster. A zero value means that it is unlimited. It is recommended to use the Daily Link Limit instead. - **Daily Bandwidth Used** - the value in bytes that a user has already used to download from this hoster. Usually zero unless you send authentication to the endpoint. This will return accurate values. ### Authorization Optional authorization. Authorization is not required in this endpoint unless you want to get the user's live data. Requires an API key using the Authorization Bearer Header to get the live and accurate data for **Daily Link Used** and **Daily Bandwidth Used**.                                                                                                                                                                                                                              |
| [EditWebDownloadItem](#editwebdownloaditem)                                     | ### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [GetUpStatus](#getupstatus)                                                     | ### Overview Simply gets if the TorBox API is available for use or not. Also might include information about downtimes. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [GetStats](#getstats)                                                           | ### Overview Simply gets general stats about the TorBox service. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [GetChangelogsRssFeed](#getchangelogsrssfeed)                                   | ### Overview Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for people who want updates on the TorBox changelog. Includes all the necessary items to make this compatible with RSS feed viewers for notifications, and proper HTML viewing. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [GetChangelogsJSON](#getchangelogsjson)                                         | ### Overview Gets most recent changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML for developers to easily render the text in a fancy yet simple way. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [GetSpeedtestFiles](#getspeedtestfiles)                                         | ### Overview Gets CDN speedtest files. This can be used for speedtesting TorBox for users or other usages, such as checking download speeds for verification. Provides all necessary data such as region, server name, and even coordinates. Uses the requesting IP to determine if the server is the closest to the user. You also have the ability to choose between long tests or short tests via the "test_length" parameter. You may also force the region selection by passing the "region" as a specific region. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [GetRssNotificationFeed](#getrssnotificationfeed)                               | ### Overview Gets your notifications in an RSS Feed which allows you to use them with RSS Feed readers or notification services that can take RSS Feeds and listen to updates. As soon as a notification goes to your account, it will be added to your feed. ### Authorization Requires an API key using as a query parameter using the `token` key.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [GetNotificationFeed](#getnotificationfeed)                                     | ### Overview Gets your notifications in a JSON object that is easily parsable compared to the RSS Feed. Gives all the same data as the RSS Feed. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [ClearAllNotifications](#clearallnotifications)                                 | ### Overview Marks all of your notifications as read and deletes them permanently. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [ClearSingleNotification](#clearsinglenotification)                             | ### Overview Marks a single notification as read and permanently deletes it from your notifications. Requires a `notification_id` which can be found by getting your notification feed. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [SendTestNotification](#sendtestnotification)                                   | ### Overview Sends a test notification to all enabled notification types. This can be useful for validating setups. No need for any body in this request. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [RefreshAPIToken](#refreshapitoken)                                             | ### Overview If you want a new API token, or your old one has been compromised, you may request a new one. If you happen to forget the token, go the [TorBox settings page ](https://torbox.app/settings) and copy the current one. If this still doesn't work, you may contact us at our support email for a new one. ### Authorization Requires an API key using the Authorization Bearer Header as well as passing the `session_token` from the website to be passed in the body. You can find the `session_token` in the localStorage of your browser on any TorBox.app page under the key `torbox_session_token`. This is a temporary token that only lasts for 1 hour, which is why it is used here to verify the identity of a user as well as their API token.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [GetUserData](#getuserdata)                                                     | ### Overview Gets a users account data and information. ### Plans `0` is Free plan `1` is Essential plan (_$3 plan_) `2` is Pro plan (_$10 plan_) `3` is Standard plan (_$5 plan_) ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [AddReferralToAccount](#addreferraltoaccount)                                   | ### Overview Automatically adds a referral code to the user's account. This can be used for developers to automatically add their referral to user's accounts who use their service. This will not override any referral a user already has. If they already have one, then it cannot be changed using this endpoint. It can only be done by the user on the [website](https://torbox.app/subscription). ### Authorization Requires an API key using the Authorization Bearer Header. Use the user's API key, not your own.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetConfirmationCode](#getconfirmationcode)                                     | ### Overview Requests a 6 digit code to be sent to the user's email for verification. Used to verify a user actually wants to perform a potentially dangerous action. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| [StartDeviceCodeAuthorizationProcess](#startdevicecodeauthorizationprocess)     | ### Overview Starts the device code authorization process. Returns all the necessary information to complete the rest of the flow. This includes the following: - `device_code` - The code your app will use for consecutive requests for getting the user token. - `code` - This is the 6 digit code that you show to the user to be entered at the following verification*url. - `verification_url` - This url which you will show as a QR code. - `friendly_verification_url` - The friendly short url which you can show as plain text to user if they cannot use the QR code method. You may also submit the `app` parameter to this request to show your app name on the `verification_url` page. This concretes user trust in authorizing your application to use their TorBox account. This is not required, but is a good idea. Default is "Third Party App". This name is also only allowed to be alpha-numeric characters (\_A-Z,0-9*). Anything else will be filtered our when being shown to the user. The `friendly_verification_url` will always show "Third Party App" to keep the URL short. The device_code is only valid for 10 minutes. Make sure the user is aware of this time limit so they submit and login before this time limit is over. ### Authorization None required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [GetTokenFromDeviceCodeProcess](#gettokenfromdevicecodeprocess)                 | ### Overview Gets the user's API token from the device code authorization process. This endpoint is expected to be polled in interval until it returns success, or until the expiry has passed. All data you need for this request is in the [original request](https://www.postman.com/wamy-dev/workspace/torbox/request/29572726-744bc274-88b7-4faf-a833-e30573c8c1e0?action=share&creator=29572726&ctx=documentation), which includes the interval to be polled by, the expiry time to stop at, and `device_code` to be sent to the API to get the correct data. ### Authorization None required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [GetReferralData](#getreferraldata)                                             | ### Overview Use this endpoint to get the user's referral data. This includes information such as how many referrals they have made, and their referral code. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [GetSubscriptions](#getsubscriptions)                                           | ### Overview Gets a users subscriptions. This includes active and inactive subscriptions for all payment processors. It should be noted that "plan_code" is not normalized to the standard plan codes such as 1, 2, or 3. These "plan_code"s are payment processor specific and used internally. If you want to display these subscriptions, use the "plan_name". For displaying information to users, use the "status", "gateway", and "plan_name". No other information is relevant to the user, and the rest is used for TorBox internally. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetTransactions](#gettransactions)                                             | ### Overview Use this endpoint to get the user's list of transactions. This data isn't exactly relevant to a third party developer, unless you want to build full parity with TorBox's dashboard, or simply want to show this information. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [GetTransactionPdf](#gettransactionpdf)                                         | ### Overview Generates a PDF of the transaction information. Requires the transaction ID from the "Get Transactions" request. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [EditUserSettings](#editusersettings)                                           | ### Overview Edits a user's TorBox settings. These values take effect immediately. Please view the body to read about allowed options and parameters. To get the user settings, simply use the "[Get User Data](https://www.postman.com/torbox/workspace/torbox-api/request/29572726-a3775d9b-f3ee-49cd-98a4-fc909ed2923b?action=share&source=copy-link&creator=29572726)" endpoint, while sending the "?settings=true" parameter ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [GetUserStats](#getuserstats)                                                   | ### Overview Gets a users stats. This data is currently split into two sections, "general" and "bandwidth". General shows general data, such as "Total Data Downloaded", "Total Data Egress", "Overall Ratio", and "Total Downloaded". Bandwidth shows the amount of bandwidth a user has downloaded within the past 30 days. We don't store this data for longer than 30 days, and it can be made more or less granular. Makes cool graphs! ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [AddRssFeed](#addrssfeed)                                                       | ### Overview Allows adding an RSS feed to your account. This API gives you a lot of customization options, but is best used with the UI on the website. This endpoint only works for Pro users (plan: 2). ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [ControlRssFeed](#controlrssfeed)                                               | ### Overview Controls an RSS Feed. By sending the RSS feed's ID and the type of operation you want to perform, it will perform said action on the RSS feed checker. Operations are either: - **Update** `forces an update on the rss feed` - **Delete** `deletes the rss feed from your account permanently. This also deletes all associated RSS feed items. This does not remove queued items spawned from the RSS feed.` - **Pause** `pauses checking the rss feed on the scan interval` - **Resume** `resumes a paused rss feed` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| [EditRssFeed](#editrssfeed)                                                     | ### Overview Allows you to edit the RSS feed based on the RSS feed ID passed. Allows you to change everything about an RSS feed, except for the URL. When updating the "rss_type" or the "torrent_seeding" options, it will update all of the RSS items associated with the RSS feed, including all previous items. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [GetUserRssFeeds](#getuserrssfeeds)                                             | ### Overview Gets all of a user's RSS feeds located on their account. Can only be accessed on a Pro (plan: 2) account. Also allows you to get a specific by passing an "id" in the parameters. ##### Type Key The "type" key in the response is used for showing what type of downloads are spawned from this RSS feed. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct. - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.` - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.` - "webdl" `uses generic links and adds as a web download when downloaded.` ##### Status Key The "status" key tells the user selected status of the RSS feed. These statuses are **not** the same as the status key for the RSS items associated with the RSS feed. Each feed also includes a "status_message" key which gives a user friendly message about anything about the RSS feed, such as errors and explainations. - "active" `the RSS feed is set to be scraped.` - "paused" `the RSS feed will not be scraped as it is paused by the user.` - "error" `the RSS feed had an error while scraping, and cannot continue.` ##### State Key The "state" key tells you exactly in real time, what the RSS feed is doing. This is useful for realtime status updates. - "idle" `the RSS feed is idle, and is currently not doing anything in the current moment.` - "scraping" `the RSS feed is currently scraping new RSS items.` - "downloading" `the RSS feed is currently downloading the newly scraped RSS items.` ##### Torrent Seeding Key The "torrent_seeding" key is returned with every response, regardless of the type of the RSS feed. This controls what seeding setting the RSS feed is using when it adds new torrents. These are the same options as when adding a torrent using the /torrent/createtorrent endpoint. \> Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. - 1 `auto, the torrent will seed based on what the user has in their settings.` - 2 `always, TorBox will be instructed to seed this torrent regardless of any other settings.` - 3 `never, TorBox will not seed this torrent.` ### Authorization Requires an API key using the Authorization Bearer Header. |
| [GetRssFeedItems](#getrssfeeditems)                                             | ### Overview Gets the first 10,000 RSS feed items associated with an RSS feed. RSS feed items are scraped items from the RSS feed. They contain a direct download link as well as the content name, and additional information. ##### Type Key The "type" key in the response is used for showing what type of downloads this RSS feed item is. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct. When the parent RSS feed's type is changed, it changes all of the items types as well. - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.` - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.` - "webdl" `uses generic links and adds as a web download when downloaded.` ##### Status Key The "status" key tells the status of the RSS feed item. This is useful for seeing what the item is doing, and what it will do next. Each item also includes a "status_message" key which gives a user friendly message about anything about the RSS feed item, such as errors and explainations. - "new" `the RSS feed item is set to be downloaded.` - "ignored" `the RSS feed item has been ignored due to the parents RSS feed settings, as it doesn't match the chosen settings. The item will not be downloaded.` - "error" `the RSS feed item had an error while attempting to download. The item will be retried on the next scrape.` - "downloaded" `the RSS feed item has already been downloaded to the user's account. It will not be downloaded again.` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [GetAllJobs](#getalljobs)                                                       | ### Overview Gets all the jobs attached to a user account. This is good for an overall view of the jobs, such as on a dashboard, or something similar. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [GetSpecificJob](#getspecificjob)                                               | ### Overview Gets a specifc job using the Job's ID. To get the ID, you will have to Get All Jobs, and get the ID you want. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [GetAllJobsByHash](#getalljobsbyhash)                                           | ### Overview Gets all jobs that match a specific hash. Good for checking on specific downloads such as a download page, that could contain a lot of jobs. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [CancelSpecificJob](#cancelspecificjob)                                         | ### Overview Cancels a job or deletes the job. Cancels while in progess (pending, uploading), or deletes the job any other time. It will delete it from the database completely. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [GetQueuedDownloads](#getqueueddownloads)                                       | ### Overview Retrieves all of a user's queued downloads by type. If you want to get all 3 types, "torrent", "usenet" and "webdl" then you will need to run this request 3 times, each with the different type. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [ControlQueuedDownloads](#controlqueueddownloads)                               | ### Overview Controls a queued torrent. By sending the queued torrent's ID and the type of operation you want to perform, it will perform that action on the queued torrent. Operations are either: - **Delete** `deletes the queued download from your account` - **Start** `starts a queued download, cannot be used with the "all" parameter` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [CreateStream](#createstream)                                                   | ### Overview Creates a new stream based off of the data provided. #### Selecting Subtitles and Audios We use indexes to determine which audio and subtitle you would like to select. To view all available indexes, you must request this endpoint intially with "null" or 0, for the chosen indexes. This will return a metadata dictionary in the response, which will contain the list of audios and subtitles available for this stream. When choosing a subtitle or audio stream, use the **relative index**, starting from 0, rather than the "index" value. This is important for selecting the correct index. Once your user gets a selection of audios and subtitles available, you can submit this to the endpoint again to get the chosen stream. - **null** - Selecting no subtitle. Not available for audio as one must always be selected. - **0** - Selecting first subtitle or audio #### Selecting Resolution Similar to the subtitles and audios, we use indexes to determine the resolution chosen. Unlike the above, these are predetermined. We do not upscale video, so if for example, you pass a 1080p video, and request it with resolution, 7, it will stay at 1080p. Please see the below for options. - **null** - Original video resolution. No downscaling will take place. - **0** - 144p - **1** - 240p - **2** - 360p - **3** - 480p - **4** - 720p - **5** - 1080p - **6** - 1440p - **7** - 2160p #### Types This endpoint allows you to handle torrents, usenet and web downloads all in one. This is why the "type" parameter is available. It defaults to "torrent". - **torrent** - Torrents - **usenet** - Usenet downloads - **webdownload** - Web Downloads #### Important Notes - Only text based subtitles are supported. If a non-text based subtitle is selected, it will be ignored. - If a subtitle or audio index does not exist on the stream, it will be ignored. - There are a lot of edge-case videos, if something doesn't work, don't hesitate to report the download hash and file_id of the failing video file. - This is a **Pro** only feature. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                       |
| [GetStreamData](#getstreamdata)                                                 | ### Overview Gets stream data from an already created stream. Requires the presigned token and user API token from the original request. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

## CreateTorrent

### Overview Creates a torrent under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. ### Rate Limits This endpoint has two limits that apply at the same time: - **300 requests per minute:** a shared ceiling that _every_ request counts against, whether the torrent is cached or uncached. - **60 uncached torrents per hour:** an additional, stricter limit that only uncached torrents count against. The important thing is that cached and uncached requests both draw from the same 300/min pool, but uncached requests _also_ draw from their own separate 60/hour pool. An uncached request therefore counts twice: once against the per-minute limit, and once against the per-hour uncached limit. **Example:** In a single minute you add 250 cached + 50 uncached torrents. That's 300 total, so you've hit the per-minute limit for that minute. But those 50 uncached torrents also came out of your 60/hour uncached budget, leaving only **10 uncached torrents** for the remaining 59 minutes of that hour. (You could still add plenty more _cached_ torrents during that hour, since cached requests aren't capped hourly: only by the 300/min ceiling.) ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/createtorrent`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| apiVersion           | string               | ✅       |                             |
| createTorrentRequest | CreateTorrentRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.CreateTorrentRequest{
  File: mainapisdk.Nullable[[]byte]([]byte{}),
  Magnet: mainapisdk.Nullable[string]("string"),
  Seed: mainapisdk.Nullable[string]("int"),
  AllowZip: mainapisdk.Nullable[string]("boolean"),
  Name: mainapisdk.Nullable[string]("string"),
  AsQueued: mainapisdk.Nullable[string]("boolean"),
  AddOnlyIfCached: mainapisdk.Nullable[string]("boolean"),
}

response, err := client.MainAPISDK1.CreateTorrent(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AsyncCreateTorrent

### Overview Creates a torrent asynchrously under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. The request will return instantly, not returning any data other than a success message. This is the same for failures, as they will return the error in the background, rather than immediately. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/asynccreatetorrent`

**Parameters**

| Name                      | Type                      | Required | Description                 |
| :------------------------ | :------------------------ | :------- | :-------------------------- |
| ctx                       | Context                   | ✅       | Default go language context |
| apiVersion                | string                    | ✅       |                             |
| asyncCreateTorrentRequest | AsyncCreateTorrentRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.AsyncCreateTorrentRequest{
  File: mainapisdk.Nullable[[]byte]([]byte{}),
  Magnet: mainapisdk.Nullable[string]("string"),
  Seed: mainapisdk.Nullable[string]("int"),
  AllowZip: mainapisdk.Nullable[string]("boolean"),
  Name: mainapisdk.Nullable[string]("string"),
  AsQueued: mainapisdk.Nullable[string]("boolean"),
  AddOnlyIfCached: mainapisdk.Nullable[string]("boolean"),
}

response, err := client.MainAPISDK1.AsyncCreateTorrent(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlTorrent

### Overview Controls a torrent. By sending the torrent's ID and the type of operation you want to perform, it will send that request to the torrent client. Operations are either: - **Reannounce** `reannounces the torrent to get new peers` - **Delete** `deletes the torrent from the client and your account permanently` - **Resume** `resumes a paused torrent` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/controltorrent`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ControlTorrent(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLink

### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/requestdl`

**Parameters**

| Name       | Type                             | Required | Description                   |
| :--------- | :------------------------------- | :------- | :---------------------------- |
| ctx        | Context                          | ✅       | Default go language context   |
| apiVersion | string                           | ✅       |                               |
| params     | RequestDownloadLinkRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinkRequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  TorrentID: mainapisdk.Nullable[string]("torrent_id"),
  FileID: mainapisdk.Nullable[string]("file_id"),
  ZipLink: mainapisdk.Nullable[string]("zip_link"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  Redirect: mainapisdk.Nullable[string]("redirect"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLink(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksGet

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/requestdl/batch`

**Parameters**

| Name       | Type                                 | Required | Description                   |
| :--------- | :----------------------------------- | :------- | :---------------------------- |
| ctx        | Context                              | ✅       | Default go language context   |
| apiVersion | string                               | ✅       |                               |
| params     | RequestDownloadLinksGetRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksGetRequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  TorrentID: mainapisdk.Nullable[string]("torrent_id"),
  FileIds: mainapisdk.Nullable[string]("file_ids"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLinksGet(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksPost

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/requestdl/batch`

**Parameters**

| Name                                       | Type                                       | Required | Description                   |
| :----------------------------------------- | :----------------------------------------- | :------- | :---------------------------- |
| ctx                                        | Context                                    | ✅       | Default go language context   |
| apiVersion                                 | string                                     | ✅       |                               |
| mainAPISDKRequestDownloadLinksPostRequest1 | MainAPISDKRequestDownloadLinksPostRequest1 | ✅       |                               |
| params                                     | RequestDownloadLinksPostRequestParams      | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksPostRequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  TorrentID: mainapisdk.Nullable[string]("torrent_id"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}


request := mainapisdk1.MainAPISDKRequestDownloadLinksPostRequest1{
  FileIds: mainapisdk.Nullable[[]float64]([]float64{}),
}

response, err := client.MainAPISDK1.RequestDownloadLinksPost(context.Background(), "api_version", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTorrentList

### Overview Gets the user's torrent list. This gives you the needed information to perform other torrent actions. This information only gets updated every 600 seconds, or when the _Request Update On Torrent_ request is sent to the _relay API_. #### Download States: - "downloading" -\> The torrent is currently downloading. - "uploading" -\> The torrent is currently seeding. - "stalled (no seeds)" -\> The torrent is trying to download, but there are no seeds connected to download from. - "paused" -\> The torrent is paused. - "completed" -\> The torrent is completely downloaded. Do not use this for download completion status. - "cached" -\> The torrent is cached from the server. - "metaDL" -\> The torrent is downloading metadata from the hoard. - "checkingResumeData" -\> The torrent is checking resumable data. All other statuses are basic qBittorrent states. [Check out here for the full list](<https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#torrent-management>). ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/mylist`

**Parameters**

| Name       | Type                        | Required | Description                   |
| :--------- | :-------------------------- | :------- | :---------------------------- |
| ctx        | Context                     | ✅       | Default go language context   |
| apiVersion | string                      | ✅       |                               |
| params     | GetTorrentListRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetTorrentListRequestParams{
  BypassCache: mainapisdk.Nullable[string]("bypass_cache"),
  ID: mainapisdk.Nullable[string]("id"),
  Offset: mainapisdk.Nullable[string]("offset"),
  Limit: mainapisdk.Nullable[string]("limit"),
  Active: mainapisdk.Nullable[string]("active"),
  Files: mainapisdk.Nullable[string]("files"),
}

response, err := client.MainAPISDK1.GetTorrentList(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTorrentCachedAvailability

### Overview Takes in a list of comma separated torrent hashes and checks if the torrent is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/checkcached`

**Parameters**

| Name       | Type                                      | Required | Description                   |
| :--------- | :---------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                   | ✅       | Default go language context   |
| apiVersion | string                                    | ✅       |                               |
| params     | GetTorrentCachedAvailabilityRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetTorrentCachedAvailabilityRequestParams{
  Hash: mainapisdk.Nullable[string]("hash"),
  Format: mainapisdk.Nullable[string]("object"),
  ListFiles: mainapisdk.Nullable[string]("list_files"),
}

response, err := client.MainAPISDK1.GetTorrentCachedAvailability(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTorrentCachedAvailabilityBatch

### Overview Takes in a list of comma separated torrent hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/checkcached`

**Parameters**

| Name                                     | Type                                     | Required | Description                 |
| :--------------------------------------- | :--------------------------------------- | :------- | :-------------------------- |
| ctx                                      | Context                                  | ✅       | Default go language context |
| apiVersion                               | string                                   | ✅       |                             |
| getTorrentCachedAvailabilityBatchRequest | GetTorrentCachedAvailabilityBatchRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.GetTorrentCachedAvailabilityBatchRequest{
  Hashes: mainapisdk.Nullable[[]string]([]string{}),
}

response, err := client.MainAPISDK1.GetTorrentCachedAvailabilityBatch(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ExportTorrentData

### Overview Exports the magnet or torrent file. Requires a type to be passed. If type is **magnet**, it will return a JSON response with the magnet as a string in the _data_ key. If type is **file**, it will return a bittorrent file as a download. Not compatible with cached downloads. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/exportdata`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| apiVersion | string                         | ✅       |                               |
| params     | ExportTorrentDataRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.ExportTorrentDataRequestParams{
  TorrentID: mainapisdk.Nullable[string]("torrent_id"),
  Type: mainapisdk.Nullable[string]("type"),
}

response, err := client.MainAPISDK1.ExportTorrentData(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTorrentInfo

### Overview A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant. **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.** ### Authorization None required.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/torrents/torrentinfo`

**Parameters**

| Name       | Type                        | Required | Description                   |
| :--------- | :-------------------------- | :------- | :---------------------------- |
| ctx        | Context                     | ✅       | Default go language context   |
| apiVersion | string                      | ✅       |                               |
| params     | GetTorrentInfoRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetTorrentInfoRequestParams{
  Hash: mainapisdk.Nullable[string]("hash"),
  Timeout: mainapisdk.Nullable[string]("30"),
  UseCacheLookup: mainapisdk.Nullable[string]("false"),
}

response, err := client.MainAPISDK1.GetTorrentInfo(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTorrentInfo1

### Overview Same as the GET route, but allows posting data such as magnet, and torrent files. Hashes will have precedence over magnets, and magnets will have precedence over torrent files. Only proper torrent files are accepted. At least one of hash, magnet, or torrent file is required. A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant. **It is not recommended to use this for checking TorBox's cache. Please use the GET /checkcache endpoint instead.** ### Authorization None required.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/torrentinfo`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| apiVersion            | string                | ✅       |                             |
| getTorrentInfoRequest | GetTorrentInfoRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.GetTorrentInfoRequest{
  Magnet: mainapisdk.Nullable[string]("string"),
  File: mainapisdk.Nullable[[]byte]([]byte{}),
  Hash: mainapisdk.Nullable[string]("string"),
  Timeout: mainapisdk.Nullable[string]("integer"),
  UseCacheLookup: mainapisdk.Nullable[string]("bool"),
  PeersOnly: mainapisdk.Nullable[string]("bool"),
}

response, err := client.MainAPISDK1.GetTorrentInfo1(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditTorrentItem

### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `PUT`
- Endpoint: `/{api_version}/api/torrents/edittorrent`

**Parameters**

| Name                   | Type                   | Required | Description                 |
| :--------------------- | :--------------------- | :------- | :-------------------------- |
| ctx                    | Context                | ✅       | Default go language context |
| apiVersion             | string                 | ✅       |                             |
| editTorrentItemRequest | EditTorrentItemRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.EditTorrentItemRequest{
  TorrentID: mainapisdk.Nullable[float64](float64(88)),
  Name: mainapisdk.Nullable[string]("string"),
  Tags: mainapisdk.Nullable[[]string]([]string{}),
  AlternativeHashes: mainapisdk.Nullable[[]string]([]string{}),
  Airlocked: mainapisdk.Nullable[bool](true),
}

response, err := client.MainAPISDK1.EditTorrentItem(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MagnetToFile

### Overview Converts any magnet to a torrent file. Does not require the magnet to be on your TorBox account or in the TorBox cache. Free endpoint for public use. ### Authorization None required.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/torrents/magnettofile`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| apiVersion          | string              | ✅       |                             |
| magnetToFileRequest | MagnetToFileRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.MagnetToFileRequest{
  Magnet: mainapisdk.Nullable[string]("string"),
}

response, err := client.MainAPISDK1.MagnetToFile(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateUsenetDownload

### Overview Creates a usenet download under your account. Simply send **either** a link, or an nzb file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. #### Post Processing Options: All post processing options that the Usenet client will perform before TorBox's own processing to make the files available. It is recommended you either don't send this parameter, or keep it at `-1` for default, which will give only the wanted files. - `-1`, Default. This runs repairs, and extractions as well as deletes the source files leaving only the wanted downloaded files. - `0`, None. No post-processing at all. Just download all the files (including all PAR2). TorBox will still do its normal processing to make the download available, but no repairs, or extraction will take place. - `1`, Repair. Download files and do a PAR2 verification. If the verification fails, download more PAR2 files and attempt to repair the files. - `2`, Repair and Unpack. Download all files, do a PAR2 verification and unpack the files. The final folder will also include the RAR and ZIP files. - `3`, Repair, Unpack and Delete. Download all files, do a PAR2 verification, unpack the files to the final folder and delete the source files. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/createusenetdownload`

**Parameters**

| Name                        | Type                        | Required | Description                 |
| :-------------------------- | :-------------------------- | :------- | :-------------------------- |
| ctx                         | Context                     | ✅       | Default go language context |
| apiVersion                  | string                      | ✅       |                             |
| createUsenetDownloadRequest | CreateUsenetDownloadRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.CreateUsenetDownloadRequest{
  File: mainapisdk.Nullable[[]byte]([]byte{}),
  Link: mainapisdk.Nullable[string]("string"),
  Name: mainapisdk.Nullable[string]("string"),
  Password: mainapisdk.Nullable[string]("string"),
  PostProcessing: mainapisdk.Nullable[string]("integer"),
  AsQueued: mainapisdk.Nullable[string]("boolean"),
  AddOnlyIfCached: mainapisdk.Nullable[string]("boolean"),
}

response, err := client.MainAPISDK1.CreateUsenetDownload(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlUsenetDownload

### Overview Controls a usenet download. By sending the usenet download's ID and the type of operation you want to perform, it will send that request to the usenet client. Operations are either: - **Delete** `deletes the nzb from the client and your account permanently` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/controlusenetdownload`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ControlUsenetDownload(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLink1

### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/requestdl`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| apiVersion | string                            | ✅       |                               |
| params     | RequestDownloadLink1RequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLink1RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  UsenetID: mainapisdk.Nullable[string]("usenet_id"),
  FileID: mainapisdk.Nullable[string]("file_id"),
  ZipLink: mainapisdk.Nullable[string]("zip_link"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  Redirect: mainapisdk.Nullable[string]("redirect"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLink1(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksGet1

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/requestdl/batch`

**Parameters**

| Name       | Type                                  | Required | Description                   |
| :--------- | :------------------------------------ | :------- | :---------------------------- |
| ctx        | Context                               | ✅       | Default go language context   |
| apiVersion | string                                | ✅       |                               |
| params     | RequestDownloadLinksGet1RequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksGet1RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  UsenetID: mainapisdk.Nullable[string]("usenet_id"),
  FileIds: mainapisdk.Nullable[string]("file_ids"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLinksGet1(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksPost1

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/requestdl/batch`

**Parameters**

| Name                                       | Type                                       | Required | Description                   |
| :----------------------------------------- | :----------------------------------------- | :------- | :---------------------------- |
| ctx                                        | Context                                    | ✅       | Default go language context   |
| apiVersion                                 | string                                     | ✅       |                               |
| mainAPISDKRequestDownloadLinksPostRequest2 | MainAPISDKRequestDownloadLinksPostRequest2 | ✅       |                               |
| params                                     | RequestDownloadLinksPost1RequestParams     | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksPost1RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  UsenetID: mainapisdk.Nullable[string]("usenet_id"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}


request := mainapisdk1.MainAPISDKRequestDownloadLinksPostRequest2{
  FileIds: mainapisdk.Nullable[[]float64]([]float64{}),
}

response, err := client.MainAPISDK1.RequestDownloadLinksPost1(context.Background(), "api_version", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUsenetList

### Overview Gets the user's usenet download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live usenet downloads. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/mylist`

**Parameters**

| Name       | Type                       | Required | Description                   |
| :--------- | :------------------------- | :------- | :---------------------------- |
| ctx        | Context                    | ✅       | Default go language context   |
| apiVersion | string                     | ✅       |                               |
| params     | GetUsenetListRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetUsenetListRequestParams{
  BypassCache: mainapisdk.Nullable[string]("bypass_cache"),
  ID: mainapisdk.Nullable[string]("id"),
  Offset: mainapisdk.Nullable[string]("offset"),
  Limit: mainapisdk.Nullable[string]("limit"),
  Active: mainapisdk.Nullable[string]("active"),
  Files: mainapisdk.Nullable[string]("files"),
}

response, err := client.MainAPISDK1.GetUsenetList(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUsenetCachedAvailability

### Overview Takes in a list of comma separated usenet hashes and checks if the usenet download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a usenet download, you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/checkcached`

**Parameters**

| Name       | Type                                     | Required | Description                   |
| :--------- | :--------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                  | ✅       | Default go language context   |
| apiVersion | string                                   | ✅       |                               |
| params     | GetUsenetCachedAvailabilityRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetUsenetCachedAvailabilityRequestParams{
  Hash: mainapisdk.Nullable[string]("hash"),
  Format: mainapisdk.Nullable[string]("format"),
  ListFiles: mainapisdk.Nullable[string]("list_files"),
}

response, err := client.MainAPISDK1.GetUsenetCachedAvailability(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUsenetCachedAvailabilityBatch

### Overview Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/checkcached`

**Parameters**

| Name                                    | Type                                    | Required | Description                 |
| :-------------------------------------- | :-------------------------------------- | :------- | :-------------------------- |
| ctx                                     | Context                                 | ✅       | Default go language context |
| apiVersion                              | string                                  | ✅       |                             |
| getUsenetCachedAvailabilityBatchRequest | GetUsenetCachedAvailabilityBatchRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.GetUsenetCachedAvailabilityBatchRequest{
  Hashes: mainapisdk.Nullable[[]string]([]string{}),
}

response, err := client.MainAPISDK1.GetUsenetCachedAvailabilityBatch(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditUsenetItem

### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `PUT`
- Endpoint: `/{api_version}/api/usenet/editusenetdownload`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| apiVersion            | string                | ✅       |                             |
| editUsenetItemRequest | EditUsenetItemRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.EditUsenetItemRequest{
  UsenetDownloadID: mainapisdk.Nullable[float64](float64(21)),
  Name: mainapisdk.Nullable[string]("string"),
  Tags: mainapisdk.Nullable[[]string]([]string{}),
  AlternativeHashes: mainapisdk.Nullable[[]string]([]string{}),
  Airlocked: mainapisdk.Nullable[bool](true),
}

response, err := client.MainAPISDK1.EditUsenetItem(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetNewsServerConnectionInformation

### Overview Gets the TorBox News Server Connection information. This is a static path, and is not dynamic. ### Authorization No authentication required.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/provider/connection`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetNewsServerConnectionInformation(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetNewsServerAccount

### Overview Gets the user's TorBox News Server Account with authentication details. Does not show password if not creating the account. Account creation happens on the first request. If you need to get the password, or reset the password, use the `/usenet/provider/account/resetpw` `POST` endpoint. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/provider/account`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetNewsServerAccount(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ResetNewsServerAccountPassword

### Overview Resets the user's TorBox News Server account password. This password is automatically generated, and is not saved, it is only shown once. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/provider/account/resetpw`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ResetNewsServerAccountPassword(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateWebDownload

### Overview Creates a web download under your account. Simply send a link to any file on the internet. Once it has been checked, it will begin downloading assuming your account has available active download slots, and they aren't too large. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/webdl/createwebdownload`

**Parameters**

| Name                     | Type                     | Required | Description                 |
| :----------------------- | :----------------------- | :------- | :-------------------------- |
| ctx                      | Context                  | ✅       | Default go language context |
| apiVersion               | string                   | ✅       |                             |
| createWebDownloadRequest | CreateWebDownloadRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.CreateWebDownloadRequest{
  Link: mainapisdk.Nullable[string]("string"),
  Password: mainapisdk.Nullable[string]("string"),
  Name: mainapisdk.Nullable[string]("string"),
  AsQueued: mainapisdk.Nullable[string]("boolean"),
  AddOnlyIfCached: mainapisdk.Nullable[string]("boolean"),
}

response, err := client.MainAPISDK1.CreateWebDownload(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlWebDownload

### Overview Controls a web download. By sending the web download's ID and the type of operation you want to perform, it will send that request to the debrid client. Operations are either: - **Delete** `deletes the download from the client and your account permanently` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/webdl/controlwebdownload`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ControlWebDownload(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLink2

### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/requestdl`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| apiVersion | string                            | ✅       |                               |
| params     | RequestDownloadLink2RequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLink2RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  WebID: mainapisdk.Nullable[string]("web_id"),
  FileID: mainapisdk.Nullable[string]("file_id"),
  ZipLink: mainapisdk.Nullable[string]("zip_link"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  Redirect: mainapisdk.Nullable[string]("redirect"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLink2(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksGet2

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/requestdl/batch`

**Parameters**

| Name       | Type                                  | Required | Description                   |
| :--------- | :------------------------------------ | :------- | :---------------------------- |
| ctx        | Context                               | ✅       | Default go language context   |
| apiVersion | string                                | ✅       |                               |
| params     | RequestDownloadLinksGet2RequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksGet2RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  WebID: mainapisdk.Nullable[string]("web_id"),
  FileIds: mainapisdk.Nullable[string]("file_ids"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}

response, err := client.MainAPISDK1.RequestDownloadLinksGet2(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLinksPost2

### Overview Requests multiple download links from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/webdl/requestdl/batch`

**Parameters**

| Name                                       | Type                                       | Required | Description                   |
| :----------------------------------------- | :----------------------------------------- | :------- | :---------------------------- |
| ctx                                        | Context                                    | ✅       | Default go language context   |
| apiVersion                                 | string                                     | ✅       |                               |
| mainAPISDKRequestDownloadLinksPostRequest3 | MainAPISDKRequestDownloadLinksPostRequest3 | ✅       |                               |
| params                                     | RequestDownloadLinksPost2RequestParams     | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.RequestDownloadLinksPost2RequestParams{
  Token: mainapisdk.Nullable[string]("token"),
  WebID: mainapisdk.Nullable[string]("web_id"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
  AppendName: mainapisdk.Nullable[string]("append_name"),
}


request := mainapisdk1.MainAPISDKRequestDownloadLinksPostRequest3{
  FileIds: mainapisdk.Nullable[[]float64]([]float64{}),
}

response, err := client.MainAPISDK1.RequestDownloadLinksPost2(context.Background(), "api_version", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebDownloadList

### Overview Gets the user's web download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live web downloads. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/mylist`

**Parameters**

| Name       | Type                            | Required | Description                   |
| :--------- | :------------------------------ | :------- | :---------------------------- |
| ctx        | Context                         | ✅       | Default go language context   |
| apiVersion | string                          | ✅       |                               |
| params     | GetWebDownloadListRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetWebDownloadListRequestParams{
  BypassCache: mainapisdk.Nullable[string]("bypass_cache"),
  ID: mainapisdk.Nullable[string]("id"),
  Offset: mainapisdk.Nullable[string]("offset"),
  Limit: mainapisdk.Nullable[string]("limit"),
  Active: mainapisdk.Nullable[string]("active"),
  Files: mainapisdk.Nullable[string]("files"),
}

response, err := client.MainAPISDK1.GetWebDownloadList(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebDownloadCachedAvailability

### Overview Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a web download you can read [this document](https://support.torbox.app/en/articles/13681109-technical-getting-hashes-for-searches) for all of the strategies. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/checkcached`

**Parameters**

| Name       | Type                                          | Required | Description                   |
| :--------- | :-------------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                       | ✅       | Default go language context   |
| apiVersion | string                                        | ✅       |                               |
| params     | GetWebDownloadCachedAvailabilityRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetWebDownloadCachedAvailabilityRequestParams{
  Hash: mainapisdk.Nullable[string]("hash"),
  Format: mainapisdk.Nullable[string]("object"),
  ListFiles: mainapisdk.Nullable[string]("false"),
}

response, err := client.MainAPISDK1.GetWebDownloadCachedAvailability(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebDownloadCachedAvailabilityBatch

### Overview Takes in a list of comma separated hashes and checks if the torrent is cached. With the POST endpoint you can submit a nearly unlimited hashes compared to the GET endpoint which is limited by url parameters. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/webdl/checkcached`

**Parameters**

| Name                                         | Type                                               | Required | Description                   |
| :------------------------------------------- | :------------------------------------------------- | :------- | :---------------------------- |
| ctx                                          | Context                                            | ✅       | Default go language context   |
| apiVersion                                   | string                                             | ✅       |                               |
| getWebDownloadCachedAvailabilityBatchRequest | GetWebDownloadCachedAvailabilityBatchRequest       | ✅       |                               |
| params                                       | GetWebDownloadCachedAvailabilityBatchRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetWebDownloadCachedAvailabilityBatchRequestParams{
  Format: mainapisdk.Nullable[string]("format"),
  ListFiles: mainapisdk.Nullable[string]("list_files"),
}


request := mainapisdk1.GetWebDownloadCachedAvailabilityBatchRequest{
  Hashes: mainapisdk.Nullable[[]string]([]string{}),
}

response, err := client.MainAPISDK1.GetWebDownloadCachedAvailabilityBatch(context.Background(), "api_version", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetHosterList

### Overview A dynamic list of hosters that TorBox is capable of downloading through its paid service. - **Name** - a clean name for display use, the well known name of the service, should be recognizable to users. - **Domains** - an array of known domains that the hoster uses. While each may serve a different purpose it is still included. - **URL** - the main url of the service. This should take you to the home page or a service page of the hoster. - **Icon** - a square image, usually a favicon or logo, that represents the service, should be recognizable as the hoster's icon. - **Status** - whether this hoster can be used on TorBox or not at the current time. It is usually a good idea to check this value before submitting a download to TorBox's servers for download. - **Type** - values are either "hoster" or "stream". Both do the same thing, but is good to differentiate services used for different things. - **Note** - a string value (or null) that may give helpful information about the current status or state of a hoster. This can and should be shown to end users. - **Daily Link Limit** - the number of downloads a user can use per day. As a user submits links, once they hit this number, the API will deny them from adding anymore of this type of link. A zero value means that it is unlimited. - **Daily Link Used** - the number of downloads a user has already used. Usually zero unless you send authentication to the endpoint. This will return accurate values. - **Daily Bandwidth Limit** - the value in bytes that a user is allowed to download from this hoster. A zero value means that it is unlimited. It is recommended to use the Daily Link Limit instead. - **Daily Bandwidth Used** - the value in bytes that a user has already used to download from this hoster. Usually zero unless you send authentication to the endpoint. This will return accurate values. ### Authorization Optional authorization. Authorization is not required in this endpoint unless you want to get the user's live data. Requires an API key using the Authorization Bearer Header to get the live and accurate data for **Daily Link Used** and **Daily Bandwidth Used**.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/hosters`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetHosterList(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditWebDownloadItem

### Overview Updates a download item based on the item's ID. Sending a PUT request with valid items overwrites the previous items entirely, so use this endpoint with caution. Please make sure to read the restrictions below as they are important. ### Restrictions - Does not affect the cached database (or any other user's item). - Name must pass validation: - Above 1 character long. - Below 200 characters long. - Cannot contain any non-url or non filesystem safe characters. (This is removed automatically). - Cannot contain any leading or trailing whitespace. You can have spaces in the middle. (This is removed automatically). - Tags follow the same rules as the name. - Alternative hashes must be of MD5, SHA1, or SHA256. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `PUT`
- Endpoint: `/{api_version}/api/webdl/editwebdownload`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| apiVersion                 | string                     | ✅       |                             |
| editWebDownloadItemRequest | EditWebDownloadItemRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.EditWebDownloadItemRequest{
  WebdlID: mainapisdk.Nullable[float64](float64(16)),
  Name: mainapisdk.Nullable[string]("string"),
  Tags: mainapisdk.Nullable[[]string]([]string{}),
  AlternativeHashes: mainapisdk.Nullable[[]string]([]string{}),
  Airlocked: mainapisdk.Nullable[bool](true),
}

response, err := client.MainAPISDK1.EditWebDownloadItem(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUpStatus

### Overview Simply gets if the TorBox API is available for use or not. Also might include information about downtimes. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetUpStatus(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetStats

### Overview Simply gets general stats about the TorBox service. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/stats`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetStats(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetChangelogsRssFeed

### Overview Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for people who want updates on the TorBox changelog. Includes all the necessary items to make this compatible with RSS feed viewers for notifications, and proper HTML viewing. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/changelogs/rss`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| apiVersion | string                            | ✅       |                               |
| params     | GetChangelogsRssFeedRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetChangelogsRssFeedRequestParams{
  Offset: mainapisdk.Nullable[string]("0"),
  Limit: mainapisdk.Nullable[string]("10"),
}

response, err := client.MainAPISDK1.GetChangelogsRssFeed(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetChangelogsJSON

### Overview Gets most recent changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML for developers to easily render the text in a fancy yet simple way. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/changelogs/json`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| apiVersion | string                         | ✅       |                               |
| params     | GetChangelogsJSONRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetChangelogsJSONRequestParams{
  Offset: mainapisdk.Nullable[string]("0"),
  Limit: mainapisdk.Nullable[string]("10"),
}

response, err := client.MainAPISDK1.GetChangelogsJSON(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSpeedtestFiles

### Overview Gets CDN speedtest files. This can be used for speedtesting TorBox for users or other usages, such as checking download speeds for verification. Provides all necessary data such as region, server name, and even coordinates. Uses the requesting IP to determine if the server is the closest to the user. You also have the ability to choose between long tests or short tests via the "test_length" parameter. You may also force the region selection by passing the "region" as a specific region. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/speedtest`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| apiVersion | string                         | ✅       |                               |
| params     | GetSpeedtestFilesRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetSpeedtestFilesRequestParams{
  TestLength: mainapisdk.Nullable[string]("short"),
  Region: mainapisdk.Nullable[string]("region"),
  UserIP: mainapisdk.Nullable[string]("user_ip"),
}

response, err := client.MainAPISDK1.GetSpeedtestFiles(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetRssNotificationFeed

### Overview Gets your notifications in an RSS Feed which allows you to use them with RSS Feed readers or notification services that can take RSS Feeds and listen to updates. As soon as a notification goes to your account, it will be added to your feed. ### Authorization Requires an API key using as a query parameter using the `token` key.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/notifications/rss`

**Parameters**

| Name       | Type                                | Required | Description                   |
| :--------- | :---------------------------------- | :------- | :---------------------------- |
| ctx        | Context                             | ✅       | Default go language context   |
| apiVersion | string                              | ✅       |                               |
| params     | GetRssNotificationFeedRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetRssNotificationFeedRequestParams{
  Token: mainapisdk.Nullable[string]("token"),
}

response, err := client.MainAPISDK1.GetRssNotificationFeed(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetNotificationFeed

### Overview Gets your notifications in a JSON object that is easily parsable compared to the RSS Feed. Gives all the same data as the RSS Feed. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/notifications/mynotifications`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetNotificationFeed(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ClearAllNotifications

### Overview Marks all of your notifications as read and deletes them permanently. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/notifications/clear`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ClearAllNotifications(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ClearSingleNotification

### Overview Marks a single notification as read and permanently deletes it from your notifications. Requires a `notification_id` which can be found by getting your notification feed. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/notifications/clear/{notification_id}`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| apiVersion     | string  | ✅       |                             |
| notificationID | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ClearSingleNotification(context.Background(), "api_version", "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SendTestNotification

### Overview Sends a test notification to all enabled notification types. This can be useful for validating setups. No need for any body in this request. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/notifications/test`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.SendTestNotification(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RefreshAPIToken

### Overview If you want a new API token, or your old one has been compromised, you may request a new one. If you happen to forget the token, go the [TorBox settings page ](https://torbox.app/settings) and copy the current one. If this still doesn't work, you may contact us at our support email for a new one. ### Authorization Requires an API key using the Authorization Bearer Header as well as passing the `session_token` from the website to be passed in the body. You can find the `session_token` in the localStorage of your browser on any TorBox.app page under the key `torbox_session_token`. This is a temporary token that only lasts for 1 hour, which is why it is used here to verify the identity of a user as well as their API token.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/user/refreshtoken`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.RefreshAPIToken(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUserData

### Overview Gets a users account data and information. ### Plans `0` is Free plan `1` is Essential plan (_$3 plan_) `2` is Pro plan (_$10 plan_) `3` is Standard plan (_$5 plan_) ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/me`

**Parameters**

| Name       | Type                     | Required | Description                   |
| :--------- | :----------------------- | :------- | :---------------------------- |
| ctx        | Context                  | ✅       | Default go language context   |
| apiVersion | string                   | ✅       |                               |
| params     | GetUserDataRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetUserDataRequestParams{
  Settings: mainapisdk.Nullable[string]("false"),
}

response, err := client.MainAPISDK1.GetUserData(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddReferralToAccount

### Overview Automatically adds a referral code to the user's account. This can be used for developers to automatically add their referral to user's accounts who use their service. This will not override any referral a user already has. If they already have one, then it cannot be changed using this endpoint. It can only be done by the user on the [website](https://torbox.app/subscription). ### Authorization Requires an API key using the Authorization Bearer Header. Use the user's API key, not your own.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/user/addreferral`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| apiVersion | string                            | ✅       |                               |
| params     | AddReferralToAccountRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.AddReferralToAccountRequestParams{
  Referral: mainapisdk.Nullable[string]("referral"),
}

response, err := client.MainAPISDK1.AddReferralToAccount(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetConfirmationCode

### Overview Requests a 6 digit code to be sent to the user's email for verification. Used to verify a user actually wants to perform a potentially dangerous action. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/getconfirmation`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetConfirmationCode(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartDeviceCodeAuthorizationProcess

### Overview Starts the device code authorization process. Returns all the necessary information to complete the rest of the flow. This includes the following: - `device_code` - The code your app will use for consecutive requests for getting the user token. - `code` - This is the 6 digit code that you show to the user to be entered at the following verification*url. - `verification_url` - This url which you will show as a QR code. - `friendly_verification_url` - The friendly short url which you can show as plain text to user if they cannot use the QR code method. You may also submit the `app` parameter to this request to show your app name on the `verification_url` page. This concretes user trust in authorizing your application to use their TorBox account. This is not required, but is a good idea. Default is "Third Party App". This name is also only allowed to be alpha-numeric characters (\_A-Z,0-9*). Anything else will be filtered our when being shown to the user. The `friendly_verification_url` will always show "Third Party App" to keep the URL short. The device_code is only valid for 10 minutes. Make sure the user is aware of this time limit so they submit and login before this time limit is over. ### Authorization None required.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/auth/device/start`

**Parameters**

| Name       | Type                                             | Required | Description                   |
| :--------- | :----------------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                          | ✅       | Default go language context   |
| apiVersion | string                                           | ✅       |                               |
| params     | StartDeviceCodeAuthorizationProcessRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.StartDeviceCodeAuthorizationProcessRequestParams{
  App: mainapisdk.Nullable[string]("Third Party App"),
}

response, err := client.MainAPISDK1.StartDeviceCodeAuthorizationProcess(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTokenFromDeviceCodeProcess

### Overview Gets the user's API token from the device code authorization process. This endpoint is expected to be polled in interval until it returns success, or until the expiry has passed. All data you need for this request is in the [original request](https://www.postman.com/wamy-dev/workspace/torbox/request/29572726-744bc274-88b7-4faf-a833-e30573c8c1e0?action=share&creator=29572726&ctx=documentation), which includes the interval to be polled by, the expiry time to stop at, and `device_code` to be sent to the API to get the correct data. ### Authorization None required.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/user/auth/device/token`

**Parameters**

| Name                                 | Type                                 | Required | Description                 |
| :----------------------------------- | :----------------------------------- | :------- | :-------------------------- |
| ctx                                  | Context                              | ✅       | Default go language context |
| apiVersion                           | string                               | ✅       |                             |
| getTokenFromDeviceCodeProcessRequest | GetTokenFromDeviceCodeProcessRequest | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


request := mainapisdk1.GetTokenFromDeviceCodeProcessRequest{
  DeviceCode: mainapisdk.Nullable[string]("XXXXXXXXXXXXXXXXXXXXXXXXX"),
}

response, err := client.MainAPISDK1.GetTokenFromDeviceCodeProcess(context.Background(), "api_version", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetReferralData

### Overview Use this endpoint to get the user's referral data. This includes information such as how many referrals they have made, and their referral code. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/referraldata`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetReferralData(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSubscriptions

### Overview Gets a users subscriptions. This includes active and inactive subscriptions for all payment processors. It should be noted that "plan_code" is not normalized to the standard plan codes such as 1, 2, or 3. These "plan_code"s are payment processor specific and used internally. If you want to display these subscriptions, use the "plan_name". For displaying information to users, use the "status", "gateway", and "plan_name". No other information is relevant to the user, and the rest is used for TorBox internally. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/subscriptions`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetSubscriptions(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTransactions

### Overview Use this endpoint to get the user's list of transactions. This data isn't exactly relevant to a third party developer, unless you want to build full parity with TorBox's dashboard, or simply want to show this information. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/transactions`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetTransactions(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTransactionPdf

### Overview Generates a PDF of the transaction information. Requires the transaction ID from the "Get Transactions" request. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/transaction/pdf`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| apiVersion | string                         | ✅       |                               |
| params     | GetTransactionPdfRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetTransactionPdfRequestParams{
  TransactionID: mainapisdk.Nullable[string]("transaction_id"),
}

response, err := client.MainAPISDK1.GetTransactionPdf(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditUserSettings

### Overview Edits a user's TorBox settings. These values take effect immediately. Please view the body to read about allowed options and parameters. To get the user settings, simply use the "[Get User Data](https://www.postman.com/torbox/workspace/torbox-api/request/29572726-a3775d9b-f3ee-49cd-98a4-fc909ed2923b?action=share&source=copy-link&creator=29572726)" endpoint, while sending the "?settings=true" parameter ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `PUT`
- Endpoint: `/{api_version}/api/user/settings/editsettings`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.EditUserSettings(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUserStats

### Overview Gets a users stats. This data is currently split into two sections, "general" and "bandwidth". General shows general data, such as "Total Data Downloaded", "Total Data Egress", "Overall Ratio", and "Total Downloaded". Bandwidth shows the amount of bandwidth a user has downloaded within the past 30 days. We don't store this data for longer than 30 days, and it can be made more or less granular. Makes cool graphs! ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/user/stats`

**Parameters**

| Name       | Type                      | Required | Description                   |
| :--------- | :------------------------ | :------- | :---------------------------- |
| ctx        | Context                   | ✅       | Default go language context   |
| apiVersion | string                    | ✅       |                               |
| params     | GetUserStatsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetUserStatsRequestParams{
  General: mainapisdk.Nullable[string]("true"),
  Bandwidth: mainapisdk.Nullable[string]("false"),
  BandwidthGrouping: mainapisdk.Nullable[string]("day"),
}

response, err := client.MainAPISDK1.GetUserStats(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddRssFeed

### Overview Allows adding an RSS feed to your account. This API gives you a lot of customization options, but is best used with the UI on the website. This endpoint only works for Pro users (plan: 2). ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/rss/addrss`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.AddRssFeed(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlRssFeed

### Overview Controls an RSS Feed. By sending the RSS feed's ID and the type of operation you want to perform, it will perform said action on the RSS feed checker. Operations are either: - **Update** `forces an update on the rss feed` - **Delete** `deletes the rss feed from your account permanently. This also deletes all associated RSS feed items. This does not remove queued items spawned from the RSS feed.` - **Pause** `pauses checking the rss feed on the scan interval` - **Resume** `resumes a paused rss feed` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/rss/controlrss`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ControlRssFeed(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditRssFeed

### Overview Allows you to edit the RSS feed based on the RSS feed ID passed. Allows you to change everything about an RSS feed, except for the URL. When updating the "rss_type" or the "torrent_seeding" options, it will update all of the RSS items associated with the RSS feed, including all previous items. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/rss/modifyrss`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.EditRssFeed(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUserRssFeeds

### Overview Gets all of a user's RSS feeds located on their account. Can only be accessed on a Pro (plan: 2) account. Also allows you to get a specific by passing an "id" in the parameters. ##### Type Key The "type" key in the response is used for showing what type of downloads are spawned from this RSS feed. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct. - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.` - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.` - "webdl" `uses generic links and adds as a web download when downloaded.` ##### Status Key The "status" key tells the user selected status of the RSS feed. These statuses are **not** the same as the status key for the RSS items associated with the RSS feed. Each feed also includes a "status_message" key which gives a user friendly message about anything about the RSS feed, such as errors and explainations. - "active" `the RSS feed is set to be scraped.` - "paused" `the RSS feed will not be scraped as it is paused by the user.` - "error" `the RSS feed had an error while scraping, and cannot continue.` ##### State Key The "state" key tells you exactly in real time, what the RSS feed is doing. This is useful for realtime status updates. - "idle" `the RSS feed is idle, and is currently not doing anything in the current moment.` - "scraping" `the RSS feed is currently scraping new RSS items.` - "downloading" `the RSS feed is currently downloading the newly scraped RSS items.` ##### Torrent Seeding Key The "torrent_seeding" key is returned with every response, regardless of the type of the RSS feed. This controls what seeding setting the RSS feed is using when it adds new torrents. These are the same options as when adding a torrent using the /torrent/createtorrent endpoint. \> Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. - 1 `auto, the torrent will seed based on what the user has in their settings.` - 2 `always, TorBox will be instructed to seed this torrent regardless of any other settings.` - 3 `never, TorBox will not seed this torrent.` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/rss/getfeeds`

**Parameters**

| Name       | Type                         | Required | Description                   |
| :--------- | :--------------------------- | :------- | :---------------------------- |
| ctx        | Context                      | ✅       | Default go language context   |
| apiVersion | string                       | ✅       |                               |
| params     | GetUserRssFeedsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetUserRssFeedsRequestParams{
  ID: mainapisdk.Nullable[string]("id"),
}

response, err := client.MainAPISDK1.GetUserRssFeeds(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetRssFeedItems

### Overview Gets the first 10,000 RSS feed items associated with an RSS feed. RSS feed items are scraped items from the RSS feed. They contain a direct download link as well as the content name, and additional information. ##### Type Key The "type" key in the response is used for showing what type of downloads this RSS feed item is. This affects how the item is downloaded, as well as what is scraped, so make sure it is correct. When the parent RSS feed's type is changed, it changes all of the items types as well. - **"**torrent**"** `uses .torrent files and magnet links and adds as torrent when downloaded.` - "usenet" `uses .nzb files and nzb links and adds as a usenet download when downloaded.` - "webdl" `uses generic links and adds as a web download when downloaded.` ##### Status Key The "status" key tells the status of the RSS feed item. This is useful for seeing what the item is doing, and what it will do next. Each item also includes a "status_message" key which gives a user friendly message about anything about the RSS feed item, such as errors and explainations. - "new" `the RSS feed item is set to be downloaded.` - "ignored" `the RSS feed item has been ignored due to the parents RSS feed settings, as it doesn't match the chosen settings. The item will not be downloaded.` - "error" `the RSS feed item had an error while attempting to download. The item will be retried on the next scrape.` - "downloaded" `the RSS feed item has already been downloaded to the user's account. It will not be downloaded again.` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/rss/getfeeditems`

**Parameters**

| Name       | Type                         | Required | Description                   |
| :--------- | :--------------------------- | :------- | :---------------------------- |
| ctx        | Context                      | ✅       | Default go language context   |
| apiVersion | string                       | ✅       |                               |
| params     | GetRssFeedItemsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetRssFeedItemsRequestParams{
  RssFeedID: mainapisdk.Nullable[string]("rss_feed_id"),
}

response, err := client.MainAPISDK1.GetRssFeedItems(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetAllJobs

### Overview Gets all the jobs attached to a user account. This is good for an overall view of the jobs, such as on a dashboard, or something similar. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/integration/jobs`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetAllJobs(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSpecificJob

### Overview Gets a specifc job using the Job's ID. To get the ID, you will have to Get All Jobs, and get the ID you want. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/integration/job/{job_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| jobID      | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetSpecificJob(context.Background(), "api_version", "job_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetAllJobsByHash

### Overview Gets all jobs that match a specific hash. Good for checking on specific downloads such as a download page, that could contain a lot of jobs. ### Statuses - "pending" -\> Upload is still waiting in the queue. Waiting for spot to upload. - "uploading" -\> Upload is uploading to the proper remote. Progress will be updated as upload continues. - "completed" -\> Upload has successfully been uploaded. Progress will be at 1, and the download URL will be populated. - "failed" -\> The upload has failed. Check the Detail key for information. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/integration/jobs/{hash}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| hash       | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.GetAllJobsByHash(context.Background(), "api_version", "hash")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CancelSpecificJob

### Overview Cancels a job or deletes the job. Cancels while in progess (pending, uploading), or deletes the job any other time. It will delete it from the database completely. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `DELETE`
- Endpoint: `/{api_version}/api/integration/job/{job_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| jobID      | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.CancelSpecificJob(context.Background(), "api_version", "job_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetQueuedDownloads

### Overview Retrieves all of a user's queued downloads by type. If you want to get all 3 types, "torrent", "usenet" and "webdl" then you will need to run this request 3 times, each with the different type. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/queued/getqueued`

**Parameters**

| Name       | Type                            | Required | Description                   |
| :--------- | :------------------------------ | :------- | :---------------------------- |
| ctx        | Context                         | ✅       | Default go language context   |
| apiVersion | string                          | ✅       |                               |
| params     | GetQueuedDownloadsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetQueuedDownloadsRequestParams{
  BypassCache: mainapisdk.Nullable[string]("bypass_cache"),
  ID: mainapisdk.Nullable[string]("id"),
  Offset: mainapisdk.Nullable[string]("offset"),
  Limit: mainapisdk.Nullable[string]("limit"),
  Type: mainapisdk.Nullable[string]("torrent"),
}

response, err := client.MainAPISDK1.GetQueuedDownloads(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlQueuedDownloads

### Overview Controls a queued torrent. By sending the queued torrent's ID and the type of operation you want to perform, it will perform that action on the queued torrent. Operations are either: - **Delete** `deletes the queued download from your account` - **Start** `starts a queued download, cannot be used with the "all" parameter` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/queued/controlqueued`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)

response, err := client.MainAPISDK1.ControlQueuedDownloads(context.Background(), "api_version")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateStream

### Overview Creates a new stream based off of the data provided. #### Selecting Subtitles and Audios We use indexes to determine which audio and subtitle you would like to select. To view all available indexes, you must request this endpoint intially with "null" or 0, for the chosen indexes. This will return a metadata dictionary in the response, which will contain the list of audios and subtitles available for this stream. When choosing a subtitle or audio stream, use the **relative index**, starting from 0, rather than the "index" value. This is important for selecting the correct index. Once your user gets a selection of audios and subtitles available, you can submit this to the endpoint again to get the chosen stream. - **null** - Selecting no subtitle. Not available for audio as one must always be selected. - **0** - Selecting first subtitle or audio #### Selecting Resolution Similar to the subtitles and audios, we use indexes to determine the resolution chosen. Unlike the above, these are predetermined. We do not upscale video, so if for example, you pass a 1080p video, and request it with resolution, 7, it will stay at 1080p. Please see the below for options. - **null** - Original video resolution. No downscaling will take place. - **0** - 144p - **1** - 240p - **2** - 360p - **3** - 480p - **4** - 720p - **5** - 1080p - **6** - 1440p - **7** - 2160p #### Types This endpoint allows you to handle torrents, usenet and web downloads all in one. This is why the "type" parameter is available. It defaults to "torrent". - **torrent** - Torrents - **usenet** - Usenet downloads - **webdownload** - Web Downloads #### Important Notes - Only text based subtitles are supported. If a non-text based subtitle is selected, it will be ignored. - If a subtitle or audio index does not exist on the stream, it will be ignored. - There are a lot of edge-case videos, if something doesn't work, don't hesitate to report the download hash and file_id of the failing video file. - This is a **Pro** only feature. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/stream/createstream`

**Parameters**

| Name       | Type                      | Required | Description                   |
| :--------- | :------------------------ | :------- | :---------------------------- |
| ctx        | Context                   | ✅       | Default go language context   |
| apiVersion | string                    | ✅       |                               |
| params     | CreateStreamRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.CreateStreamRequestParams{
  ID: mainapisdk.Nullable[string]("id"),
  FileID: mainapisdk.Nullable[string]("0"),
  Type: mainapisdk.Nullable[string]("torrent"),
  ChosenSubtitleIndex: mainapisdk.Nullable[string]("chosen_subtitle_index"),
  ChosenAudioIndex: mainapisdk.Nullable[string]("0"),
  ChosenResolutionIndex: mainapisdk.Nullable[string]("chosen_resolution_index"),
}

response, err := client.MainAPISDK1.CreateStream(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetStreamData

### Overview Gets stream data from an already created stream. Requires the presigned token and user API token from the original request. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/stream/getstreamdata`

**Parameters**

| Name       | Type                       | Required | Description                   |
| :--------- | :------------------------- | :------- | :---------------------------- |
| ctx        | Context                    | ✅       | Default go language context   |
| apiVersion | string                     | ✅       |                               |
| params     | GetStreamDataRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/main-api-sdk"
  "example.com/main-api-sdk/mainapisdk1"
)

config := mainapisdk.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := mainapisdk.NewMainAPISDK(config)


params := mainapisdk1.GetStreamDataRequestParams{
  PresignedToken: mainapisdk.Nullable[string]("presigned_token"),
  Token: mainapisdk.Nullable[string]("token"),
  ChosenSubtitleIndex: mainapisdk.Nullable[string]("chosen_subtitle_index"),
  ChosenAudioIndex: mainapisdk.Nullable[string]("chosen_audio_index"),
  ChosenResolutionIndex: mainapisdk.Nullable[string]("chosen_resolution_index"),
}

response, err := client.MainAPISDK1.GetStreamData(context.Background(), "api_version", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
