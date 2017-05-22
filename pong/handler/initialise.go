package handler

import (
	config "github.com/chrusty/micro-pingpong/pong/config"

	logrus "github.com/sirupsen/logrus"
)

var (
	logger        = logrus.WithFields(logrus.Fields{"logger": "handler"})
	serviceConfig config.Config
)

func Initialise(serviceConfigFromMain config.Config) error {

	logger.Debug("Initialising handler")

	// Store this in a package-wide var:
	serviceConfig = serviceConfigFromMain

	return nil
}
