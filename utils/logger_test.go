package utils

import (
	"go.uber.org/zap"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger1 := GetLogger(zap.NewDevelopment)
	logger2 := GetLogger(zap.NewDevelopment)
	if logger1 != logger2 {
		t.Error("Expected both loggers have the same reference")
	}
}
