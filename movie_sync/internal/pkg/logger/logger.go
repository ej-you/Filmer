// Package logger provides Logger interface to log with different levels.
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Init sets up main logger for application.
func Init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
