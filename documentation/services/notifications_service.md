# NotificationsService

A list of all methods in the `NotificationsService` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description                                                                                                                                                                                                                                                                                                                                           |
| :-------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetRssNotificationFeed](#getrssnotificationfeed)   | ### Overview Gets your notifications in an RSS Feed which allows you to use them with RSS Feed readers or notification services that can take RSS Feeds and listen to updates. As soon as a notification goes to your account, it will be added to your feed. ### Authorization Requires an API key using as a query parameter using the `token` key. |
| [GetNotificationFeed](#getnotificationfeed)         | ### Overview Gets your notifications in a JSON object that is easily parsable compared to the RSS Feed. Gives all the same data as the RSS Feed. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                         |
| [ClearAllNotifications](#clearallnotifications)     | ### Overview Marks all of your notifications as read and deletes them permanently. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                                                                                       |
| [ClearSingleNotification](#clearsinglenotification) | ### Overview Marks a single notification as read and permanently deletes it from your notifications. Requires a `notification_id` which can be found by getting your notification feed. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                  |
| [SendTestNotification](#sendtestnotification)       | ### Overview Sends a test notification to all enabled notification types. This can be useful for validating setups. No need for any body in this request. ### Authorization Requires an API key using the Authorization Bearer Header.                                                                                                                |

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
  "torbox-sdk-go/pkg/torboxapiconfig"
  "torbox-sdk-go/pkg/torboxapi"
  "torbox-sdk-go/pkg/notifications"
)

config := torboxapiconfig.NewConfig()
client := torboxapi.NewTorboxApi(config)


params := notifications.GetRssNotificationFeedRequestParams{}


response, err := client.Notifications.GetRssNotificationFeed(context.Background(), "apiVersion", params)
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

`GetNotificationFeedOkResponse`

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

response, err := client.Notifications.GetNotificationFeed(context.Background(), "apiVersion")
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

response, err := client.Notifications.ClearAllNotifications(context.Background(), "apiVersion")
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
| notificationId | string  | ✅       |                             |

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

response, err := client.Notifications.ClearSingleNotification(context.Background(), "apiVersion", "notificationId")
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

response, err := client.Notifications.SendTestNotification(context.Background(), "apiVersion")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
