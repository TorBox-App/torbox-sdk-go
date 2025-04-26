# GeneralService

A list of all methods in the `GeneralService` service. Click on the method name to view detailed information about that method.

| Methods                                       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| :-------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetUpStatus](#getupstatus)                   | ### Overview Simply gets if the TorBox API is available for use or not. Also might include information about downtimes. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                 |
| [GetStats](#getstats)                         | ### Overview Simply gets general stats about the TorBox service. ### Authorization None needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [GetChangelogsRssFeed](#getchangelogsrssfeed) | ### Overview Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for people who want updates on the TorBox changelog. Includes all the necessary items to make this compatible with RSS feed viewers for notifications, and proper HTML viewing. ### Authorization None needed.                                                                                                                                                                                       |
| [GetChangelogsJson](#getchangelogsjson)       | ### Overview Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML and markdown for developers to easily render the text in a fancy yet simple way. ### Authorization None needed.                                                                                                                                                                  |
| [GetSpeedtestFiles](#getspeedtestfiles)       | ### Overview Gets CDN speedtest files. This can be used for speedtesting TorBox for users or other usages, such as checking download speeds for verification. Provides all necessary data such as region, server name, and even coordinates. Uses the requesting IP to determine if the server is the closest to the user. You also have the ability to choose between long tests or short tests via the "test_length" parameter. You may also force the region selection by passing the "region" as a specific region. ### Authorization None needed. |

## GetUpStatus

### Overview Simply gets if the TorBox API is available for use or not. Also might include information about downtimes. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`GetUpStatusOkResponse`

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

response, err := client.General.GetUpStatus(context.Background())
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

`GetStatsOkResponse`

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

response, err := client.General.GetStats(context.Background(), "apiVersion")
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
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)

response, err := client.General.GetChangelogsRssFeed(context.Background(), "apiVersion")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetChangelogsJson

### Overview Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML and markdown for developers to easily render the text in a fancy yet simple way. ### Authorization None needed.

- HTTP Method: `GET`
- Endpoint: `/{api_version}/api/changelogs/json`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| apiVersion | string  | ✅       |                             |

**Return Type**

`GetChangelogsJsonOkResponse`

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

response, err := client.General.GetChangelogsJson(context.Background(), "apiVersion")
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

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"

  "torbox-sdk-go/pkg/general"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := general.GetSpeedtestFilesRequestParams{

}

response, err := client.General.GetSpeedtestFiles(context.Background(), "apiVersion", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
