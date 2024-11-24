package httptransport

type ContentType string

const (
	ContentTypeJson              ContentType = "json"
	ContentTypeXml               ContentType = "xml"
	ContentTypePdf               ContentType = "pdf"
	ContentTypeImage             ContentType = "image"
	ContentTypeFile              ContentType = "file"
	ContentTypeBinary            ContentType = "binary"
	ContentTypeFormUrlEncoded    ContentType = "form"
	ContentTypeText              ContentType = "text"
	ContentTypeMultipartFormData ContentType = "multipartFormData"
)
