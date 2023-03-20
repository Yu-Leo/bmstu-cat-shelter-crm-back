package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/logger"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type catRoutes struct {
	catService *services.CatService
	logger     logger.Interface
}

func NewCatRoutes(handler *gin.RouterGroup, catService *services.CatService, logger logger.Interface) {
	uR := &catRoutes{
		catService: catService,
		logger:     logger,
	}

	catHandlerGroup := handler.Group("/cats")
	{
		catHandlerGroup.POST("/", uR.CreateCat)
		catHandlerGroup.GET("/", uR.GetCatsList)
	}
}

func (r *catRoutes) CreateCat(c *gin.Context) {
	requestData := models.CreateCatRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newCatId, err := r.catService.CreateCat(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, *newCatId)
}

func (r *catRoutes) GetCatsList(c *gin.Context) {
	catsList, err := r.catService.GetCatsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *catsList)
}
