package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/pkg/utils"
	"Filmer/client/internal/repository"
)

const apiRequestTimeout = 5 * time.Second // timeout for requests to REST API

// REST API client implementation
type restAPIClient struct {
	cfg            *config.Config
	requestTimeout time.Duration
}

// REST API constructor
func NewRestAPI(cfg *config.Config) repository.RestAPI {
	return &restAPIClient{
		cfg:            cfg,
		requestTimeout: apiRequestTimeout,
	}
}

func (api restAPIClient) GetMovie(authToken string, kinopoiskID int) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)}
	url := fmt.Sprintf("%s/api/v1/films/full-info/%d", api.cfg.RestAPI.Host, kinopoiskID)
	if err := api.sendGET(url, headers, nil, apiResp); err != nil {
		return nil, fmt.Errorf("get full movie info using rest api: send get request: %w", err)
	}
	return apiResp, nil
}

// Get stared, want and watched user movies
func (api restAPIClient) GetCategory(authToken, category string, queryParams repository.CategoryUserMoviesIn) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)}
	url := fmt.Sprintf("%s/api/v1/films/%s", api.cfg.RestAPI.Host, category)
	if err := api.sendGET(url, headers, queryParams, apiResp); err != nil {
		return nil, fmt.Errorf("get %s user movies using rest api: send get request: %w", category, err)
	}
	return apiResp, nil
}

func (api restAPIClient) PostCategory(authToken, category, movieID string) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)}
	url := fmt.Sprintf("%s/api/v1/films/%s/%s", api.cfg.RestAPI.Host, movieID, category)
	if err := api.sendPOST(url, headers, nil, apiResp); err != nil {
		return nil, fmt.Errorf("set %s user movie using rest api: send post request: %w", category, err)
	}
	return apiResp, nil
}

// Search movies using keyword (with paginaiton)
func (api restAPIClient) GetSearchMovies(authToken string, queryParams *repository.SearchMoviesIn) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)}
	query := map[string][]string{
		"q":    []string{queryParams.Keyword},
		"page": []string{strconv.Itoa(queryParams.Page)},
	}
	if err := api.sendGET(api.cfg.RestAPI.Host+"/api/v1/kinopoisk/films/search", headers, query, apiResp); err != nil {
		return nil, fmt.Errorf("search movies using rest api: send get request: %w", err)
	}
	return apiResp, nil
}

// Login user
func (api restAPIClient) Login(body repository.AuthIn) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	if err := api.sendPOST(api.cfg.RestAPI.Host+"/api/v1/user/login", nil, body, apiResp); err != nil {
		return nil, fmt.Errorf("login using rest api: send post request: %w", err)
	}
	return apiResp, nil
}

// Sign up user
func (api restAPIClient) SignUp(body repository.AuthIn) (*repository.APIResponse, error) {
	apiResp := new(repository.APIResponse)
	if err := api.sendPOST(api.cfg.RestAPI.Host+"/api/v1/user/sign-up", nil, body, apiResp); err != nil {
		return nil, fmt.Errorf("sign up using rest api: send post request: %w", err)
	}
	return apiResp, nil
}

// Logout user
func (api restAPIClient) Logout(authToken string) error {
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authToken)}
	if err := api.sendPOST(api.cfg.RestAPI.Host+"/api/v1/user/logout", headers, nil, nil); err != nil {
		return fmt.Errorf("logout using rest api: send post request: %w", err)
	}
	return nil
}

// Parse error from response
func (api restAPIClient) parseError(resp *http.Response) error {
	fiberErr := new(fiber.Error)
	// read error response
	bytesErrorMessage, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read error response: %w", err)
	}
	// decode response to struct
	if err := json.Unmarshal(bytesErrorMessage, &fiberErr); err != nil {
		return fmt.Errorf("decode error response: %w", err)
	}
	fiberErr.Code = resp.StatusCode
	// return processed error
	return fiberErr
}

// Send prepared request to REST API
func (api restAPIClient) sendRequest(req *http.Request, outStruct *repository.APIResponse) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	// if resp code is not 2xx
	if !utils.StatusCode2xx(resp.StatusCode) {
		return fmt.Errorf("got api error: %w", api.parseError(resp))
	}
	if outStruct == nil {
		return nil
	}
	// read response body and decode it to outstruct
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}
	err = json.Unmarshal(bytesData, &outStruct)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	return nil
}

// Send GET request to REST API
func (api restAPIClient) sendGET(url string, headers map[string]string, queryParams map[string][]string, outStruct *repository.APIResponse) error {
	reqContext, cancel := context.WithTimeout(context.Background(), api.requestTimeout)
	defer cancel()

	// create request
	req, err := http.NewRequestWithContext(reqContext, "GET", url, http.NoBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	// add headers
	// add headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if queryParams != nil {
		// add query-params
		query := req.URL.Query()
		for k, v := range queryParams {
			for _, val := range v {
				query.Add(k, val)
			}
		}
		req.URL.RawQuery = query.Encode()
	}
	// send prepared request
	return api.sendRequest(req, outStruct)
}

// Send POST request to REST API
func (api restAPIClient) sendPOST(url string, headers map[string]string, body any, outStruct *repository.APIResponse) error {
	reqContext, cancel := context.WithTimeout(context.Background(), api.requestTimeout)
	defer cancel()

	var bodyReader io.Reader
	if body == nil {
		bodyReader = http.NoBody
	} else {
		// create io.Reader from struct body
		bytesBody, err := json.Marshal(body)

		if err != nil {
			return fmt.Errorf("process body: %w", err)
		}
		bodyReader = bytes.NewBuffer(bytesBody)
	}

	// create request
	req, err := http.NewRequestWithContext(reqContext, "POST", url, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	// add headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// send prepared request
	return api.sendRequest(req, outStruct)
}
