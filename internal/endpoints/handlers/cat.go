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
		catHandlerGroup.POST("", uR.CreateCat)
		catHandlerGroup.GET("", uR.GetCatsList)
		catHandlerGroup.GET("/:chip_number", uR.GetCat)
		catHandlerGroup.DELETE("/:chip_number", uR.DeleteCat)

	}
}

// CreateCat
// @Summary     Create cat
// @ID          createCat
// @Tags  	    cats
// @Accept      json
// @Produce     json
// @Param createCatObject body models.CreateCatRequest true "Parameters for creating a cat."
// @Success     201 {object} models.CatChipNumber
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /cats [post]
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
		if err == apperror.ValidationError {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, *newCatId)
}

// GetCatsList
// @Summary     Get cats list
// @ID          getCatsList
// @Tags  	    cats
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Cat
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /cats [get]
func (r *catRoutes) GetCatsList(c *gin.Context) {
	catsList, err := r.catService.GetCatsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *catsList)
}

// GetCat
// @Summary     Get cat
// @ID          getCat
// @Tags  	    cats
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Success     200 {object} models.Cat
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /cats/{chip_number} [get]
func (r *catRoutes) GetCat(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")
	cat, err := r.catService.GetCat(models.CatChipNumber{ChipNumber: chipNumber})
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	if cat == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, *cat)
}

// DeleteCat
// @Summary     Delete cat
// @ID          deleteCat
// @Tags  	    cats
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /cats/{chip_number} [delete]
func (r *catRoutes) DeleteCat(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")
	err := r.catService.DeleteCat(models.CatChipNumber{ChipNumber: chipNumber})
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
