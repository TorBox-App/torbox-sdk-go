package torbox_test

// Runs against a Prism mock of the spec: TORBOX_MOCK_URL=http://127.0.0.1:4011 go test ./...

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	torbox "github.com/TorBox-App/torbox-sdk-go/v9"
)

func mock(t *testing.T) *torbox.Client {
	url := os.Getenv("TORBOX_MOCK_URL")
	if url == "" {
		t.Skip("TORBOX_MOCK_URL not set")
	}
	return torbox.New("00000000-0000-4000-8000-000000000001", torbox.Options{BaseURL: url, MaxRetries: -1, Timeout: 10 * time.Second})
}

func TestEnvelopeIsChecked(t *testing.T) {
	c := mock(t)
	res, err := torbox.Result(c.UserAPI.GetUserData(context.Background()).Execute())
	if err != nil {
		t.Fatal(err)
	}
	if res.GetData().Email == nil {
		t.Fatal("no email in data")
	}
}

func TestMultipartUpload(t *testing.T) {
	c := mock(t)
	f, err := os.CreateTemp(t.TempDir(), "*.torrent")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("d8:announce0:e")
	f.Seek(0, 0)
	res, err := torbox.Result(c.TorrentsAPI.CreateTorrent(context.Background()).File(f).Execute())
	if err != nil {
		t.Fatal(err)
	}
	if res.GetData().TorrentId == nil {
		t.Fatal("no torrent_id")
	}
}

func TestDownloadLinkTakesTheKeyAsToken(t *testing.T) {
	c := mock(t)
	if _, err := torbox.Result(c.TorrentsAPI.RequestTorrentDownloadLink(context.Background()).Token(c.APIKey()).TorrentId(1).FileId(0).Execute()); err != nil {
		t.Fatal(err)
	}
}

func TestErrorsAreTyped(t *testing.T) {
	url := os.Getenv("TORBOX_MOCK_URL")
	if url == "" {
		t.Skip("TORBOX_MOCK_URL not set")
	}
	c := torbox.New("", torbox.Options{BaseURL: url, MaxRetries: -1}) // no key -> 401
	_, err := torbox.Result(c.UserAPI.GetUserData(context.Background()).Execute())
	var e *torbox.Error
	if !errors.As(err, &e) || e.Status != 401 || e.Detail == "" {
		t.Fatalf("want *torbox.Error 401, got %#v", err)
	}
}

func TestRetriesOn429WithRetryAfter(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(429)
			w.Write([]byte(`{"success":false,"error":"RATE_LIMIT_EXCEEDED","detail":"300 per 1 minute","data":null}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":true,"error":null,"detail":"ok","data":{"email":"a@b.c"}}`))
	}))
	defer srv.Close()
	c := torbox.New("k", torbox.Options{BaseURL: srv.URL})
	res, err := torbox.Result(c.UserAPI.GetUserData(context.Background()).Execute())
	data := res.GetData()
	if err != nil || calls.Load() != 2 || data.GetEmail() != "a@b.c" {
		t.Fatalf("calls=%d err=%v", calls.Load(), err)
	}
}
