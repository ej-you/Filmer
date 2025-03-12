package utils

import (
	"time"

	"testing"
)

func TestToNextDayDuration(t *testing.T) {
	t.Log("Get duration from now to the next day")

	now := time.Now().UTC()

	nextDay := ToNextDayDuration(now)

	t.Logf("now time: %v", now)
	t.Logf("duration: %v", nextDay)
}
