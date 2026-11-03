package httptransport

// ContentType represents the serialization format for request and response bodies.
// Used to determine how to serialize request data and deserialize response data.
type ContentType string

const (
	ContentTypeJSON              ContentType = "json"
	ContentTypeXML               ContentType = "xml"
	ContentTypePdf               ContentType = "pdf"
	ContentTypeImage             ContentType = "image"
	ContentTypeFile              ContentType = "file"
	ContentTypeBinary            ContentType = "binary"
	ContentTypeFormUrlEncoded    ContentType = "form"
	ContentTypeText              ContentType = "text"
	ContentTypeMultipartFormData ContentType = "multipartFormData"
)
