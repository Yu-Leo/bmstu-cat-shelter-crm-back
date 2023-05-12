package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/errors"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/services"
)

type roomRoutes struct {
	roomService *services.RoomService
	logger      *logrus.Logger
}

func NewRoomRoutes(handler *gin.RouterGroup, roomService *services.RoomService, logger *logrus.Logger) {
	router := &roomRoutes{
		roomService: roomService,
		logger:      logger,
	}

	handlerGroup := handler.Group("/rooms")
	{
		handlerGroup.POST("", router.CreateRoom)
		handlerGroup.GET("", router.GetRoomsList)
		handlerGroup.GET("/:number", router.GetRoom)
		handlerGroup.PUT("/:number", router.UpdateRoom)
		handlerGroup.DELETE("/:number", router.DeleteRoom)
	}
}

// CreateRoom
// @Summary     Create room
// @ID          createRoom
// @Tags  	    rooms
// @Accept      json
// @Produce     json
// @Param createRoomObject body models.CreateRoomRequest true "Parameters for creating a room."
// @Success     201 {object} models.RoomNumber
// @Failure	    400 {object} errors.ErrorJSON
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /rooms [post]
func (r *roomRoutes) CreateRoom(c *gin.Context) {
	requestData := models.CreateRoomRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{
			Message:          errors.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newRoomChipNumber, err := r.roomService.CreateRoom(requestData)
	if err != nil {
		if err == errors.RoomNumberAlreadyExists {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, newRoomChipNumber)
}

// GetRoomsList
// @Summary     Get rooms list
// @ID          getRoomsList
// @Tags  	    rooms
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Room
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /rooms [get]
func (r *roomRoutes) GetRoomsList(c *gin.Context) {
	roomsList, err := r.roomService.GetRoomsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *roomsList)
}

// GetRoom
// @Summary     Get room
// @ID          getRoom
// @Tags  	    rooms
// @Accept      json
// @Produce     json
// @Param number path string true "Number"
// @Success     200 {object} models.Room
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /rooms/{number} [get]
func (r *roomRoutes) GetRoom(c *gin.Context) {
	number := c.Params.ByName("number")
	room, err := r.roomService.GetRoom(models.RoomNumber(number))
	if err == errors.RoomNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, *room)
}

// DeleteRoom
// @Summary     Delete room
// @ID          deleteRoom
// @Tags  	    rooms
// @Accept      json
// @Produce     json
// @Param number path string true "Number"
// @Success     204
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /rooms/{number} [delete]
func (r *roomRoutes) DeleteRoom(c *gin.Context) {
	number := c.Params.ByName("number")
	err := r.roomService.DeleteRoom(models.RoomNumber(number))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateRoom
// @Summary     Update room
// @ID          updateRoom
// @Tags  	    rooms
// @Accept      json
// @Produce     json
// @Param number path string true "Number"
// @Param createRoomObject body models.CreateRoomRequest true "Parameters for updating a room."
// @Success     204
// @Failure	    500 {object} errors.ErrorJSON
// @Router      /rooms/{number} [put]
func (r *roomRoutes) UpdateRoom(c *gin.Context) {
	number := c.Params.ByName("number")
	requestData := models.CreateRoomRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorJSON{
			Message:          errors.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}
	err = r.roomService.UpdateRoom(models.RoomNumber(number), requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
