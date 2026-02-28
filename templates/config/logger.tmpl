package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var AppLogger *logrus.Logger

func LoadLogger() {
	AppLogger = logrus.New()
	file, _ := os.OpenFile("go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	multiWriter := io.MultiWriter(os.Stdout, file)
	AppLogger.SetFormatter(&logrus.JSONFormatter{})
	AppLogger.SetOutput(multiWriter)
}
