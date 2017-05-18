package logger

import (
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

type LogrusLogger struct {
	Logger *logrus.Logger
}

func NewLogrusFileLogger(file *os.File) *LogrusLogger {
	logger := logrus.New()
	logger.Out = file
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04",
	}
	return &LogrusLogger{
		Logger: logger,
	}
}

func NewLogrusStdLogger() *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04",
	}
	return &LogrusLogger{
		Logger: logger,
	}
}

func NewLogrusFileStdLogger(file *os.File) *LogrusLogger {
	logger := logrus.New()
	logger.Out = io.MultiWriter(file, os.Stdout)
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04",
	}
	return &LogrusLogger{
		Logger: logger,
	}
}

func (l *LogrusLogger) Println(args ...interface{}) {
	l.Logger.Println(args)
}
func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}
func (l *LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args)
}
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args)
}
