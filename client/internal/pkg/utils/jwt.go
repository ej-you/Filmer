package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Return days, hours, minutes and seconds until token is expire
func GetJWTExpirationData(accessToken string) (int, int, int, int, error) {
	token, _, err := jwt.NewParser().ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parse token: %w", err)
	}
	// get time.Time instance of expiration time
	expirationTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("parse token expiration time: %w", err)
	}
	// get time duration to expiration time
	expirationDuration := expirationTime.Time.Sub(time.Now().UTC())
	// if token already expired
	if expirationDuration < 0 {
		return 0, 0, 0, 0, nil
	}

	allSeconds := int(expirationDuration.Seconds())
	// calc days, hours, minutes and seconds to expiration time separately
	h := allSeconds / 3600 // non-normalized
	m := (allSeconds - h*3600) / 60
	s := allSeconds - h*3600 - m*60
	d := h / 24
	h = h - d*24

	return d, h, m, s, nil
}
