package services

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
)

type RoomService struct {
	repository repositories.RoomRepository
}

func NewRoomService(roomRepository repositories.RoomRepository) *RoomService {
	return &RoomService{
		repository: roomRepository,
	}
}

func (s RoomService) CreateRoom(requestData models.CreateRoomRequest) (models.RoomNumber, error) {
	return s.repository.Create(context.Background(), requestData)
}

func (s RoomService) GetRoomsList() (*[]models.Room, error) {
	return s.repository.GetList(context.Background())
}

func (s RoomService) GetRoom(roomNumber models.RoomNumber) (*models.Room, error) {
	return s.repository.Get(context.Background(), roomNumber)
}

func (s RoomService) DeleteRoom(roomNumber models.RoomNumber) error {
	return s.repository.Delete(context.Background(), roomNumber)
}

func (s RoomService) UpdateRoom(roomNumber models.RoomNumber, requestData models.CreateRoomRequest) error {
	return s.repository.Update(context.Background(), roomNumber, requestData)
}
