package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/app"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		logrus.Fatal("Config error: %s", err)
	}

	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logrus.Error("Invalid log type: %s", err)
	}
	logger.SetLevel(level)

	app.Run(cfg, logger)
}
