# GetTorrentInfoOkResponse

**Properties**

| Name    | Type                                  | Required | Description |
| :------ | :------------------------------------ | :------- | :---------- |
| Data    | torrents.GetTorrentInfoOkResponseData | ❌       |             |
| Detail  | string                                | ❌       |             |
| Error   | any                                   | ❌       |             |
| Success | bool                                  | ❌       |             |

# GetTorrentInfoOkResponseData

**Properties**

| Name     | Type                  | Required | Description |
| :------- | :-------------------- | :------- | :---------- |
| Files    | []torrents.DataFiles2 | ❌       |             |
| Hash     | string                | ❌       |             |
| Name     | string                | ❌       |             |
| Peers    | float64               | ❌       |             |
| Seeds    | float64               | ❌       |             |
| Size     | float64               | ❌       |             |
| Trackers | []any                 | ❌       |             |

# DataFiles2

**Properties**

| Name | Type    | Required | Description |
| :--- | :------ | :------- | :---------- |
| Name | string  | ❌       |             |
| Size | float64 | ❌       |             |
