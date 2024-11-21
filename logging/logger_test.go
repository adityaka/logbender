package logging

import (
	"log"
	"os"
	"path"
	"testing"
)

func removeIfFileExists(filePath string) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if !stat.Mode().IsRegular() {
		log.Fatalf("Path passed is not a regular file %s", filePath)
	}
	err = os.Remove(filePath)
	if err != nil {
		log.Fatalf("Error deleting file %s", filePath)
	}
}

func createLogFile(filePath string) *os.File {
	logWriter, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return logWriter
}

func TestLog(t *testing.T) {
	l := Log{}
	tempDir := os.TempDir()
	fileFullPath := path.Join(tempDir, "logbender_log_test.log")
	removeIfFileExists(fileFullPath)
	logWriter := createLogFile(fileFullPath)
	l.logger = log.Logger{}
	l.logger.SetOutput(logWriter)
	l.logLevel = LogLevelDebug
	l.logSeparator = '|'
	l.Info("Hello World")
	l.Error("Error in World")
	l.Debug("Debugging in testing world")
	l.Verbose("This won't get logged")
}

func TestLogManager(t *testing.T) {
	logFilePath := path.Join(os.TempDir(), "logmanager_test.log")
	removeIfFileExists(logFilePath)
	createLogFile(logFilePath)
	l := GetLogger(LoggingConfig{Name: "testing_logmanager", FileFullPath: logFilePath})
	if l == nil {
		log.Fatal("Failed to initialize Logger")
	}
	l.Info("Hello World")
	l.Error("Error in World")
	l.Debug("Debugging in testing world")
	l.Verbose("This won't get logged")
}
