// Hand-written layer of the TorBox Go SDK (sdk-config/overlays/go). Not generated.

package torbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DefaultBaseURL is the TorBox API.
const DefaultBaseURL = "https://api.torbox.app"

// SDKVersion is the version of this SDK.
const SDKVersion = "9.3.1-20261125.4"

// Error is an answer with success=false or an error status. Code is the stable error code (e.g. BAD_TOKEN),
// Detail the message for people. Status is 0 when a 200 answer said success=false.
type Error struct {
	Status int
	Code   string
	Detail string
	Data   any
	Header http.Header
}

func (e *Error) Error() string { return strings.TrimSpace(fmt.Sprintf("%d %s: %s", e.Status, e.Code, e.Detail)) }

// Options configures New.
type Options struct {
	BaseURL    string        // default DefaultBaseURL
	Timeout    time.Duration // per attempt, default 30s
	MaxRetries int           // retries for 429/503, default 3; set -1 for none
	HTTPClient *http.Client  // base client; its Transport is wrapped with the retries
}

// Client is the generated APIClient (one field per API group: TorrentsAPI, UsenetAPI, UserAPI, ...) set up with
// the API key, retries and timeouts.
type Client struct {
	*APIClient
	apiKey string
}

// New returns a client for the TorBox API. apiKey is the key from https://torbox.app/settings.
func New(apiKey string, opts ...Options) *Client {
	o := Options{}
	if len(opts) > 0 {
		o = opts[0]
	}
	if o.BaseURL == "" {
		o.BaseURL = DefaultBaseURL
	}
	if o.Timeout == 0 {
		o.Timeout = 30 * time.Second
	}
	if o.MaxRetries == 0 {
		o.MaxRetries = 3
	}
	base := http.DefaultTransport
	if o.HTTPClient != nil && o.HTTPClient.Transport != nil {
		base = o.HTTPClient.Transport
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: strings.TrimRight(o.BaseURL, "/")}}
	cfg.UserAgent = "torbox-go/" + SDKVersion
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	cfg.HTTPClient = &http.Client{Timeout: o.Timeout, Transport: &retryTransport{base: base, maxRetries: max(o.MaxRetries, 0)}}
	return &Client{APIClient: NewAPIClient(cfg), apiKey: apiKey}
}

// APIKey is the key the client was created with. The download link endpoints take it as their Token:
//
//	link, err := torbox.Result(c.TorrentsAPI.RequestTorrentDownloadLink(ctx).Token(c.APIKey()).TorrentId(1).Execute())
func (c *Client) APIKey() string { return c.apiKey }

type envelope interface {
	GetSuccess() bool
	GetError() string
	GetDetail() string
}

// Result checks an Execute() result: an error status becomes *Error, and so does a 200 answer with
// success=false. Use resp.GetData() on what it returns.
//
//	res, err := torbox.Result(c.UserAPI.GetUserData(ctx).Execute())
func Result[E envelope](resp E, httpResp *http.Response, err error) (E, error) {
	if err != nil {
		return resp, AsError(err, httpResp)
	}
	if !resp.GetSuccess() {
		return resp, &Error{Status: 0, Code: resp.GetError(), Detail: resp.GetDetail()}
	}
	return resp, nil
}

// AsError turns an error of a generated call into *Error when the API answered with its error envelope (or a
// 422 validation error); other errors (network, timeouts) are returned unchanged.
func AsError(err error, httpResp *http.Response) error {
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return err
	}
	status := 0
	var header http.Header
	if httpResp != nil {
		status, header = httpResp.StatusCode, httpResp.Header
	}
	var body struct {
		Error  *string         `json:"error"`
		Detail json.RawMessage `json:"detail"`
		Data   any             `json:"data"`
	}
	if json.Unmarshal(apiErr.Body(), &body) != nil {
		return &Error{Status: status, Detail: apiErr.Error(), Header: header}
	}
	var detail string
	if json.Unmarshal(body.Detail, &detail) != nil {
		// FastAPI request validation (422) is {"detail": [...]}, not the envelope
		return &Error{Status: status, Code: "VALIDATION_ERROR", Detail: string(body.Detail), Header: header}
	}
	code := ""
	if body.Error != nil {
		code = *body.Error
	}
	return &Error{Status: status, Code: code, Detail: detail, Data: body.Data, Header: header}
}

// retryTransport retries 429 and 503, waiting for Retry-After or backing off exponentially.
type retryTransport struct {
	base       http.RoundTripper
	maxRetries int
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		r := req
		if attempt > 0 && req.Body != nil {
			if req.GetBody == nil {
				return nil, errors.New("torbox: cannot retry a request whose body cannot be replayed")
			}
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			r = req.Clone(req.Context())
			r.Body = body
		}
		resp, err := t.base.RoundTrip(r)
		if err != nil || (resp.StatusCode != 429 && resp.StatusCode != 503) || attempt >= t.maxRetries {
			return resp, err
		}
		wait := time.Duration(500*(1<<attempt)) * time.Millisecond
		if s, perr := strconv.Atoi(resp.Header.Get("Retry-After")); perr == nil && s > 0 {
			wait = time.Duration(s) * time.Second
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		select {
		case <-time.After(wait):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}
}
