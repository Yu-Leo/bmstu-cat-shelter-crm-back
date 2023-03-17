package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/endpoints/rest"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories/sqlite"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/logger"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqliteStorage"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/httpserver"
)

func Run(cfg *config.Config, l logger.Interface) {
	storage := sqliteStorage.NewStorage("db.db")
	if err := storage.Init(context.Background()); err != nil {
		fmt.Println(err)
		l.Error(fmt.Sprintf("SQLite database open errr: %e", err))
		return
	}
	defer storage.DB.Close()

	l.Info("Run application")

	if cfg.Server.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.Default()
	addRouter(ginEngine, l, storage)

	httpServer := httpserver.New(ginEngine, cfg.Server.Host, cfg.Server.Port)
	l.Info(fmt.Sprintf("Run server on %s:%d", cfg.Server.Host, cfg.Server.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info(fmt.Sprintf("Catch the %s signal", s.String()))
	case err := <-httpServer.Notify():
		l.Error(fmt.Sprintf("HTTPServer notify error: %e", err))
	}

	err := httpServer.Shutdown()
	if err == nil {
		l.Info("Shutdown HTTPServer")
	} else {
		l.Error(fmt.Sprintf("HTTPServer shutdown error: %e", err))
	}
}

func addRouter(ginEngine *gin.Engine, l logger.Interface, storage *sqliteStorage.Storage) {
	catRepository := sqlite.NewSqliteCatRepository(storage)
	catService := services.NewCatService(catRepository)
	rest.NewRouter(ginEngine, l, catService)
}
