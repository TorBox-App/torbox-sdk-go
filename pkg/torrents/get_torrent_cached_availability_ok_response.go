package torrents

import (
	"torbox-sdk-go/internal/marshal"
	"torbox-sdk-go/internal/unmarshal"
)

type GetTorrentCachedAvailabilityOkResponse struct {
	Data    *GetTorrentCachedAvailabilityOkResponseData `json:"data,omitempty"`
	Detail  *string                                     `json:"detail,omitempty"`
	Success *bool                                       `json:"success,omitempty"`
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetData(data GetTorrentCachedAvailabilityOkResponseData) {
	g.Data = &data
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetData() *GetTorrentCachedAvailabilityOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetTorrentCachedAvailabilityOkResponseData struct {
	Data1Array []Data1
	Data2      *Data2
	AnyValue   any
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetData1Array(data1Array []Data1) {
	g.Data1Array = data1Array
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetData1Array() []Data1 {
	if g == nil {
		return nil
	}
	return g.Data1Array
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetData2(data2 Data2) {
	g.Data2 = &data2
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetData2() *Data2 {
	if g == nil {
		return nil
	}
	return g.Data2
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetAnyValue(anyValue any) {
	g.AnyValue = anyValue
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetAnyValue() any {
	if g == nil {
		return nil
	}
	return g.AnyValue
}

func (g GetTorrentCachedAvailabilityOkResponseData) MarshalJSON() ([]byte, error) {
	return marshal.FromComplexObject(g)
}

func (g *GetTorrentCachedAvailabilityOkResponseData) UnmarshalJSON(data []byte) error {
	return unmarshal.ToComplexObject[GetTorrentCachedAvailabilityOkResponseData](data, g)
}

type Data1 struct {
	Files []Data1Files `json:"files,omitempty"`
	Hash  *string      `json:"hash,omitempty"`
	Name  *string      `json:"name,omitempty"`
	Size  *float64     `json:"size,omitempty"`
}

func (d *Data1) SetFiles(files []Data1Files) {
	d.Files = files
}

func (d *Data1) GetFiles() []Data1Files {
	if d == nil {
		return nil
	}
	return d.Files
}

func (d *Data1) SetHash(hash string) {
	d.Hash = &hash
}

func (d *Data1) GetHash() *string {
	if d == nil {
		return nil
	}
	return d.Hash
}

func (d *Data1) SetName(name string) {
	d.Name = &name
}

func (d *Data1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Data1) SetSize(size float64) {
	d.Size = &size
}

func (d *Data1) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

type Data1Files struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (d *Data1Files) SetName(name string) {
	d.Name = &name
}

func (d *Data1Files) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Data1Files) SetSize(size float64) {
	d.Size = &size
}

func (d *Data1Files) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

type Data2 struct {
	Xxxxxxxxxxxxxxxxxx  *Xxxxxxxxxxxxxxxxxx  `json:"XXXXXXXXXXXXXXXXXX,omitempty"`
	Xxxxxxxxxxxxxxxxxxx *Xxxxxxxxxxxxxxxxxxx `json:"XXXXXXXXXXXXXXXXXXX,omitempty"`
}

func (d *Data2) SetXxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxxxxxx Xxxxxxxxxxxxxxxxxx) {
	d.Xxxxxxxxxxxxxxxxxx = &xxxxxxxxxxxxxxxxxx
}

func (d *Data2) GetXxxxxxxxxxxxxxxxxx() *Xxxxxxxxxxxxxxxxxx {
	if d == nil {
		return nil
	}
	return d.Xxxxxxxxxxxxxxxxxx
}

func (d *Data2) SetXxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxxxxxxx Xxxxxxxxxxxxxxxxxxx) {
	d.Xxxxxxxxxxxxxxxxxxx = &xxxxxxxxxxxxxxxxxxx
}

func (d *Data2) GetXxxxxxxxxxxxxxxxxxx() *Xxxxxxxxxxxxxxxxxxx {
	if d == nil {
		return nil
	}
	return d.Xxxxxxxxxxxxxxxxxxx
}

type Xxxxxxxxxxxxxxxxxx struct {
	Hash *string  `json:"hash,omitempty"`
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (x *Xxxxxxxxxxxxxxxxxx) SetHash(hash string) {
	x.Hash = &hash
}

func (x *Xxxxxxxxxxxxxxxxxx) GetHash() *string {
	if x == nil {
		return nil
	}
	return x.Hash
}

func (x *Xxxxxxxxxxxxxxxxxx) SetName(name string) {
	x.Name = &name
}

func (x *Xxxxxxxxxxxxxxxxxx) GetName() *string {
	if x == nil {
		return nil
	}
	return x.Name
}

func (x *Xxxxxxxxxxxxxxxxxx) SetSize(size float64) {
	x.Size = &size
}

func (x *Xxxxxxxxxxxxxxxxxx) GetSize() *float64 {
	if x == nil {
		return nil
	}
	return x.Size
}

type Xxxxxxxxxxxxxxxxxxx struct {
	Hash *string  `json:"hash,omitempty"`
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (x *Xxxxxxxxxxxxxxxxxxx) SetHash(hash string) {
	x.Hash = &hash
}

func (x *Xxxxxxxxxxxxxxxxxxx) GetHash() *string {
	if x == nil {
		return nil
	}
	return x.Hash
}

func (x *Xxxxxxxxxxxxxxxxxxx) SetName(name string) {
	x.Name = &name
}

func (x *Xxxxxxxxxxxxxxxxxxx) GetName() *string {
	if x == nil {
		return nil
	}
	return x.Name
}

func (x *Xxxxxxxxxxxxxxxxxxx) SetSize(size float64) {
	x.Size = &size
}

func (x *Xxxxxxxxxxxxxxxxxxx) GetSize() *float64 {
	if x == nil {
		return nil
	}
	return x.Size
}
