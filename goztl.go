package goztl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.1.36"
	userAgent      = "goztl/" + libraryVersion
	mediaType      = "application/json"
)

// Client manages communication with Zentral API.
type Client struct {
	// HTTP client used to communicate with the Zentral API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	// Inventory
	JMESPathChecks    JMESPathChecksService
	MetaBusinessUnits MetaBusinessUnitsService
	Tags              TagsService
	Taxonomies        TaxonomiesService
	// MDM
	MDMArtifacts          MDMArtifactsService
	MDMBlueprints         MDMBlueprintsService
	MDMBlueprintArtifacts MDMBlueprintArtifactsService
	MDMFileVaultConfigs   MDMFileVaultConfigsService
	MDMProfiles           MDMProfilesService
	// Monolith
	MonolithCatalogs             MonolithCatalogsService
	MonolithConditions           MonolithConditionsService
	MonolithEnrollments          MonolithEnrollmentsService
	MonolithManifests            MonolithManifestsService
	MonolithManifestCatalogs     MonolithManifestCatalogsService
	MonolithManifestSubManifests MonolithManifestSubManifestsService
	MonolithSubManifests         MonolithSubManifestsService
	MonolithSubManifestPkgInfos  MonolithSubManifestPkgInfosService
	// Munki
	MunkiConfigurations MunkiConfigurationsService
	MunkiEnrollments    MunkiEnrollmentsService
	// Osquery
	OsqueryATC                OsqueryATCService
	OsqueryConfigurations     OsqueryConfigurationsService
	OsqueryConfigurationPacks OsqueryConfigurationPacksService
	OsqueryEnrollments        OsqueryEnrollmentsService
	OsqueryFileCategories     OsqueryFileCategoriesService
	OsqueryPacks              OsqueryPacksService
	OsqueryQueries            OsqueryQueriesService
	// Santa
	SantaConfigurations SantaConfigurationsService
	SantaEnrollments    SantaEnrollmentsService
	SantaRules          SantaRulesService

	// Zentral API token
	token string

	// Optional extra HTTP headers to set on every request to the API.
	headers map[string]string
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, maximum number of items to return
	Limit int `url:"limit,omitempty"`

	// For paginated result sets, starting position of the query in relation
	// to the complete set of unpaginated items
	Offset int `url:"offset,omitempty"`
}

// Response is a Zentral response. This wraps the standard http.Response returned from Zentral.
type Response struct {
	*http.Response
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// NewClient returns a new Zentral API client with the given base URL and API token.
func NewClient(httpClient *http.Client, bu string, token string, opts ...ClientOpt) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(bu)
	if err != nil {
		return nil, err
	}

	cleanToken := strings.Trim(strings.TrimSpace(token), "'")

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		token:     cleanToken,
	}
	// Inventory
	c.JMESPathChecks = &JMESPathChecksServiceOp{client: c}
	c.MetaBusinessUnits = &MetaBusinessUnitsServiceOp{client: c}
	c.Tags = &TagsServiceOp{client: c}
	c.Taxonomies = &TaxonomiesServiceOp{client: c}
	// MDM
	c.MDMArtifacts = &MDMArtifactsServiceOp{client: c}
	c.MDMBlueprints = &MDMBlueprintsServiceOp{client: c}
	c.MDMBlueprintArtifacts = &MDMBlueprintArtifactsServiceOp{client: c}
	c.MDMFileVaultConfigs = &MDMFileVaultConfigsServiceOp{client: c}
	c.MDMProfiles = &MDMProfilesServiceOp{client: c}
	// Monolith
	c.MonolithCatalogs = &MonolithCatalogsServiceOp{client: c}
	c.MonolithConditions = &MonolithConditionsServiceOp{client: c}
	c.MonolithEnrollments = &MonolithEnrollmentsServiceOp{client: c}
	c.MonolithManifests = &MonolithManifestsServiceOp{client: c}
	c.MonolithManifestCatalogs = &MonolithManifestCatalogsServiceOp{client: c}
	c.MonolithManifestSubManifests = &MonolithManifestSubManifestsServiceOp{client: c}
	c.MonolithSubManifests = &MonolithSubManifestsServiceOp{client: c}
	c.MonolithSubManifestPkgInfos = &MonolithSubManifestPkgInfosServiceOp{client: c}
	// Munki
	c.MunkiConfigurations = &MunkiConfigurationsServiceOp{client: c}
	c.MunkiEnrollments = &MunkiEnrollmentsServiceOp{client: c}
	// Osquery
	c.OsqueryATC = &OsqueryATCServiceOp{client: c}
	c.OsqueryConfigurations = &OsqueryConfigurationsServiceOp{client: c}
	c.OsqueryConfigurationPacks = &OsqueryConfigurationPacksServiceOp{client: c}
	c.OsqueryEnrollments = &OsqueryEnrollmentsServiceOp{client: c}
	c.OsqueryFileCategories = &OsqueryFileCategoriesServiceOp{client: c}
	c.OsqueryPacks = &OsqueryPacksServiceOp{client: c}
	c.OsqueryQueries = &OsqueryQueriesServiceOp{client: c}
	// Santa
	c.SantaConfigurations = &SantaConfigurationsServiceOp{client: c}
	c.SantaEnrollments = &SantaEnrollmentsServiceOp{client: c}
	c.SantaRules = &SantaRulesServiceOp{client: c}

	c.headers = make(map[string]string)

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}

// SetRequestHeaders sets optional HTTP headers on the client that are sent on each HTTP request.
func SetRequestHeaders(headers map[string]string) ClientOpt {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}
		return nil
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}

	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", mediaType)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", c.token))
	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}

	return &response
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := DoRequestWithClient(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Ensure the response body is fully read and closed
		// before we reconnect, so that we reuse the same TCPConnection.
		// Close the previous response's body. But read at least some of
		// the body so if it's small the underlying TCP connection will be
		// re-used. No need to check for errors: if it fails, the Transport
		// won't reuse it anyway.
		const maxBodySlurpSize = 2 << 10
		if resp.ContentLength == -1 || resp.ContentLength <= maxBodySlurpSize {
			io.CopyN(io.Discard, resp.Body, maxBodySlurpSize)
		}

		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

// DoRequestWithClient submits an HTTP request using the specified client.
func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
// If the API error response does not include the request ID in its body, the one from its header will be used.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		errorResponse.Message = string(data)
	}

	return errorResponse
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
