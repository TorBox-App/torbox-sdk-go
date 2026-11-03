package httptransport

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"example.com/main-api-sdk/internal/contenttypes"
	"example.com/main-api-sdk/internal/serialization"
	"example.com/main-api-sdk/internal/utils"
	"example.com/main-api-sdk/mainapisdkconfig"
)

type paramMap struct {
	Key   string
	Value string
}

// AuthScheme identifies which security scheme an operation requires.
// Used by auth handlers in the request chain to decide whether to apply
// their credentials on a per-operation basis: an operation that only
// declares apiKey security should not also receive Basic or Bearer
// headers just because those credentials happen to be configured.
type AuthScheme string

const (
	AuthSchemeAPIKey AuthScheme = "apiKey"
	AuthSchemeBasic  AuthScheme = "basic"
	AuthSchemeBearer AuthScheme = "bearer"
	AuthSchemeOAuth2 AuthScheme = "oauth2"
)

// ShouldApplyAuthScheme returns true when the given scheme should be
// applied for a request whose accepted schemes list is `schemes`.
//
//   - schemes == nil  → no per-operation preference recorded (e.g. caller
//     hand-constructed the Request); apply every configured credential
//     for backward compatibility with non-generated callers.
//   - schemes == []   → operation declared `security: []` (explicit
//     no-auth); apply nothing.
//   - schemes == [a, b] → operation accepts either scheme; the matching
//     handlers run.
func ShouldApplyAuthScheme(schemes []AuthScheme, scheme AuthScheme) bool {
	if schemes == nil {
		return true
	}
	for _, s := range schemes {
		if s == scheme {
			return true
		}
	}
	return false
}

// Request represents an HTTP request with all necessary configuration and parameters.
// Handles path/query/header serialization, content types, and authentication scopes.
type Request struct {
	Context             context.Context
	Method              string
	Path                string
	Headers             map[string]string
	QueryParams         map[string]string
	PathParams          map[string]string
	Options             any
	Body                any
	Config              mainapisdkconfig.Config
	ContentType         ContentType
	ResponseContentType ContentType
	// SecuritySchemes lists the auth schemes the operation accepts.
	// Per-handler dispatch uses ShouldApplyAuthScheme to decide whether
	// to apply its credentials. Defaults to nil for hand-built Requests,
	// which preserves the legacy "apply every configured credential"
	// behaviour.
	SecuritySchemes []AuthScheme
}

// NewRequest creates a new Request with default settings.
// Initializes maps for headers, query params, and path params, and sets default content types to JSON.
func NewRequest(ctx context.Context, method string, path string, config mainapisdkconfig.Config) Request {
	return Request{
		Context:             ctx,
		Method:              method,
		Path:                path,
		Headers:             make(map[string]string),
		QueryParams:         make(map[string]string),
		PathParams:          make(map[string]string),
		Config:              config,
		ContentType:         ContentTypeJSON,
		ResponseContentType: ContentTypeJSON,
	}
}

// Clone creates a deep copy of the Request including all maps.
// Returns a new Request with copied values to prevent mutation of the original.
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

func (r *Request) GetBaseURL() string {
	return r.Config.BaseURL
}

func (r *Request) SetBaseURL(baseURL string) {
	r.Config.SetBaseURL(baseURL)
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

func (r *Request) SetSecuritySchemes(schemes []AuthScheme) {
	r.SecuritySchemes = schemes
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

// validateBaseURLForHTTP checks that the configured base URL is usable for HTTP requests.
// It rejects empty values, unresolved {{...}} placeholders (e.g. from Postman collections),
// and non-absolute URLs.
func validateBaseURLForHTTP(baseURL string) error {
	s := strings.TrimSpace(baseURL)
	if s == "" {
		return fmt.Errorf("invalid base URL: empty; set BaseURL to a valid URL")
	}
	u, err := url.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid base URL %q: could not parse as URL: %w", s, err)
	}
	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid base URL %q: must be an absolute URL with scheme and host (for example https://api.example.com)", s)
	}
	return nil
}

func (r *Request) CreateHTTPRequest() (*http.Request, error) {
	if err := validateBaseURLForHTTP(r.Config.BaseURL); err != nil {
		return nil, err
	}

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
	if err != nil {
		return nil, fmt.Errorf("cannot build HTTP request for URL %q: %w", requestUrl, err)
	}

	httpRequest.Header = r.getRequestHeaders()
	serializeCookieParams(httpRequest, r.Options)

	return httpRequest, nil
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

	return strings.TrimRight(r.Config.BaseURL, "/") + requestPath + requestOptions
}

func (r *Request) bodyToBytesReader() (*bytes.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

	if r.ContentType == ContentTypeJSON {
		return contenttypes.ToJSON(r.Body)
	} else if r.ContentType == ContentTypeMultipartFormData {
		bytesReader, contentTypeHeader, err := contenttypes.ToFormData(r.Body)
		if err != nil {
			return bytesReader, err
		}

		r.Headers["Content-Type"] = contentTypeHeader

		return bytesReader, err
	} else if r.ContentType == ContentTypeFormUrlEncoded {
		return contenttypes.ToFormUrlEncoded(r.Body)
	} else if r.ContentType == ContentTypeXML {
		return contenttypes.ToXML(r.Body)
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

// serializeCookieParams reads fields tagged with `cookieParam:"name"` from obj
// and adds each as an http.Cookie to the request using req.AddCookie.
// Nil pointer fields (unset optional params) are skipped.
func serializeCookieParams(req *http.Request, obj any) {
	if obj == nil {
		return
	}

	values := utils.GetReflectValue(reflect.ValueOf(obj))
	for i := 0; i < values.NumField(); i++ {
		key, found := values.Type().Field(i).Tag.Lookup("cookieParam")
		if shouldSkipField(found, values.Field(i)) {
			continue
		}

		field := values.Field(i)
		fieldValue := utils.GetReflectValue(field)
		fieldKind := utils.GetReflectKind(fieldValue.Type())
		switch fieldKind {
		case reflect.Array, reflect.Slice:
			fieldInfo := values.Type().Field(i)
			explode, explodeFound := fieldInfo.Tag.Lookup("explode")
			if !explodeFound || explode == "true" {
				for j := 0; j < field.Len(); j++ {
					req.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(field.Index(j))})
				}
			} else {
				var serializedValues []string
				for j := 0; j < field.Len(); j++ {
					serializedValues = append(serializedValues, fmt.Sprint(field.Index(j)))
				}
				if len(serializedValues) > 0 {
					req.AddCookie(&http.Cookie{Name: key, Value: strings.Join(serializedValues, ",")})
				}
			}
		case reflect.Struct:
			if isNullableType(fieldValue.Type()) {
				if !fieldValue.FieldByName("IsNull").Bool() {
					innerValue := fieldValue.FieldByName("Value")
					innerKind := innerValue.Kind()
					innerFieldInfo := values.Type().Field(i)
					innerExplode, innerExplodeFound := innerFieldInfo.Tag.Lookup("explode")
					if innerKind == reflect.Slice || innerKind == reflect.Array {
						if !innerExplodeFound || innerExplode == "true" {
							for j := 0; j < innerValue.Len(); j++ {
								req.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(innerValue.Index(j))})
							}
						} else {
							var serializedValues []string
							for j := 0; j < innerValue.Len(); j++ {
								serializedValues = append(serializedValues, fmt.Sprint(innerValue.Index(j)))
							}
							if len(serializedValues) > 0 {
								req.AddCookie(&http.Cookie{Name: key, Value: strings.Join(serializedValues, ",")})
							}
						}
					} else {
						req.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(innerValue)})
					}
				}
				// null → omit the cookie entirely
			} else {
				req.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(fieldValue)})
			}
		default:
			req.AddCookie(&http.Cookie{Name: key, Value: fmt.Sprint(fieldValue)})
		}
	}
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
			fieldInfo := values.Type().Field(i)
			explode, explodeFound := fieldInfo.Tag.Lookup("explode")
			if !explodeFound || explode == "true" {
				for i := 0; i < field.Len(); i++ {
					headers.Add(key, fmt.Sprint(field.Index(i)))
				}
			} else {
				var serializedValues []string
				for i := 0; i < field.Len(); i++ {
					serializedValues = append(serializedValues, fmt.Sprint(field.Index(i)))
				}
				if len(serializedValues) > 0 {
					headers.Add(key, strings.Join(serializedValues, ","))
				}
			}
		case reflect.Struct:
			if isNullableType(fieldValue.Type()) {
				if !fieldValue.FieldByName("IsNull").Bool() {
					innerValue := fieldValue.FieldByName("Value")
					innerKind := innerValue.Kind()
					innerFieldInfo := values.Type().Field(i)
					innerExplode, innerExplodeFound := innerFieldInfo.Tag.Lookup("explode")
					if innerKind == reflect.Slice || innerKind == reflect.Array {
						if !innerExplodeFound || innerExplode == "true" {
							for j := 0; j < innerValue.Len(); j++ {
								headers.Add(key, fmt.Sprint(innerValue.Index(j)))
							}
						} else {
							var serializedValues []string
							for j := 0; j < innerValue.Len(); j++ {
								serializedValues = append(serializedValues, fmt.Sprint(innerValue.Index(j)))
							}
							if len(serializedValues) > 0 {
								headers.Add(key, strings.Join(serializedValues, ","))
							}
						}
					} else {
						headers.Add(key, fmt.Sprint(innerValue))
					}
				}
				// null → omit the header entirely
			} else {
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
			}
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
			if isNullableType(field.Type()) {
				queryParams = append(queryParams, serializeNullableAsQueryParams(key, field, values.Type().Field(i))...)
			} else {
				objectParams := serialization.SerializeObject(key, field.Interface())
				for _, p := range objectParams {
					queryParams = append(queryParams, paramMap{Key: p.Key, Value: p.Value})
				}
			}
		} else {
			queryParams = append(queryParams, paramMap{Key: key, Value: fmt.Sprint(field)})
		}
	}

	return queryParams
}

// isNullableType reports whether t is a param.Nullable[T] — a struct with exactly
// two exported fields named "Value" and "IsNull".
func isNullableType(t reflect.Type) bool {
	if t.Kind() != reflect.Struct || t.NumField() != 2 {
		return false
	}
	_, hasValue := t.FieldByName("Value")
	_, hasIsNull := t.FieldByName("IsNull")
	return hasValue && hasIsNull
}

// serializeNullableAsQueryParams serializes a param.Nullable[T] value as query params.
// Returns nil when the value is null or the inner slice is nil/empty (param omitted).
func serializeNullableAsQueryParams(key string, field reflect.Value, fieldInfo reflect.StructField) []paramMap {
	if field.FieldByName("IsNull").Bool() {
		return nil
	}
	innerValue := field.FieldByName("Value")
	innerKind := innerValue.Kind()
	if innerKind == reflect.Slice || innerKind == reflect.Array {
		if innerValue.IsNil() {
			return nil // nil slice = zero value / not set, omit
		}
		return serializeArrayFieldToQueryParams(key, innerValue, fieldInfo)
	}
	return []paramMap{{Key: key, Value: fmt.Sprint(innerValue)}}
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
