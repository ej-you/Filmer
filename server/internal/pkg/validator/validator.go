package validator

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	enlocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	govalidator "github.com/go-playground/validator/v10"
	entrans "github.com/go-playground/validator/v10/translations/en"

	"Filmer/server/internal/pkg/httperror"
)

// Validator interface for HTTP requests to REST API
type Validator interface {
	Validate(s any) error
}

// Validator implementation
type restValidator struct {
	validatorInstance *govalidator.Validate
	translator        ut.Translator
}

// Validator constructor
func NewValidator() Validator {
	en := enlocale.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := govalidator.New(govalidator.WithRequiredStructEnabled())
	err := entrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return &restValidator{validate, trans}
}

// Validate given struct s (using pointer to this struct) with error handling to HTTPError
func (v restValidator) Validate(s any) error {
	err := v.validatorInstance.Struct(s)
	if err == nil { // NOT err
		return nil
	}

	// assert error to validatorModule.ValidationErrors
	var validateErrors govalidator.ValidationErrors
	if !errors.As(err, &validateErrors) {
		return err
	}
	// handle error messages
	rawTranstaledMap := validateErrors.Translate(v.translator)
	// for concat string
	transtaledStringSlice := make([]string, 0, len(rawTranstaledMap))
	// sort out errors and concat them into string
	var tempSlice []string
	var key string
	for k, v := range rawTranstaledMap {
		tempSlice = strings.Split(k, ".")
		key = tempSlice[len(tempSlice)-1]

		transtaledStringSlice = append(transtaledStringSlice, key+": "+v)
	}
	return httperror.NewHTTPError(http.StatusBadRequest, strings.Join(transtaledStringSlice, " | "), fmt.Errorf("validate error"))
}
