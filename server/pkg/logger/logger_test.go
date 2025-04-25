package logger

import (
	"testing"

	"Filmer/server/config"
)

func TestLogger(t *testing.T) {
	t.Log("Try to init Logger")

	fMessage := "message"
	logger := NewLogger(config.NewConfig())
	t.Logf("logger type: %T", logger)

	logger.Debug("debug sample")
	logger.Debugf("debugf sample: %v", fMessage)
	logger.Info("info sample")
	logger.Infof("infof sample: %v", fMessage)
	logger.Error("error sample")
	logger.Errorf("errorf sample: %v", fMessage)
}
