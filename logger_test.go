package logger_test

import (
	"github.com/win-d/logger"
	"os"
	"runtime"
	"testing"
)

func TestSetDir(t *testing.T) {
	dir := getTestDir()
	err := logger.SetDir(dir)
	if err != nil {
		t.Error(err.Error())
	}

	err = os.Remove(dir)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestWrite(t *testing.T) {
	dir := getTestDir()
	err := logger.SetDir(dir)
	if err != nil {
		t.Error(err.Error())
	}

	logger.Write("Test string")
	err = os.RemoveAll(dir)
	if err != nil {
		t.Error(err.Error())
	}
}

func getTestDir() string {
	switch runtime.GOOS {
	case "windows":
		return "D:\\TestLogDir"
	default:
		return "/tmp/testLogDir"
	}
}
