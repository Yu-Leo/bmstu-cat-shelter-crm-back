package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/endpoints/rest"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories/mock"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/logger"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/config"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/httpserver"
)

func Run(cfg *config.Config, l logger.Interface) {
	l.Info("Run application")

	if cfg.Server.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.Default()
	addRouter(ginEngine, l)

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

func addRouter(ginEngine *gin.Engine, l logger.Interface) {
	catRepository := mock.NewMockCatRepository()
	catService := services.NewCatService(catRepository)
	rest.NewRouter(ginEngine, l, catService)
}
