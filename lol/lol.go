package lol

import(
	"net/url"
	"net/http"
	"sync"

)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://api.github.com/"
	userAgent      = "go-lol/" + libraryVersion
)



// A Client manages communication with the LOL API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public LOL API, but can be
	// set to a domain endpoint to use with LOL RIOT. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the LOL API.
	UserAgent string
}

// NewClient returns a new LOL API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	return c
}