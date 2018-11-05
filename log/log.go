package log

import (
	"io"

	logger "github.com/sirupsen/logrus"
)

// Fields : alias to logger.Fields
type Fields map[string]interface{}

// Disabled : indicates if the log are disabled
var Disabled = false

func init() {
	logger.SetFormatter(&logger.TextFormatter{FullTimestamp: true, TimestampFormat: "02/01|15:04:05"})
}

// Debug is a convenient alias for logger.Debug
func Debug(msg string, ctx Fields) {
	if !Disabled {
		logger.WithFields(logger.Fields(ctx)).Debug(msg)
	}
}

// Info is a convenient alias for logger.Info
func Info(msg string, ctx Fields) {
	if !Disabled {
		logger.WithFields(logger.Fields(ctx)).Info(msg)
	}
}

// Warn is a convenient alias for logger.Warn
func Warn(msg string, ctx Fields) {
	if !Disabled {
		logger.WithFields(logger.Fields(ctx)).Warn(msg)
	}
}

// Error is a convenient alias for logger.Error
func Error(msg string, ctx Fields) {
	if !Disabled {
		logger.WithFields(logger.Fields(ctx)).Error(msg)
	}
}

// Fatal is a convenient alias for logger.Fatal
func Fatal(msg string, ctx Fields) {
	if !Disabled {
		logger.WithFields(logger.Fields(ctx)).Fatal(msg)
	}
}

// SetLevelDebug : alias to logger DebugLevel
func SetLevelDebug() {
	logger.SetLevel(logger.DebugLevel)
}

// SetLevelWarn : alias to logger WarnLevel
func SetLevelWarn() {
	logger.SetLevel(logger.WarnLevel)
}

// SetLevelError : alias to logger ErrorLevel
func SetLevelError() {
	logger.SetLevel(logger.ErrorLevel)
}

// SetOutput : alias to logger SetOutput
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}
