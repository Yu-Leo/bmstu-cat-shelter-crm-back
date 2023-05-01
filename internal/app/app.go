package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/endpoints"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/httpserver"
)

func Run(cfg *config.Config, l *logrus.Logger) {
	storage, err := sqlitedb.NewStorage(cfg.Storage.Path)
	if err != nil {
		l.Error(fmt.Sprintf("SQLite database open error: %e", err))
		return
	}
	err = storage.Init(context.Background())
	if err != nil {
		l.Error(fmt.Sprintf("SQLite database init error: %e", err))
		return
	}
	defer storage.DB.Close()

	l.Info("Start application")

	if cfg.Server.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.Default()
	addRouter(ginEngine, l, storage)

	httpServer := httpserver.New(ginEngine, cfg.Server.Host, cfg.Server.Port)
	l.Info(fmt.Sprintf("Run server on http://%s:%d", cfg.Server.Host, cfg.Server.Port))
	l.Info(fmt.Sprintf("Open Swagger UI on http://%s:%d/swagger/index.html", cfg.Server.Host, cfg.Server.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info(fmt.Sprintf("Catch the %s signal", s.String()))
	case err := <-httpServer.Notify():
		l.Error(fmt.Sprintf("HTTPServer notify error: %e", err))
	}

	err = httpServer.Shutdown()
	if err == nil {
		l.Info("Shutdown HTTPServer")
	} else {
		l.Error(fmt.Sprintf("HTTPServer shutdown error: %e", err))
	}
}

func addRouter(ginEngine *gin.Engine, l *logrus.Logger, storage *sqlitedb.Storage) {
	catRepository := repositories.NewSqliteCatRepository(storage)
	guardianRepository := repositories.NewSqliteGuardianRepository(storage)

	catService := services.NewCatService(catRepository)
	guardianService := services.NewGuardianService(guardianRepository)

	endpoints.NewRouter(ginEngine, l, catService, guardianService)
}
