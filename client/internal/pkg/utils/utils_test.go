package utils

import (
	"testing"
)

var authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI4MjE1NDQsInN1YiI6IjViNjk3OGVkLTZjYjQtNDk2Zi04ZGIzLTYzY2RlZDc5YTg1YyJ9.wf7S8UvU5mBenzdHIRwqwxjs0zs7ihC2qaUkxwPM1DY"

func TestGetJWTExpirationData(t *testing.T) {
	t.Log("Test getting JWT expiration data")

	d, h, m, s, err := GetJWTExpirationData(authToken)
	if err != nil {
		t.Fatalf("Failed to got JWT expiration data: %v", err)
	}
	t.Logf("Successfully got JWT expiration data: %dd %dh %dm %ds", d, h, m, s)
}
