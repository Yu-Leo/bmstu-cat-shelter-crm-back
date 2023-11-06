package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories/mocks"
)

func TestCreateRoom(t *testing.T) {
	roomRequest := models.CreateRoomRequest{Number: models.RoomNumber("A1"), Status: "Status"}

	roomRepo := mocks.NewRoomRepository(t)
	roomRepo.On("Create", context.Background(), roomRequest).Return(models.RoomNumber("A1"), nil)

	roomService := NewRoomService(roomRepo)

	roomNumber, err := roomService.CreateRoom(roomRequest)

	assert.Equal(t, roomRequest.Number, roomNumber)
	assert.NoError(t, err)
}
