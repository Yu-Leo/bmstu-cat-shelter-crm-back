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

	app.Run(cfg, logger)
}
