package kinopoisk_api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mailru/easyjson"
	fiber "github.com/gofiber/fiber/v2"
	retryHTTP "github.com/hashicorp/go-retryablehttp"

	"server/settings"
)


const retryAttemps = 2 // кол-во попыток после первой неудачной
const minRetryWait = 2*time.Second // минимальное время ожидания между повторными попытками
const sendRequestTimeout = 3*time.Second // таймаут на запрос


// структура для парсинга JSON-ответа от API с ошибкой
//easyjson:json
type kinopoiskApiError struct {
	Message	string `json:"message"`
}

// структура для создания запроса к API
type apiGetRequest struct {
	URL 		string
	APIKey 		string
	QueryParams map[string]string
}

// парсинг ошибки из полученного ответа
func (api *apiGetRequest) parseError(response *http.Response) error {
	var rawErr kinopoiskApiError

	// декодирование ответа с ошибкой в структуру
	if err := easyjson.UnmarshalFromReader(response.Body, &rawErr); err != nil {
		return fiber.NewError(500, fmt.Sprintf("parse error: failed to decode error answer: %v", err))
	}
	// возврат обработанной ошибки
	return fmt.Errorf("parse error: %w", fiber.NewError(response.StatusCode, rawErr.Message))
}

// отправка запроса и обработка ответа (outStruct - указатель на структуру)
func (api *apiGetRequest) sendRequest(outStruct easyjson.Unmarshaler) error {
	var err error

	// создание запроса
	req, err := http.NewRequest("GET", api.URL, nil)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to send request: %v", err))
	}
	// добавление API ключа в заголовок запроса
	req.Header.Set("X-API-KEY", api.APIKey)

	// добавление query-параметров
	queryParams := req.URL.Query()
	for k, v := range api.QueryParams {	
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	// клиент и его настройка на auto-retry
	client := retryHTTP.NewClient()
	client.HTTPClient = &http.Client{Timeout: sendRequestTimeout}
	client.RetryWaitMin = minRetryWait
	client.RetryMax = retryAttemps
	client.Logger = settings.APILog

	// оборачивание запроса для auto-retry
	retryReq, err := retryHTTP.FromRequest(req)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to wrap request for retry: %v", err))
	}
	// отправка запроса
	resp, err := client.Do(retryReq)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to do request: %v", err))
	}
	defer resp.Body.Close()
	// если получен ответ с ошибкой
	if resp.StatusCode != 200 {
		// парсинг ошибки
		return api.parseError(resp)
	}

	// при успешном запросе - декодирование ответа в структуру
	if err := easyjson.UnmarshalFromReader(resp.Body, outStruct); err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to decode answer: %v", err))
	}
	return nil
}
