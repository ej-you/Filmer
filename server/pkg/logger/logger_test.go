package logger

import (
	"testing"
)


func TestLogger(t *testing.T) {
	t.Log("Try to init Logger")

	fMessage := "message"
	logger := NewLogger()
	t.Logf("logger type: %T", logger)

	logger.Debug("debug sample")
	logger.Debugf("debugf sample: %v", fMessage)
	logger.Info("info sample")
	logger.Infof("infof sample: %v", fMessage)
	logger.Error("error sample")
	logger.Errorf("errorf sample: %v", fMessage)
}
