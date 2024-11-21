package logging

import (
	"log"
	"time"
)

type LogLevel int

const (
	LogLevelFatal LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
	LogLevelVerbose
)

type Log struct {
	logger       log.Logger
	logSeparator rune
	logLevel     LogLevel
}

func (l *Log) getLevelString(level LogLevel) string {
	switch level {
	case LogLevelFatal:
		{
			return "FATAL"
		}
	case LogLevelError:
		{
			return "ERROR"
		}
	case LogLevelInfo:
		{
			return "INFO"
		}
	case LogLevelDebug:
		{
			return "DEBUG"
		}
	case LogLevelVerbose:
		{
			return "VERBOSE"
		}
	default:
		{
			return "UNKNOWN"
		}
	}

}

func (l *Log) internalWriter(message string, messageLevel LogLevel) {
	if messageLevel > l.logLevel {
		return
	}
	timeStamp := time.Now().UTC().Format(time.RFC3339Nano)
	levelString := l.getLevelString(messageLevel)
	l.logger.Printf("%s %c %s %c %s\n", timeStamp, l.logSeparator, levelString, l.logSeparator, message)
}

func (l *Log) Fatal(message string) {
	l.internalWriter(message, LogLevelFatal)
}

func (l *Log) Error(message string) {
	l.internalWriter(message, LogLevelError)
}

func (l *Log) Info(message string) {
	l.internalWriter(message, LogLevelInfo)
}

func (l *Log) Debug(message string) {
	l.internalWriter(message, LogLevelDebug)
}

func (l *Log) Verbose(message string) {
	l.internalWriter(message, LogLevelVerbose)
}
