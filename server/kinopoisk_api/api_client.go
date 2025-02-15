package kinopoisk_api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mailru/easyjson"
	fiber "github.com/gofiber/fiber/v2"

	"server/settings"
)


const apiUrl = "https://kinopoiskapiunofficial.tech/api"
const sendRequestTimeout = 3*time.Second


// структура для парсинга JSON-ответа от API с ошибкой
//easyjson:json
type KinopoiskApiError struct {
	Message	string `json:"message"`
}

// парсинг ошибки из полученного ответа
func parseError(url string, response *http.Response) error {
	var rawErr KinopoiskApiError

	// декодирование ответа с ошибкой в структуру
	if err := easyjson.UnmarshalFromReader(response.Body, &rawErr); err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to decode error answer: request to %q: %v", url, err))
	}
	// возврат Fiber-ошибки
	return fiber.NewError(response.StatusCode, rawErr.Message)
}


// отправка запроса и обработка ответа (outStruct - указатель на структуру)
func sendRequest(req *http.Request, url string, outStruct easyjson.Unmarshaler) error {
// func sendRequest(req *http.Request, url string, outStruct any) error {
	var err error

	client := &http.Client{Timeout: sendRequestTimeout}
	// добавление API ключа в заголовок запроса
	req.Header.Set("X-API-KEY", settings.KinopoiskApiKey)

	// отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to do request to %q: %v", url, err))
	}
	defer resp.Body.Close()
	// если получен ответ с ошибкой
	if resp.StatusCode != 200 {
		// парсинг ошибки
		return parseError(url, resp)
	}

	// если для ответа не передана структура, то пропускаем парсинг ответа в неё
	if outStruct == nil {
		return nil
	}
	// при успешном запросе - декодирование ответа в структуру
	if err := easyjson.UnmarshalFromReader(resp.Body, outStruct); err != nil {
		return fiber.NewError(500, fmt.Sprintf("failed to decode answer: request to %q: %v", url, err))
	}
	return nil
}
