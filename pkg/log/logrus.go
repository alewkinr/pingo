package log

import "github.com/sirupsen/logrus"

// SetUpLogging — настраиваем логгирование
func SetUpLogging() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	return log
}
