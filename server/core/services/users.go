package services

import (
	"golang.org/x/crypto/bcrypt"
	fiber "github.com/gofiber/fiber/v2"
)


// кодирование пароля в хэш
func EncodePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// возвращаем 400, потому что скорее всего ошибка длины пароля
		return nil, fiber.NewError(400, "failed to encode password: " + err.Error())
	}
	return hash, nil
}

// проверка введённого юзером пароля на совпадение с хэшем из БД
func PasswordIsCorrect(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
