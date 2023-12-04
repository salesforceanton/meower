package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func LogError(message, point string) {
	logrus.WithFields(logrus.Fields{
		"point":   point,
		"message": message,
	}).Error()
}

func LogInfo(message string) {
	logrus.Info(message)
}
