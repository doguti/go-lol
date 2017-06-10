package lol

import(
	"net/url"
	"net/http"
	"sync"

	"context"
	"io"
	"encoding/json"
	"fmt"
	"bytes"
)

const (
	libraryVersion = "0.1"
	region         = "euw1"
	defaultBaseURL = "https://"+ region +".api.riotgames.com/lol/"
	profileIconURL = "http://ddragon.leagueoflegends.com/cdn/6.24.1/img/profileicon/"
	userAgent      = "go-lol/" + libraryVersion

	summonerURL    = "summoner/v3/summoners"
	championURL    = "platform/v3/champions"
)

// A Client manages communication with the LOL API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public LOL API, but can be
	// set to a domain endpoint to use with LOL RIOT. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	ProfileIconURL *url.URL

	// User agent used when communicating with the LOL API.
	UserAgent string

	keyLol string

	common service

	// Services used for talking to different parts of the LOL API.
	Summoners          *SummonerService
	Champions          *ChampionService


	//EndPoints
	SummonerURL string
	ChampionURL string
}

type service struct {
	client *Client
}

// NewClient returns a new LOL API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).

func NewClient(httpClient *http.Client, key string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	profileIconURL, _ := url.Parse(profileIconURL)
	c := &Client{
		client: httpClient,
		BaseURL: baseURL,
		ProfileIconURL: profileIconURL,
		UserAgent: userAgent,
		SummonerURL: summonerURL,
		ChampionURL: championURL}
	c.common.client = c
	c.keyLol = key
	c.Summoners = (*SummonerService)(&c.common)
	c.Champions = (*ChampionService)(&c.common)
	return c
}


// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%v?api_key=%s",urlStr,c.keyLol))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a GitHub API response. This wraps the standard http.Response
// returned from GitHub and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v ",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message)
}

//HTTP STATUS CODES
/*
400	Bad request
401	Unauthorized
403	Forbidden
404	Data not found
405	Method not allowed
415	Unsupported media type
429	Rate limit exceeded
500	Internal server error
502	Bad gateway
503	Service unavailable
504	Gateway timeout
*/

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	switch r.StatusCode {
	default:
		return errorResponse
	}
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)

	if err != nil {

		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}
	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("api_key")) > 0 {
		params.Set("api_key", "PRIVATE")
		uri.RawQuery = params.Encode()
	}
	return uri
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }