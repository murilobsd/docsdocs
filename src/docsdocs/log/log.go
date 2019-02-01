package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Settings sets the log settings
func Settings(format, out, levelName string) {
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
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}

// DocsLogger struct
type DocsLogger struct {
	*logrus.Logger
}

// NewDocsLogger create new logger
func NewDocsLogger() Logger {
	logger := &DocsLogger{Logger: log}

	return logger

}

// Debug logger
func (d *DocsLogger) Debug(args ...interface{}) {
	d.Logger.Debug(args...)
}

// Info logger
func (d *DocsLogger) Info(args ...interface{}) {
	d.Logger.Info(args...)
}

// Warn logger
func (d *DocsLogger) Warn(args ...interface{}) {
	d.Logger.Warn(args...)
}

// Error logger
func (d *DocsLogger) Error(args ...interface{}) {
	d.Logger.Error(args...)
}
