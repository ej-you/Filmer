// Package kinopoisk provides interface for sending requests and
// parse responses from kinopoisk API.
package kinopoisk

import (
	"fmt"
	"io"
	"net/http"
	"time"

	retryHTTP "github.com/hashicorp/go-retryablehttp"

	"Filmer/server/internal/pkg/httperror"
	"Filmer/server/internal/pkg/jsonify"
)

const (
	retryAttempts      = 2               // attempts amount after first failed
	minRetryWait       = 2 * time.Second // min wait time between retries
	sendRequestTimeout = 3 * time.Second // request timeout
)

// API client interface.
type API interface {
	SendGET(outStruct any) error
}

// API client for GET-requests.
type apiGet struct {
	url         string
	apiKey      string
	queryParams map[string]string
	jsonify     jsonify.JSONify
}

// Returns API interface.
func NewAPI(url, apiKey string, queryParams map[string]string, jsonify jsonify.JSONify) API {
	return &apiGet{
		url:         url,
		apiKey:      apiKey,
		queryParams: queryParams,
		jsonify:     jsonify,
	}
}

// struct for parsing error JSON-response from API
//
//easyjson:json
type apiError struct {
	Message string `json:"message"`
}

// Parse error from response.
func (a apiGet) parseError(resp *http.Response) error {
	var rawErr apiError

	// if 404 error code
	if resp.StatusCode == http.StatusNotFound {
		return httperror.New(http.StatusNotFound,
			"movie not found", fmt.Errorf("got not found error"))
	}

	bytesErrorMessage, err := io.ReadAll(resp.Body)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"parse error: failed to read error answer", err)
	}
	// decode response to struct
	if err := a.jsonify.Unmarshal(bytesErrorMessage, &rawErr); err != nil {
		return httperror.New(http.StatusInternalServerError,
			"parse error: failed to decode error answer", err)
	}
	// return processed error
	return httperror.New(resp.StatusCode,
		rawErr.Message, fmt.Errorf("parsed error"))
}

// Send request and process response (outStruct - pointer to struct).
func (a apiGet) SendGET(outStruct any) error {
	var err error

	// create request
	req, err := http.NewRequest("GET", a.url, http.NoBody)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to send request", err)
	}
	// add API key to request headers
	req.Header.Set("X-API-KEY", a.apiKey)

	// add query-params
	queryParams := req.URL.Query()
	for k, v := range a.queryParams {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	// http client and auto-retry set up
	client := retryHTTP.NewClient()
	client.HTTPClient = &http.Client{Timeout: sendRequestTimeout}
	client.RetryWaitMin = minRetryWait
	client.RetryMax = retryAttempts

	// wrap request for auto-retry
	retryReq, err := retryHTTP.FromRequest(req)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to wrap request for retry", err)
	}
	// send request
	resp, err := client.Do(retryReq)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to do request", err)
	}
	defer resp.Body.Close()
	// if error reqponse was received
	if resp.StatusCode != http.StatusOK {
		return a.parseError(resp)
	}

	// if request is success - decode reqponse to struct
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to read answer", err)
	}
	if err := a.jsonify.Unmarshal(bytesData, outStruct); err != nil {
		return httperror.New(http.StatusInternalServerError,
			"failed to decode answer", err)
	}
	return nil
}
