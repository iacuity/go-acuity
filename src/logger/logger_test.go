package logger_test

import (
	"testing"

	"logger"
)

func TestInit(t *testing.T) {
	fileName := "testlog.log"

	err := logger.Init(fileName)

	if nil != err {
		t.Fatal("Failed to Initialize Logger")
	}

	logger.SetLogLevel(logger.INFO)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
	logger.Fatal("Fatal")
}
