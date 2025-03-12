package utils

import (
	"fmt"
	"time"
)

// Convert minutes to hours:minutes
func RawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/int(time.Hour.Minutes()), minutes%int(time.Hour.Minutes()))
}

// Get duration from the given time to time of the beginning of the next day
func ToNextDayDuration(from time.Time) time.Duration {
	// parse year, month and day of given time
	nowYear, nowMonth, nowDay := from.Date()
	// get next day time
	nextDayTime := time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.UTC).Add(24 * time.Hour)
	return nextDayTime.Sub(from)
}
