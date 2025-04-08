package helpers

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	//use farmework logrus
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("Logger initialized Using logrus")
	return log
}
