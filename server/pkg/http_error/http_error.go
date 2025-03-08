package http_errors

import (
	"fmt"
)


// HTTPError interface
type HTTPError interface {
	Error() string
	UserFriendlyMessage() string
	StatusCode() int
}


// HTTPError interface implementation
type HTTPErr struct {
	statusCode	int
	message		string
	causeErr	error
}

// HTTPError constructor
func NewHTTPError(statusCode int, message string, causeErr error) HTTPError {
	return &HTTPErr{
		statusCode: statusCode,
		message: message,
		causeErr: causeErr,
	}
}

func (this HTTPErr) StatusCode() int {
	return this.statusCode
}

func (this HTTPErr) UserFriendlyMessage() string {
	return this.message
}

func (this HTTPErr) Error() string {
	return fmt.Sprintf("%s: %v", this.message, this.causeErr)
}
