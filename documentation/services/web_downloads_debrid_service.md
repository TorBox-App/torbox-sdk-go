# WebDownloadsDebridService

A list of all methods in the `WebDownloadsDebridService` service. Click on the method name to view detailed information about that method.

| Methods                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| :-------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreateWebDownload](#createwebdownload)                               | ### Overview Creates a web download under your account. Simply send a link to any file on the internet. Once it has been checked, it will begin downloading assuming your account has available active download slots, and they aren't too large. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [ControlWebDownload](#controlwebdownload)                             | ### Overview Controls a web download. By sending the web download's ID and the type of operation you want to perform, it will send that request to the debrid client. Operations are either: - **Delete** `deletes the download from the client and your account permanently` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [RequestDownloadLink2](#requestdownloadlink2)                         | ### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 3 hours for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Permalinks Instead of generating many CDN urls by requesting this endpoint, you can instead create a permalink such as: `https://api.torbox.app/v1/api/torrents/requestdl?token=APIKEY&torrent_id=NUMBER&file_id=NUMBER&redirect=true` and when a user clicks on it, it will automatically redirect them to the CDN link. This saves requests and doesn't abuse the API. Use this method rather than saving CDN links as they are not permanent. To invalidate these permalinks, simply reset your API token or delete the item from your dashboard. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [GetWebDownloadList](#getwebdownloadlist)                             | ### Overview Gets the user's web download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live web downloads. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [GetWebDownloadCachedAvailability](#getwebdownloadcachedavailability) | ### Overview Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a web download, pass the link through an md5 hash algo and it will return the proper hash for it. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| [GetHosterList](#gethosterlist)                                       | ### Overview A dynamic list of hosters that TorBox is capable of downloading through its paid service. - **Name** - a clean name for display use, the well known name of the service, should be recognizable to users. - **Domains** - an array of known domains that the hoster uses. While each may serve a different purpose it is still included. - **URL** - the main url of the service. This should take you to the home page or a service page of the hoster. - **Icon** - a square image, usually a favicon or logo, that represents the service, should be recognizable as the hoster's icon. - **Status** - whether this hoster can be used on TorBox or not at the current time. It is usually a good idea to check this value before submitting a download to TorBox's servers for download. - **Type** - values are either "hoster" or "stream". Both do the same thing, but is good to differentiate services used for different things. - **Note** - a string value (or null) that may give helpful information about the current status or state of a hoster. This can and should be shown to end users. - **Daily Link Limit** - the number of downloads a user can use per day. As a user submits links, once they hit this number, the API will deny them from adding anymore of this type of link. A zero value means that it is unlimited. - **Daily Link Used** - the number of downloads a user has already used. Usually zero unless you send authentication to the endpoint. This will return accurate values. - **Daily Bandwidth Limit** - the value in bytes that a user is allowed to download from this hoster. A zero value means that it is unlimited. It is recommended to use the Daily Link Limit instead. - **Daily Bandwidth Used** - the value in bytes that a user has already used to download from this hoster. Usually zero unless you send authentication to the endpoint. This will return accurate values. ### Authorization Optional authorization. Authorization is not required in this endpoint unless you want to get the user's live data. Requires an API key using the Authorization Bearer Header to get the live and accurate data for **Daily Link Used** and **Daily Bandwidth Used**. |

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

`CreateWebDownloadOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/util"
  "torbox-sdk-go/pkg/webdownloadsdebrid"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


request := webdownloadsdebrid.CreateWebDownloadRequest{
  Link: util.ToPointer("Link"),
}

response, err := client.WebDownloadsDebrid.CreateWebDownload(context.Background(), "apiVersion", request)
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

| Name       | Type                            | Required | Description                   |
| :--------- | :------------------------------ | :------- | :---------------------------- |
| ctx        | Context                         | ✅       | Default go language context   |
| apiVersion | string                          | ✅       |                               |
| body       | []byte                          | ✅       |                               |
| params     | ControlWebDownloadRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

  "torbox-sdk-go/pkg/webdownloadsdebrid"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := webdownloadsdebrid.ControlWebDownloadRequestParams{

}

response, err := client.WebDownloadsDebrid.ControlWebDownload(context.Background(), "apiVersion", request, params)
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

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

  "torbox-sdk-go/pkg/webdownloadsdebrid"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := webdownloadsdebrid.RequestDownloadLink2RequestParams{

}

response, err := client.WebDownloadsDebrid.RequestDownloadLink2(context.Background(), "apiVersion", params)
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

`GetWebDownloadListOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

  "torbox-sdk-go/pkg/webdownloadsdebrid"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := webdownloadsdebrid.GetWebDownloadListRequestParams{

}

response, err := client.WebDownloadsDebrid.GetWebDownloadList(context.Background(), "apiVersion", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebDownloadCachedAvailability

### Overview Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a web download, pass the link through an md5 hash algo and it will return the proper hash for it. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/webdl/checkcached`

**Parameters**

| Name       | Type                                          | Required | Description                   |
| :--------- | :-------------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                       | ✅       | Default go language context   |
| apiVersion | string                                        | ✅       |                               |
| params     | GetWebDownloadCachedAvailabilityRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

  "torbox-sdk-go/pkg/webdownloadsdebrid"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := webdownloadsdebrid.GetWebDownloadCachedAvailabilityRequestParams{

}

response, err := client.WebDownloadsDebrid.GetWebDownloadCachedAvailability(context.Background(), "apiVersion", params)
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

`GetHosterListOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)

response, err := client.WebDownloadsDebrid.GetHosterList(context.Background(), "apiVersion")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
