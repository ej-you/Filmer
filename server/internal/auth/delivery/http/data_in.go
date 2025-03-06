package http


//easyjson:json
//	@description	данные для регистрации юзера
type signUpIn struct {
	// почта юзера
	Email 		string `json:"email" validate:"required,max=100" example:"user@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
}

//easyjson:json
// @description данные для входа юзера
type loginIn struct {
	// почта юзера
	Email 		string `json:"email" validate:"required,max=100" example:"user@gmail.com" maxLength:"100"`
	// пароль юзера
	Password 	string `json:"password" validate:"required,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
}
