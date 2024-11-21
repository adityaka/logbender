package logging

import (
	"log"
	"os"
)

var loggers map[string]*Log = make(map[string]*Log)

type LoggingConfig struct {
	Name         string
	FileFullPath string
}

func GetLogger(loggerConfig LoggingConfig) *Log {
	if loggers[loggerConfig.Name] == nil {
		filePath := loggerConfig.FileFullPath
		logWriter, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return nil
		}
		l := Log{logger: log.Logger{}, logSeparator: '|', logLevel: LogLevelDebug}
		l.logger.SetOutput(logWriter)
		loggers[loggerConfig.Name] = &l
	}
	return loggers[loggerConfig.Name]
}
