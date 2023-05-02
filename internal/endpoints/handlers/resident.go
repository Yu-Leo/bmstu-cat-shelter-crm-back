package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type residentRoutes struct {
	residentService *services.ResidentService
	logger          *logrus.Logger
}

func NewResidentRoutes(handler *gin.RouterGroup, residentService *services.ResidentService, logger *logrus.Logger) {
	router := &residentRoutes{
		residentService: residentService,
		logger:          logger,
	}

	handlerGroup := handler.Group("/residents")
	{
		handlerGroup.POST("", router.CreateResident)
		handlerGroup.GET("", router.GetResidentsList)
		handlerGroup.GET("/:chip_number", router.GetResident)
		handlerGroup.PUT("/:chip_number", router.UpdateResident)
		handlerGroup.DELETE("/:chip_number", router.DeleteResident)
	}
}

// CreateResident
// @Summary     Create resident
// @ID          createResident
// @Tags  	    residents
// @Accept      json
// @Produce     json
// @Param createResidentObject body models.CreateResidentRequest true "Parameters for creating a resident."
// @Success     201 {object} models.CatChipNumber
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /residents [post]
func (r *residentRoutes) CreateResident(c *gin.Context) {
	requestData := models.CreateResidentRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newResidentId, err := r.residentService.CreateResident(requestData)
	if err != nil {
		if err == apperror.PersonPhoneAlreadyExists {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, newResidentId)
}

// GetResidentsList
// @Summary     Get residents list
// @ID          getResidentsList
// @Tags  	    residents
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Resident
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /residents [get]
func (r *residentRoutes) GetResidentsList(c *gin.Context) {
	residentsList, err := r.residentService.GetResidentsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *residentsList)
}

// GetResident
// @Summary     Get resident
// @ID          getResident
// @Tags  	    residents
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Success     200 {object} models.Resident
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /residents/{chip_number} [get]
func (r *residentRoutes) GetResident(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")
	resident, err := r.residentService.GetResident(models.CatChipNumber(chipNumber))
	if err == apperror.ResidentNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *resident)
}

// DeleteResident
// @Summary     Delete resident
// @ID          deleteResident
// @Tags  	    residents
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /residents/{chip_number} [delete]
func (r *residentRoutes) DeleteResident(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")

	err := r.residentService.DeleteResident(models.CatChipNumber(chipNumber))
	if err == apperror.ResidentNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateResident
// @Summary     Update resident
// @ID          updateResident
// @Tags  	    residents
// @Accept      json
// @Produce     json
// @Param chip_number path string true "Chip number"
// @Param createResidentObject body models.CreateResidentRequest true "Parameters for updating a resident."
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /residents/{chip_number} [put]
func (r *residentRoutes) UpdateResident(c *gin.Context) {
	chipNumber := c.Params.ByName("chip_number")
	requestData := models.CreateResidentRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	err = r.residentService.UpdateResident(models.CatChipNumber(chipNumber), requestData)
	if err == apperror.ResidentNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
