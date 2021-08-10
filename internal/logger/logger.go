package logger

import (
	"fmt"

	"github.com/aveplen/REST/internal/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(cfg config.Logrus) (*logrus.Logger, error) {
	res := logrus.New()
	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("ConfigureLogger function: %w", err)
	}
	res.SetLevel(logLevel)
	res.SetFormatter(&logrus.JSONFormatter{})
	return res, nil
}
