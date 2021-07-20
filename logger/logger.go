package logger

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once
var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
	})
	return logger
}
