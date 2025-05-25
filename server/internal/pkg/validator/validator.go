// Package validator provides Validator interface to validate any struct.
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

var _ Validator = (*appValidator)(nil)

// Validator interface for HTTP requests to REST API.
type Validator interface {
	Validate(s any) error
}

// Validator implementation.
type appValidator struct {
	validatorInstance *govalidator.Validate
	translator        ut.Translator
}

func New() Validator {
	en := enlocale.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := govalidator.New(govalidator.WithRequiredStructEnabled())
	err := entrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return &appValidator{validate, trans}
}

// Validate validates given struct s (using pointer to this struct) and returns validate errors.
func (v appValidator) Validate(s any) error {
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

	// sort out errors and concat them into string
	transtaledStringSlice := make([]string, 0, len(rawTranstaledMap))
	for _, v := range rawTranstaledMap {
		transtaledStringSlice = append(transtaledStringSlice, strings.ToLower(v))
	}

	errMsg := strings.Join(transtaledStringSlice, " && ")
	return httperror.New(http.StatusBadRequest,
		errMsg, fmt.Errorf("validate error"))
}
