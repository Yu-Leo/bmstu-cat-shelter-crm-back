package app

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/endpoints"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/httpserver"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"
)

func Run(cfg *config.Config, l *logrus.Logger) {
	storage, err := sqlitedb.NewStorage(cfg.Storage.Path)
	if err != nil {
		l.Errorf("SQLite database open error: %e", err)
		return
	}
	if err != nil {
		l.Errorf("SQLite database init error: %e", err)
		return
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err == nil {
			l.Info("SQLite database close")
		} else {
			l.Errorf("SQLite database close error: %e", err)
		}
	}(storage.DB)

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
	guardianRepository := repositories.NewSqliteGuardianRepository(storage)
	residentRepository := repositories.NewSqliteResidentRepository(storage)
	roomRepository := repositories.NewSqliteRoomRepository(storage)

	resolver := services.NewResolver(
		services.NewGuardianService(guardianRepository),
		services.NewResidentService(residentRepository),
		services.NewRoomService(roomRepository),
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
