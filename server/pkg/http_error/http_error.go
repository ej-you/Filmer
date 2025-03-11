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
	statusCode int
	message    string
	causeErr   error
}

// HTTPError constructor
func NewHTTPError(statusCode int, message string, causeErr error) HTTPError {
	return &HTTPErr{
		statusCode: statusCode,
		message:    message,
		causeErr:   causeErr,
	}
}

func (he HTTPErr) StatusCode() int {
	return he.statusCode
}

func (he HTTPErr) UserFriendlyMessage() string {
	return he.message
}

func (he HTTPErr) Error() string {
	return fmt.Sprintf("%s: %v", he.message, he.causeErr)
}
