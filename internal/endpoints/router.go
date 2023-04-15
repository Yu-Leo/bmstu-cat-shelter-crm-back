package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/endpoints/handlers"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/logger"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"

	_ "github.com/Yu-Leo/bmstu-cat-shelter-crm-back/docs"
)

// @title           Cat Shelter CRM
// @version         1.0

// @contact.name   Lev Yuvenskiy
// @contact.email  levayu22@gmail.com

// @host      127.0.0.1:9000
// @BasePath  /

func NewRouter(ginEngine *gin.Engine, logger logger.Interface,
	catService *services.CatService) {

	// Routers
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginEngine.GET("/health", health)
	router := ginEngine.Group("")
	{
		handlers.NewCatRoutes(router, catService, logger)
	}
}

type healthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Health
// @Summary     Health check
// @ID          health
// @Tags  	    health
// @Accept      json
// @Produce     json
// @Success     200 {object} healthResponse
// @Router      /health [get]
func health(c *gin.Context) {
	var a = healthResponse{Status: 200,
		Message: "OK"}
	c.JSON(http.StatusOK, a)
}
