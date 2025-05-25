// Package httperror provides interface for HTTP errors with
// custom error code and user-friendly message.
package httperror

import (
	"fmt"
)

var _ HTTPError = (*httpError)(nil)

type HTTPError interface {
	Error() string
	UserFriendlyMessage() string
	StatusCode() int
}

// HTTPError implementation.
type httpError struct {
	statusCode int
	message    string
	causeErr   error
}

func New(statusCode int, message string, causeErr error) HTTPError {
	return &httpError{
		statusCode: statusCode,
		message:    message,
		causeErr:   causeErr,
	}
}

func (he httpError) StatusCode() int {
	return he.statusCode
}

func (he httpError) UserFriendlyMessage() string {
	return he.message
}

func (he httpError) Error() string {
	return fmt.Sprintf("%s: %v", he.message, he.causeErr)
}
