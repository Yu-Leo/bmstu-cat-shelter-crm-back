package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type guardianRoutes struct {
	guardianService *services.GuardianService
	logger          *logrus.Logger
}

func NewGuardianRoutes(handler *gin.RouterGroup, guardianService *services.GuardianService, logger *logrus.Logger) {
	uR := &guardianRoutes{
		guardianService: guardianService,
		logger:          logger,
	}

	catHandlerGroup := handler.Group("/guardians")
	{
		catHandlerGroup.POST("", uR.CreateGuardian)
		catHandlerGroup.GET("", uR.GetGuardiansList)
		catHandlerGroup.GET("/:id", uR.GetGuardian)
		catHandlerGroup.PUT("/:id", uR.UpdateGuardian)
		catHandlerGroup.DELETE("/:id", uR.DeleteGuardian)
	}
}

// CreateGuardian
// @Summary     Create guardian
// @ID          createGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param createCatObject body models.CreateGuardianRequest true "Parameters for creating a guardian."
// @Success     201 {object} models.GuardianId
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /guardians [post]
func (r *guardianRoutes) CreateGuardian(c *gin.Context) {
	requestData := models.CreateGuardianRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newGuardianId, err := r.guardianService.CreateGuardian(requestData)
	if err != nil {
		if err == apperror.PersonPhoneAlreadyExists {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, *newGuardianId)
}

// GetGuardiansList
// @Summary     Get guardians list
// @ID          getGuardiansList
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Guardian
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /guardians [get]
func (r *guardianRoutes) GetGuardiansList(c *gin.Context) {
	guardiansList, err := r.guardianService.GetGuardiansList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *guardiansList)
}

// GetGuardian
// @Summary     Get guardian
// @ID          getGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param id path int true "ID"
// @Success     200 {object} models.Guardian
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /guardians/{id} [get]
func (r *guardianRoutes) GetGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: apperror.InvalidGuardianIdMsg})
	}
	cat, err := r.guardianService.GetGuardian(models.GuardianId{Id: id})
	if err == apperror.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *cat)
}

// DeleteGuardian
// @Summary     Delete guardian
// @ID          deleteGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param id path int true "ID"
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /guardians/{id} [delete]
func (r *guardianRoutes) DeleteGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: apperror.InvalidGuardianIdMsg})
	}

	err = r.guardianService.DeleteGuardian(models.GuardianId{Id: id})
	if err == apperror.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateGuardian
// @Summary     Update guardian
// @ID          updateGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param id path int true "ID"
// @Param createGuardianObject body models.CreateGuardianRequest true "Parameters for updating a guardian."
// @Success     204
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /guardians/{id} [put]
func (r *guardianRoutes) UpdateGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: apperror.InvalidGuardianIdMsg})
	}
	requestData := models.CreateGuardianRequest{}

	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	err = r.guardianService.UpdateGuardian(models.GuardianId{Id: id}, requestData)
	if err == apperror.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
