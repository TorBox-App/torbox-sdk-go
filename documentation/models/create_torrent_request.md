# CreateTorrentRequest

**Properties**

| Name     | Type   | Required | Description                                                                                                                                                                                      |
| :------- | :----- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| AllowZip | string | ❌       | Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger.                                                                      |
| AsQueued | string | ❌       | Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional.                                    |
| File     | []byte | ❌       | The torrent's torrent file. Optional.                                                                                                                                                            |
| Magnet   | string | ❌       | The torrent's magnet link. Optional.                                                                                                                                                             |
| Name     | string | ❌       | The name you want the torrent to be. Optional.                                                                                                                                                   |
| Seed     | string | ❌       | Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. |
