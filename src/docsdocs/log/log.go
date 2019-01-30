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
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	}

	switch out {
	case "file":
		file, err := os.OpenFile("/tmp/docsdocs.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			log.Output = file
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	case "stdout":
		log.Output = os.Stdout
	default:
		log.Output = os.Stdout
	}

	switch levelName {
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "warning":
		log.Level = logrus.WarningLevel
	case "error":
		log.Level = logrus.ErrorLevel
	default:
		log.Level = logrus.DebugLevel
	}
}
