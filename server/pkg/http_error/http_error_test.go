package http_errors

import (
	"testing"
)

func TestHTTPErrorInterface(t *testing.T) {
	t.Log("Try to init HTTPError")

	httpErr := NewHTTPError(404, "entity not found")
	t.Logf("HTTPError type: %T", httpErr)

	t.Logf("HTTPError.StatusCode(): %d", httpErr.StatusCode())
	t.Logf("HTTPError.Error(): %s", httpErr.Error())
}
