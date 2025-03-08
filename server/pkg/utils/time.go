package utils

import "fmt"


// convert minutes to hours:minutes
func RawMinutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%d", minutes/60, minutes%60)
}
