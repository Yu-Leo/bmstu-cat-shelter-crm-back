package app

import (
	"context"
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
		l.Errorf("SQLite database open error: %e", err)
		return
	}
	err = storage.Init(context.Background())
	if err != nil {
		l.Errorf("SQLite database init error: %e", err)
		return
	}
	defer storage.DB.Close()

	l.Info("Start application")

	if cfg.Server.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	ginEngine := gin.Default()
	addRouters(ginEngine, l, storage)

	httpServer := httpserver.New(ginEngine, cfg.Server.Host, cfg.Server.Port)

	l.Infof("Run server on http://%s:%d", cfg.Server.Host, cfg.Server.Port)
	l.Infof("Open Swagger UI on http://%s:%d/swagger/index.html", cfg.Server.Host, cfg.Server.Port)

	trackSignals(httpServer, l)
}

func addRouters(ginEngine *gin.Engine, logger *logrus.Logger, storage *sqlitedb.Storage) {
	catRepository := repositories.NewSqliteCatRepository(storage)
	guardianRepository := repositories.NewSqliteGuardianRepository(storage)
	residentRepository := repositories.NewSqliteResidentRepository(storage)

	resolver := services.NewResolver(
		services.NewCatService(catRepository),
		services.NewGuardianService(guardianRepository),
		services.NewResidentService(residentRepository),
		logger)

	endpoints.NewRouter(ginEngine, resolver)
}

func trackSignals(server *httpserver.Server, l *logrus.Logger) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Infof("Catch the %s signal", s.String())
	case err := <-server.Notify():
		l.Errorf("HTTPServer notify error: %e", err)
	}

	err := server.Shutdown()
	if err == nil {
		l.Infof("Shutdown HTTPServer")
	} else {
		l.Errorf("HTTPServer shutdown error: %e", err)
	}
}
