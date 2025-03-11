package utils

import (
	"fmt"
	"time"
)

// convert minutes to hours:minutes
func RawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/int(time.Hour.Minutes()), minutes%int(time.Hour.Minutes()))
}
