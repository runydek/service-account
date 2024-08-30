package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func LogWarning(message string, err error) {
	log.Warnf("%s: %v", message, err)
}

func LogError(message string, err error) {
	log.Errorf("%s: %v", message, err)
}

func LogInfo(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}
