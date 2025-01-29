# UsenetService

A list of all methods in the `UsenetService` service. Click on the method name to view detailed information about that method.

| Methods                                                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| :---------------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreateUsenetDownload](#createusenetdownload)               | ### Overview Creates a usenet download under your account. Simply send **either** a link, or an nzb file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large. #### Post Processing Options: All post processing options that the Usenet client will perform before TorBox's own processing to make the files available. It is recommended you either don't send this parameter, or keep it at `-1` for default, which will give only the wanted files. - `-1`, Default. This runs repairs, and extractions as well as deletes the source files leaving only the wanted downloaded files. - `0`, None. No post-processing at all. Just download all the files (including all PAR2). TorBox will still do its normal processing to make the download available, but no repairs, or extraction will take place. - `1`, Repair. Download files and do a PAR2 verification. If the verification fails, download more PAR2 files and attempt to repair the files. - `2`, Repair and Unpack. Download all files, do a PAR2 verification and unpack the files. The final folder will also include the RAR and ZIP files. - `3`, Repair, Unpack and Delete. Download all files, do a PAR2 verification, unpack the files to the final folder and delete the source files. ### Authorization Requires an API key using the Authorization Bearer Header. |
| [ControlUsenetDownload](#controlusenetdownload)             | ### Overview Controls a usenet download. By sending the usenet download's ID and the type of operation you want to perform, it will send that request to the usenet client. Operations are either: - **Delete** `deletes the nzb from the client and your account permanently` - **Pause** `pauses a nzb's download` - **Resume** `resumes a paused usenet download` ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [RequestDownloadLink1](#requestdownloadlink1)               | ### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 1 hour for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [GetUsenetList](#getusenetlist)                             | ### Overview Gets the user's usenet download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live usenet downloads. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [GetUsenetCachedAvailability](#getusenetcachedavailability) | ### Overview Takes in a list of comma separated usenet hashes and checks if the usenet download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a usenet download, pass the link or file through an md5 hash algo and it will return the proper hash for it. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |

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

`CreateUsenetDownloadOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/usenet"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


request := usenet.CreateUsenetDownloadRequest{}
request.SetFile("")
request.SetLink("Link")

response, err := client.Usenet.CreateUsenetDownload(context.Background(), "apiVersion", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ControlUsenetDownload

### Overview Controls a usenet download. By sending the usenet download's ID and the type of operation you want to perform, it will send that request to the usenet client. Operations are either: - **Delete** `deletes the nzb from the client and your account permanently` - **Pause** `pauses a nzb's download` - **Resume** `resumes a paused usenet download` ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `POST`
- Endpoint: `/{api_version}/api/usenet/controlusenetdownload`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |
| body       | []byte  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Usenet.ControlUsenetDownload(context.Background(), "apiVersion", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RequestDownloadLink1

### Overview Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 1 hour for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing. ### Authorization Requires an API key as a parameter for the `token` parameter.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/requestdl`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| apiVersion | string                            | ✅       |                               |
| params     | RequestDownloadLink1RequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/usenet"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := usenet.RequestDownloadLink1RequestParams{}


response, err := client.Usenet.RequestDownloadLink1(context.Background(), "apiVersion", params)
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

`GetUsenetListOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/usenet"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := usenet.GetUsenetListRequestParams{}


response, err := client.Usenet.GetUsenetList(context.Background(), "apiVersion", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetUsenetCachedAvailability

### Overview Takes in a list of comma separated usenet hashes and checks if the usenet download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as: `?hash=XXXX&hash=XXXX&hash=XXXX` or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX` and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour. You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below. To get the hash of a usenet download, pass the link or file through an md5 hash algo and it will return the proper hash for it. ### Authorization Requires an API key using the Authorization Bearer Header.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/usenet/checkcached`

**Parameters**

| Name       | Type                                     | Required | Description                   |
| :--------- | :--------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                  | ✅       | Default go language context   |
| apiVersion | string                                   | ✅       |                               |
| params     | GetUsenetCachedAvailabilityRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/usenet"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := usenet.GetUsenetCachedAvailabilityRequestParams{}


response, err := client.Usenet.GetUsenetCachedAvailability(context.Background(), "apiVersion", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
