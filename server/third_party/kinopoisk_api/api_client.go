package kinopoisk_api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	retryHTTP "github.com/hashicorp/go-retryablehttp"

	"Filmer/server/pkg/jsonify"
	httpError "Filmer/server/pkg/http_error"
)


const (
	retryAttemps = 2 // attemps amount after first failed
	minRetryWait = 2*time.Second // min wait time between retries
	sendRequestTimeout = 3*time.Second // request timeout	
)


// Kinopoisk API client interface
type KinopoiskAPI interface {
	SendGET(outStruct any) error
}


// Kinopoisk API client for GET-requests
type KinopoiskAPIGet struct {
	url 		string
	apiKey 		string
	queryParams map[string]string
	jsonify		jsonify.JSONify
}

// KinopoiskAPI constructor
func NewKinopoiskAPI(url, apiKey string, queryParams map[string]string, jsonify jsonify.JSONify) KinopoiskAPI {
	return &KinopoiskAPIGet{
		url: url,
		apiKey: apiKey,
		queryParams: queryParams,
		jsonify: jsonify,
	}
}

// struct for parsing error JSON-response from API
//easyjson:json
type kinopoiskApiError struct {
	Message	string `json:"message"`
}
// Parse error from response
func (this KinopoiskAPIGet) parseError(resp *http.Response) error {
	var rawErr kinopoiskApiError

	bytesErrorMessage, err := io.ReadAll(resp.Body)
	if err != nil {
		return httpError.NewHTTPError(500, "parse error: failed to read error answer", err)
	}
	// decode response to struct
	if err := this.jsonify.Unmarshal(bytesErrorMessage, &rawErr); err != nil {
		return httpError.NewHTTPError(500, "parse error: failed to decode error answer", err)
	}
	// return processed error
	return httpError.NewHTTPError(resp.StatusCode, rawErr.Message, fmt.Errorf("parsed error"))
}

// Send request and process response (outStruct - pointer to struct)
func (this KinopoiskAPIGet) SendGET(outStruct any) error {
	var err error

	// create request
	req, err := http.NewRequest("GET", this.url, nil)
	if err != nil {
		return httpError.NewHTTPError(500, "failed to send request", err)
	}
	// add API key to request headers
	req.Header.Set("X-API-KEY", this.apiKey)

	// add query-params
	queryParams := req.URL.Query()
	for k, v := range this.queryParams {	
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	// http client and auto-retry set up
	client := retryHTTP.NewClient()
	client.HTTPClient = &http.Client{Timeout: sendRequestTimeout}
	client.RetryWaitMin = minRetryWait
	client.RetryMax = retryAttemps
	// client.Logger = settings.APILog

	// wrap request for auto-retry
	retryReq, err := retryHTTP.FromRequest(req)
	if err != nil {
		return httpError.NewHTTPError(500, "failed to wrap request for retry", err)
	}
	// send request
	resp, err := client.Do(retryReq)
	if err != nil {
		return httpError.NewHTTPError(500, "failed to do request", err)
	}
	defer resp.Body.Close()
	// if error reqponse was received
	if resp.StatusCode != 200 {
		return this.parseError(resp)
	}

	// if request is success - decode reqponse to struct
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		return httpError.NewHTTPError(500, "failed to read answer", err)
	}
	if err := this.jsonify.Unmarshal(bytesData, outStruct); err != nil {
		return httpError.NewHTTPError(500, "failed to decode answer", err)
	}
	return nil
}
