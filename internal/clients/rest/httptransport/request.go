package httptransport

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"torbox-sdk-go/internal/contenttypes"
	"torbox-sdk-go/internal/serialization"
	"torbox-sdk-go/internal/utils"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type paramMap struct {
	Key   string
	Value string
}

type Request struct {
	Context             context.Context
	Method              string
	Path                string
	Headers             map[string]string
	QueryParams         map[string]string
	PathParams          map[string]string
	Options             any
	Body                any
	Config              torboxapiconfig.Config
	ContentType         ContentType
	ResponseContentType ContentType
}

func NewRequest(ctx context.Context, method string, path string, config torboxapiconfig.Config) Request {
	return Request{
		Context:             ctx,
		Method:              method,
		Path:                path,
		Headers:             make(map[string]string),
		QueryParams:         make(map[string]string),
		PathParams:          make(map[string]string),
		Config:              config,
		ContentType:         ContentTypeJson,
		ResponseContentType: ContentTypeJson,
	}
}

func (r *Request) Clone() Request {
	if r == nil {
		return Request{
			Headers:     make(map[string]string),
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}

	clone := *r
	clone.PathParams = utils.CloneMap(r.PathParams)
	clone.Headers = utils.CloneMap(r.Headers)
	clone.QueryParams = utils.CloneMap(r.QueryParams)

	return clone
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}

func (r *Request) GetBaseUrl() string {
	return *r.Config.BaseUrl
}

func (r *Request) SetBaseUrl(baseUrl string) {
	r.Config.SetBaseUrl(baseUrl)
}

func (r *Request) GetPath() string {
	return r.Path
}

func (r *Request) SetPath(path string) {
	r.Path = path
}

func (r *Request) GetHeader(header string) string {
	return r.Headers[header]
}

func (r *Request) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *Request) GetPathParam(param string) string {
	return r.PathParams[param]
}

func (r *Request) SetPathParam(param string, value any) {
	r.PathParams[param] = fmt.Sprintf("%v", value)
}

func (r *Request) GetQueryParam(header string) string {
	return r.QueryParams[header]
}

func (r *Request) SetQueryParam(header string, value string) {
	r.QueryParams[header] = value
}

func (r *Request) GetOptions() any {
	return r.Options
}

func (r *Request) SetOptions(options any) {
	r.Options = options
}

func (r *Request) GetBody() any {
	return r.Body
}

func (r *Request) SetBody(body any) {
	r.Body = body
}

func (r *Request) GetContext() context.Context {
	return r.Context
}

func (r *Request) SetContext(ctx context.Context) {
	r.Context = ctx
}

func (r *Request) SetResponseContentType(contentType ContentType) {
	r.ResponseContentType = contentType
}

func (r *Request) CreateHttpRequest() (*http.Request, error) {
	requestUrl := r.getRequestUrl()

	requestBody, err := r.bodyToBytesReader()
	if err != nil {
		return nil, err
	}

	var httpRequest *http.Request
	if requestBody == nil {
		httpRequest, err = http.NewRequestWithContext(r.Context, r.Method, requestUrl, nil)
	} else {
		httpRequest, err = http.NewRequestWithContext(r.Context, r.Method, requestUrl, requestBody)
	}

	httpRequest.Header = r.getRequestHeaders()

	return httpRequest, err
}

func (r *Request) getRequestUrl() string {
	requestPath := r.Path
	for paramName, paramValue := range r.PathParams {
		placeholder := "{" + paramName + "}"
		requestPath = strings.ReplaceAll(requestPath, placeholder, url.PathEscape(paramValue))
	}

	requestOptions := ""
	params := r.getRequestQueryParams()
	if len(params) > 0 {
		requestOptions = fmt.Sprintf("?%s", params.Encode())
	}

	return *r.Config.BaseUrl + requestPath + requestOptions
}

func (r *Request) bodyToBytesReader() (*bytes.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

	if r.ContentType == ContentTypeJson {
		return contenttypes.ToJson(r.Body)
	} else if r.ContentType == ContentTypeMultipartFormData {
		bytesReader, contentTypeHeader, err := contenttypes.ToFormData(r.Body)
		if err != nil {
			return bytesReader, err
		}

		r.Headers["Content-Type"] = contentTypeHeader

		return bytesReader, err
	} else if r.ContentType == ContentTypeFormUrlEncoded {
		return contenttypes.ToFormUrlEncoded(r.Body)
	} else if r.ContentType == ContentTypeText {
		return contenttypes.ToText(r.Body)
	} else if r.ContentType == ContentTypeBinary || r.ContentType == ContentTypeImage {
		return contenttypes.ToBinary(r.Body)
	}

	return nil, fmt.Errorf("CreateRequestError: cannot parse request body. Content type not supported. ContentType: %v.", r.ContentType)
}

func (r *Request) getRequestQueryParams() url.Values {
	params := url.Values{}
	for key, value := range r.QueryParams {
		params.Add(key, value)
	}

	for _, p := range serializeQueryParams(r.Options) {
		params.Add(p.Key, p.Value)
	}

	return params
}

func (r *Request) getRequestHeaders() http.Header {
	headers := http.Header{}
	for key, value := range r.Headers {
		headers.Add(key, value)
	}

	serializeHeaderParams(headers, r.Options)

	return headers
}

func serializeHeaderParams(headers http.Header, obj any) {
	if obj == nil {
		return
	}

	values := utils.GetReflectValue(reflect.ValueOf(obj))
	for i := 0; i < values.NumField(); i++ {
		key, found := values.Type().Field(i).Tag.Lookup("headerParam")
		if shouldSkipField(found, values.Field(i)) {
			continue
		}

		field := values.Field(i)
		fieldValue := utils.GetReflectValue(field)
		fieldKind := utils.GetReflectKind(fieldValue.Type())
		switch fieldKind {
		case reflect.Array, reflect.Slice:
			for i := 0; i < field.Len(); i++ {
				headers.Add(key, fmt.Sprint(field.Index(i)))
			}
		case reflect.Struct:
			var serializedValue []string
			subValues := utils.GetReflectValue(fieldValue)
			for j := 0; j < subValues.NumField(); j++ {
				subKey, found := subValues.Type().Field(j).Tag.Lookup("headerParam")
				if !found {
					continue
				}
				if subKey == "" {
					subKey = subValues.Type().Field(j).Name // Default to field name if no tag
				}
				subField := subValues.Field(j)
				subFieldValue := utils.GetReflectValue(subField)
				serializedValue = append(serializedValue, subKey, fmt.Sprint(subFieldValue))
			}
			headers.Add(key, strings.Join(serializedValue, ","))
		default:
			headers.Add(key, fmt.Sprint(fieldValue))
		}
	}
}

func serializeQueryParams(obj any) []paramMap {
	queryParams := []paramMap{}

	if obj == nil {
		return queryParams
	}

	values := utils.GetReflectValue(reflect.ValueOf(obj))
	for i := 0; i < values.NumField(); i++ {
		key, found := values.Type().Field(i).Tag.Lookup("queryParam")
		if shouldSkipField(found, values.Field(i)) {
			continue
		}

		field := utils.GetReflectValue(values.Field(i))
		fieldKind := utils.GetReflectKind(field.Type())
		if fieldKind == reflect.Array || fieldKind == reflect.Slice {
			queryParams = append(queryParams, serializeArrayFieldToQueryParams(key, field, values.Type().Field(i))...)
		} else if fieldKind == reflect.Struct {
			objectParams := serialization.SerializeObject(key, field.Interface())
			for _, p := range objectParams {
				queryParams = append(queryParams, paramMap{Key: p.Key, Value: p.Value})
			}
		} else {
			queryParams = append(queryParams, paramMap{Key: key, Value: fmt.Sprint(field)})
		}
	}

	return queryParams
}

func serializeArrayFieldToQueryParams(key string, field reflect.Value, fieldInfo reflect.StructField) []paramMap {
	serializedParams := []paramMap{}
	serializedValues := []string{}
	for i := 0; i < field.Len(); i++ {
		serializedValues = append(serializedValues, fmt.Sprint(field.Index(i)))
	}

	if len(serializedValues) == 0 {
		return serializedParams
	}

	explode, found := fieldInfo.Tag.Lookup("explode")
	if !found || explode == "true" {
		for _, value := range serializedValues {
			serializedParams = append(serializedParams, paramMap{Key: key, Value: value})
		}
	} else {
		serializationStyle, _ := fieldInfo.Tag.Lookup("serializationStyle")
		delimiter := getDelimiterFromSerializationStyle(serializationStyle)
		joinedValues := strings.Join(serializedValues, delimiter)
		serializedParams = append(serializedParams, paramMap{Key: key, Value: joinedValues})
	}

	return serializedParams
}

func getDelimiterFromSerializationStyle(style string) string {
	delimiter := ","
	switch {
	case style == "form":
		delimiter = ","
	case style == "spaceDelimited":
		delimiter = " "
	case style == "pipeDelimited":
		delimiter = "|"
	}
	return delimiter
}

func shouldSkipField(found bool, field reflect.Value) bool {
	return !found || field.Type().Kind() == reflect.Pointer && field.IsNil()
}
