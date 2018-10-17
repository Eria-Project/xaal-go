package log

import (
	logger "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

func init() {
	logger.SetFormatter(&logger.TextFormatter{FullTimestamp: true, TimestampFormat: "01-02|15:04:05"})
}

// Debug is a convenient alias for logger.Debug
func Debug(msg string, ctx Fields) {
	logger.Debug(msg)
}

// Info is a convenient alias for logger.Info
func Info(msg string, ctx Fields) {
	logger.WithFields(logger.Fields(ctx)).Info(msg)
}

// Warn is a convenient alias for logger.Warn
func Warn(msg string, ctx Fields) {
	logger.Warn(msg)
}

// Error is a convenient alias for logger.Error
func Error(msg string, ctx Fields) {
	logger.Error(msg)
}

// Fatal is a convenient alias for logger.Crit
func Fatal(msg string, ctx Fields) {
	logger.Fatal(msg)
}
