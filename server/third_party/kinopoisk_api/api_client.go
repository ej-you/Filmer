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
	retryAttemps = 2 // кол-во попыток после первой неудачной
	minRetryWait = 2*time.Second // минимальное время ожидания между повторными попытками
	sendRequestTimeout = 3*time.Second // таймаут на запрос	
)


// интерфейс клиента Kinopoisk API
type KinopoiskAPI interface {
	SendGET(outStruct any) error
}


// структура для создания GET-запросов к API
type KinopoiskAPIGet struct {
	URL 		string
	APIKey 		string
	QueryParams map[string]string
	jsonify		jsonify.JSONify
}

// конструктор для типа интерфейса KinopoiskAPI
func NewKinopoiskAPI(url, apiKey string, queryParams map[string]string, jsonify jsonify.JSONify) KinopoiskAPI {
	return &KinopoiskAPIGet{
		URL: url,
		APIKey: apiKey,
		QueryParams: queryParams,
		jsonify: jsonify,
	}
}

// структура для парсинга JSON-ответа от API с ошибкой
//easyjson:json
type kinopoiskApiError struct {
	Message	string `json:"message"`
}
// парсинг ошибки из полученного ответа
func (this KinopoiskAPIGet) parseError(resp *http.Response) error {
	var rawErr kinopoiskApiError

	bytesErrorMessage, err := io.ReadAll(resp.Body)
	if err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("parse error: failed to read error answer: %v", err))
	}
	// декодирование ответа с ошибкой в структуру
	if err := this.jsonify.Unmarshal(bytesErrorMessage, &rawErr); err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("parse error: failed to decode error answer: %v", err))
	}
	// возврат обработанной ошибки
	return fmt.Errorf("parse error: %w", httpError.NewHTTPError(resp.StatusCode, rawErr.Message))
}

// отправка запроса и обработка ответа (outStruct - указатель на структуру)
func (this KinopoiskAPIGet) SendGET(outStruct any) error {
	var err error

	// создание запроса
	req, err := http.NewRequest("GET", this.URL, nil)
	if err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("failed to send request: %v", err))
	}
	// добавление API ключа в заголовок запроса
	req.Header.Set("X-API-KEY", this.APIKey)

	// добавление query-параметров
	queryParams := req.URL.Query()
	for k, v := range this.QueryParams {	
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	// клиент и его настройка на auto-retry
	client := retryHTTP.NewClient()
	client.HTTPClient = &http.Client{Timeout: sendRequestTimeout}
	client.RetryWaitMin = minRetryWait
	client.RetryMax = retryAttemps
	// client.Logger = settings.APILog

	// оборачивание запроса для auto-retry
	retryReq, err := retryHTTP.FromRequest(req)
	if err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("failed to wrap request for retry: %v", err))
	}
	// отправка запроса
	resp, err := client.Do(retryReq)
	if err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("failed to do request: %v", err))
	}
	defer resp.Body.Close()
	// если получен ответ с ошибкой
	if resp.StatusCode != 200 {
		// парсинг ошибки
		return this.parseError(resp)
	}

	// при успешном запросе - декодирование ответа в структуру
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("failed to read answer: %v", err))
	}
	if err := this.jsonify.Unmarshal(bytesData, outStruct); err != nil {
		return httpError.NewHTTPError(500, fmt.Sprintf("failed to decode answer: %v", err))
	}
	return nil
}
