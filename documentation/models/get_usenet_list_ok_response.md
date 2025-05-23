# GetUsenetListOkResponse

**Properties**

| Name    | Type                                 | Required | Description |
| :------ | :----------------------------------- | :------- | :---------- |
| Data    | []usenet.GetUsenetListOkResponseData | ❌       |             |
| Detail  | string                               | ❌       |             |
| Error   | any                                  | ❌       |             |
| Success | bool                                 | ❌       |             |

# GetUsenetListOkResponseData

**Properties**

| Name             | Type                | Required | Description |
| :--------------- | :------------------ | :------- | :---------- |
| Active           | bool                | ❌       |             |
| AuthId           | string              | ❌       |             |
| Availability     | float64             | ❌       |             |
| CreatedAt        | string              | ❌       |             |
| DownloadFinished | bool                | ❌       |             |
| DownloadPresent  | bool                | ❌       |             |
| DownloadSpeed    | float64             | ❌       |             |
| DownloadState    | string              | ❌       |             |
| Eta              | float64             | ❌       |             |
| ExpiresAt        | string              | ❌       |             |
| Files            | []usenet.DataFiles4 | ❌       |             |
| Hash             | string              | ❌       |             |
| Id               | float64             | ❌       |             |
| InactiveCheck    | float64             | ❌       |             |
| Name             | string              | ❌       |             |
| Progress         | float64             | ❌       |             |
| Server           | float64             | ❌       |             |
| Size             | float64             | ❌       |             |
| TorrentFile      | bool                | ❌       |             |
| UpdatedAt        | string              | ❌       |             |
| UploadSpeed      | float64             | ❌       |             |

# DataFiles4

**Properties**

| Name      | Type    | Required | Description |
| :-------- | :------ | :------- | :---------- |
| Id        | float64 | ❌       |             |
| Md5       | string  | ❌       |             |
| Mimetype  | string  | ❌       |             |
| Name      | string  | ❌       |             |
| S3Path    | string  | ❌       |             |
| ShortName | string  | ❌       |             |
| Size      | float64 | ❌       |             |
