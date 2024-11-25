# GetUserDataOkResponse

**Properties**

| Name    | Type                           | Required | Description |
| :------ | :----------------------------- | :------- | :---------- |
| Data    | user.GetUserDataOkResponseData | ❌       |             |
| Detail  | string                         | ❌       |             |
| Error   | any                            | ❌       |             |
| Success | bool                           | ❌       |             |

# GetUserDataOkResponseData

**Properties**

| Name             | Type          | Required | Description |
| :--------------- | :------------ | :------- | :---------- |
| AuthId           | string        | ❌       |             |
| BaseEmail        | string        | ❌       |             |
| CooldownUntil    | string        | ❌       |             |
| CreatedAt        | string        | ❌       |             |
| Customer         | string        | ❌       |             |
| Email            | string        | ❌       |             |
| Id               | float64       | ❌       |             |
| IsSubscribed     | bool          | ❌       |             |
| Plan             | float64       | ❌       |             |
| PremiumExpiresAt | string        | ❌       |             |
| Server           | float64       | ❌       |             |
| Settings         | user.Settings | ❌       |             |
| TotalDownloaded  | float64       | ❌       |             |
| UpdatedAt        | string        | ❌       |             |
| UserReferral     | string        | ❌       |             |

# Settings

**Properties**

| Name           | Type   | Required | Description |
| :------------- | :----- | :------- | :---------- |
| Anothersetting | string | ❌       |             |
| Setting        | string | ❌       |             |
