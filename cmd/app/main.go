package main

import (
	"log"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/logger"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/app"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	l := logger.NewLogger(cfg.Logger.Level)

	app.Run(cfg, l)
}
