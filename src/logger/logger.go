package logger

import (
	"os"

	"github.com/one-piece-team1/one-piece-recovery/src/config"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.GetLogLevel())
	if err != nil {
		level = logrus.DebugLevel
	}
	Logger = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProd() {
		Logger.Formatter = &logrus.JSONFormatter{}
	} else {
		Logger.Formatter = &logrus.TextFormatter{}
	}
}
