package validator

import (
	"strings"

	validatorModule "github.com/go-playground/validator/v10"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTrans "github.com/go-playground/validator/v10/translations/en"

	httpError "Filmer/server/pkg/http_error"
)


// интерфейс валидатора для HTTP-запросов к REST API
type Validator interface {
	Validate(s any) error
}


// структура для валидации входных данных
type restValidator struct {
	validatorInstance	*validatorModule.Validate
	translator			ut.Translator
}

// конструктор для валидатора
func NewValidator() Validator {
	en := enLocale.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validatorModule.New(validatorModule.WithRequiredStructEnabled())
	enTrans.RegisterDefaultTranslations(validate, trans)

	return &restValidator{validate, trans}
}

// валидация переданной через указатель структуры s с обработкой ошибок валидации в HTTPError
func (this restValidator) Validate(s any) error {
	err := this.validatorInstance.Struct(s)
	if err == nil { // NOT err
		return nil
	}

	// приводим ошибку к validatorModule.ValidationErrors
	validateErrors := err.(validatorModule.ValidationErrors)
	// обработка сообщений ошибки
	rawTranstaledMap := validateErrors.Translate(this.translator)
	// для объединенной строки
	transtaledStringSlice := make([]string, 0, len(rawTranstaledMap))
	// перебор ошибок и конкатенация их в строку
	var tempSlice []string
	var key string
	for k, v := range rawTranstaledMap {
		tempSlice = strings.Split(k, ".")
		key = tempSlice[len(tempSlice) - 1]

		transtaledStringSlice = append(transtaledStringSlice, key+": "+v)
	}
	return httpError.NewHTTPError(400, strings.Join(transtaledStringSlice, " | "))
}
