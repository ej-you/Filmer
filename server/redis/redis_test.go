package redis

import (
	"time"
	"fmt"

	"testing"
)


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"

    testToken = "knevsg5w4o7.gvie45g7nw3ci8.gw3c6g4i5v3kuwfv"
)


var startTime time.Time

func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}

func SuccessLog(t *testing.T, format string, a ...any) {
	t.Logf("\t%s\t%s", successMarker, fmt.Sprintf(format, a...))
}

func ErrorLog(t *testing.T, err error) {
	t.Logf("\t%s\tFailed: %v", failedMarker, err)
}


// connection.go
func TestGetRedisClient(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get redis connection")
	{
		redisConn := GetRedisClient()
		// если подключение не получится, то случится паника
		SuccessLog(t, "Successfully got redis connection: %v", redisConn)
	}
	logExecTime(t, &startTime)
}

// db_funcs.go
func TestSetBlacklistedToken(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set token to blacklist in Redis")
	{
		err := SetBlacklistedToken(testToken)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set token %q to blacklist", testToken)
		}
	}
	logExecTime(t, &startTime)
}

// db_funcs.go
func TestGetBlacklistedToken(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get token from blacklist from Redis")
	{
		isBlacklisted, err := GetBlacklistedToken(testToken)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully. Token %q is blacklisted: %v", testToken, isBlacklisted)
		}
	}
	logExecTime(t, &startTime)
}
