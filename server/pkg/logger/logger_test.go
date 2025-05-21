package logger

import (
	"testing"

	"Filmer/server/config"
)

func TestLogger(t *testing.T) {
	t.Log("Try to init Logger")

	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}

	fMessage := "message"
	logger := NewLogger(cfg)
	t.Logf("logger type: %T", logger)

	logger.Debug("debug sample")
	logger.Debugf("debugf sample: %v", fMessage)
	logger.Info("info sample")
	logger.Infof("infof sample: %v", fMessage)
	logger.Error("error sample")
	logger.Errorf("errorf sample: %v", fMessage)
}
