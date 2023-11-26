package logger

import (
	"github.com/sirupsen/logrus"
)

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func LogError(message, point string) {
	logrus.WithFields(logrus.Fields{
		"point":   point,
		"message": message,
	})
}
