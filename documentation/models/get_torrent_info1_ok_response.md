# GetTorrentInfo1OkResponse

**Properties**

| Name    | Type                                   | Required | Description |
| :------ | :------------------------------------- | :------- | :---------- |
| Data    | torrents.GetTorrentInfo1OkResponseData | ❌       |             |
| Detail  | string                                 | ❌       |             |
| Error   | any                                    | ❌       |             |
| Success | bool                                   | ❌       |             |

# GetTorrentInfo1OkResponseData

**Properties**

| Name     | Type                  | Required | Description |
| :------- | :-------------------- | :------- | :---------- |
| Files    | []torrents.DataFiles3 | ❌       |             |
| Hash     | string                | ❌       |             |
| Name     | string                | ❌       |             |
| Peers    | float64               | ❌       |             |
| Seeds    | float64               | ❌       |             |
| Size     | float64               | ❌       |             |
| Trackers | []any                 | ❌       |             |

# DataFiles3

**Properties**

| Name | Type    | Required | Description |
| :--- | :------ | :------- | :---------- |
| Name | string  | ❌       |             |
| Size | float64 | ❌       |             |
