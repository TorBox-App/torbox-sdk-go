# CreateTorrentOkResponse

**Properties**

| Name    | Type                                 | Required | Description |
| :------ | :----------------------------------- | :------- | :---------- |
| Data    | torrents.CreateTorrentOkResponseData | ❌       |             |
| Detail  | string                               | ❌       |             |
| Error   | any                                  | ❌       |             |
| Success | bool                                 | ❌       |             |

# CreateTorrentOkResponseData

**Properties**

| Name                   | Type    | Required | Description |
| :--------------------- | :------ | :------- | :---------- |
| ActiveLimit            | float64 | ❌       |             |
| AuthId                 | string  | ❌       |             |
| CurrentActiveDownloads | float64 | ❌       |             |
| Hash                   | string  | ❌       |             |
| QueuedId               | float64 | ❌       |             |
| TorrentId              | float64 | ❌       |             |
