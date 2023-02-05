package config

import "github.com/sirupsen/logrus"

func InitLogrus() {
	env := GetEnvVariable("ENV")
	if env != "production" {
		logrus.SetLevel(logrus.DebugLevel)
		return
	}
	logrus.SetLevel(logrus.WarnLevel)

}
