package validator

import (
	"sync"
	
	validatorModule "github.com/ej-you/go-utils/validator"
)


var once sync.Once
// структура для валидации входных данных
var validator *validatorModule.Validator


// получение структуры для валидации входных данных
func GetValidator() *validatorModule.Validator {
	once.Do(func() {
		validator = validatorModule.New()
	})
	return validator
}
