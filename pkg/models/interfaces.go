package attom

import (
	"io"
	"net/http"
)

// HTTPClient interface defines the methods we need from an
// HTTP client.  This allows us to mock HTTP clients for testing,
// and the use of alternative HTTP clients for fetching data.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Post(url, contentType string, body io.Reader) (*http.Response, error)
	Put(url, contentType string, body io.Reader) (*http.Response, error)
	Patch(url, contentType string, body io.Reader) (*http.Response, error)
	Delete(url string) (*http.Response, error)
	Head(url string) (*http.Response, error)
	Options(url string) (*http.Response, error)
}
