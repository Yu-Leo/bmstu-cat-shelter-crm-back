package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/errors"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type guardianRoutes struct {
	guardianService *services.GuardianService
	logger          *logrus.Logger
}

func NewGuardianRoutes(handler *gin.RouterGroup, guardianService *services.GuardianService, logger *logrus.Logger) {
	router := &guardianRoutes{
		guardianService: guardianService,
		logger:          logger,
	}

	handlerGroup := handler.Group("/guardians")
	{
		handlerGroup.POST("", router.CreateGuardian)
		handlerGroup.GET("", router.GetGuardiansList)
		handlerGroup.GET("/:id", router.GetGuardian)
		handlerGroup.PUT("/:id", router.UpdateGuardian)
		handlerGroup.DELETE("/:id", router.DeleteGuardian)
	}
}

// CreateGuardian
// @Summary     Create guardian
// @ID          createGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param createGuardianObject body models.CreateGuardianRequest true "Parameters for creating a guardian."
// @Success     201 {object} models.GuardianId
// @Failure	    400 {object} errors.ErrorJSON
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /guardians [post]
func (r *guardianRoutes) CreateGuardian(c *gin.Context) {
	requestData := models.CreateGuardianRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{
			Message:          errors.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newGuardianId, err := r.guardianService.CreateGuardian(requestData)
	if err != nil {
		if err == errors.PersonPhoneAlreadyExists {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, newGuardianId)
}

// GetGuardiansList
// @Summary     Get guardians list
// @ID          getGuardiansList
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Guardian
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /guardians [get]
func (r *guardianRoutes) GetGuardiansList(c *gin.Context) {
	guardiansList, err := r.guardianService.GetGuardiansList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
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
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /guardians/{id} [get]
func (r *guardianRoutes) GetGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{Message: errors.InvalidGuardianIdMsg})
	}
	guardian, err := r.guardianService.GetGuardian(models.GuardianId(id))
	if err == errors.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *guardian)
}

// DeleteGuardian
// @Summary     Delete guardian
// @ID          deleteGuardian
// @Tags  	    guardians
// @Accept      json
// @Produce     json
// @Param id path int true "ID"
// @Success     204
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /guardians/{id} [delete]
func (r *guardianRoutes) DeleteGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{Message: errors.InvalidGuardianIdMsg})
	}

	err = r.guardianService.DeleteGuardian(models.GuardianId(id))
	if err == errors.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
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
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /guardians/{id} [put]
func (r *guardianRoutes) UpdateGuardian(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{Message: errors.InvalidGuardianIdMsg})
	}
	requestData := models.CreateGuardianRequest{}

	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{
			Message:          errors.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	err = r.guardianService.UpdateGuardian(models.GuardianId(id), requestData)
	if err == errors.GuardianNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
