package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

var Log *logrus.Logger

func SetupLogger(env string) error {
	Log = logrus.New()
	switch env {
	case envLocal:
		Log.SetLevel(logrus.DebugLevel)
		Log.SetOutput(os.Stdout)
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	case envProd:
		Log.SetLevel(logrus.InfoLevel)
		Log.SetFormatter(&logrus.JSONFormatter{})
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Log.SetOutput(file)
		} else {
			return fmt.Errorf("Failed to open log file: %w", err)
		}
	default:
		return fmt.Errorf("Unknown env: %s", env)
	}
	Log.WithFields(logrus.Fields{"env": env})
	return nil
}
