package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	hoursInDay = 24
	secInMin   = 60
	secInHour  = 3600
)

// Return days, hours, minutes and seconds until token is expire
func GetJWTExpirationData(accessToken string) (days, hours, minutes, seconds int, err error) {
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
	hours = allSeconds / secInHour // non-normalized
	minutes = (allSeconds - hours*secInHour) / secInMin
	seconds = allSeconds - hours*secInHour - minutes*secInMin
	days = hours / hoursInDay
	hours -= days * hoursInDay
	return
}
