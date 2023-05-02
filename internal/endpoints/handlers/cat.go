package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type catRoutes struct {
	catService *services.CatService
	logger     *logrus.Logger
}

func NewCatRoutes(handler *gin.RouterGroup, catService *services.CatService, logger *logrus.Logger) {
	router := &catRoutes{
		catService: catService,
		logger:     logger,
	}

	handlerGroup := handler.Group("/cats")
	{
		handlerGroup.POST("", router.CreateCat)
		handlerGroup.GET("", router.GetCatsList)
		handlerGroup.GET("/:chip_number", router.GetCat)
		handlerGroup.PUT("/:chip_number", router.UpdateCat)
		handlerGroup.DELETE("/:chip_number", router.DeleteCat)
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
		if err == apperror.CatChipNumberAlreadyExists {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
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
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
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
	if err == apperror.CatNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
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
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateCat
// @Summary     Update cat
// @ID          updateCat
// @Tags  	    cats
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Param createCatObject body models.CreateCatRequest true "Parameters for updating a cat."
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /cats/{chip_number} [put]
func (r *catRoutes) UpdateCat(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")
	requestData := models.CreateCatRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}
	err = r.catService.UpdateCat(models.CatChipNumber{ChipNumber: chipNumber}, requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
