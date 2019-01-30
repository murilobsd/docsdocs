package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// LogTo sets the log settings
func LogTo(format, out, levelName string) {
	// log format
	switch format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	switch out {
	case "file":
		file, err := os.OpenFile("/tmp/docsdocs.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	case "stdout":
		log.Out = os.Stdout
	default:
		log.Out = os.Stdout
	}

	switch levelName {
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "warning":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	default:
		log.Level = logrus.DebugLevel
	}
}

// Logger interface
type Logger interface {
	Debug(...interface{})
	// Info(string, ...interface{})
	// Warn(string, ...interface{}) error
	// Error(string, ...interface{}) error
}

// DocsLogger struct
type DocsLogger struct {
	*logrus.Logger
}

// NewDocsLogger create new logger
func NewDocsLogger() Logger {
	return &DocsLogger{log}
}

// Debug logger
func (d *DocsLogger) Debug(args ...interface{}) {
	d.Logger.Debug(args...)
}
