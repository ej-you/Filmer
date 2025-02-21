package kinopoisk_api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mailru/easyjson"
	fiber "github.com/gofiber/fiber/v2"
)


const sendRequestTimeout = 3*time.Second


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
		return fiber.NewError(500, fmt.Sprintf("failed to decode error answer: request to %q: %v", api.URL, err))
	}
	// возврат Fiber-ошибки
	return fiber.NewError(response.StatusCode, rawErr.Message)
}

// отправка запроса и обработка ответа (outStruct - указатель на структуру)
func (api *apiGetRequest) sendRequest(outStruct easyjson.Unmarshaler) error {
	var err error

	client := &http.Client{Timeout: sendRequestTimeout}

	// создание запроса
	req, err := http.NewRequest("GET", api.URL, nil)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to send request to %q: %v", api.URL, err))
	}
	// добавление API ключа в заголовок запроса
	req.Header.Set("X-API-KEY", api.APIKey)

	// добавление query-параметров
	queryParams := req.URL.Query()
	for k, v := range api.QueryParams {	
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	// отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to do request to %q: %v", api.URL, err))
	}
	defer resp.Body.Close()
	// если получен ответ с ошибкой
	if resp.StatusCode != 200 {
		// парсинг ошибки
		return api.parseError(resp)
	}

	// при успешном запросе - декодирование ответа в структуру
	if err := easyjson.UnmarshalFromReader(resp.Body, outStruct); err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to decode answer: request to %q: %v", api.URL, err))
	}
	return nil
}
