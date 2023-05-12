package models

type RoomNumber string

type Room struct {
	Number RoomNumber `json:"number"`
	Status string     `json:"status"`
}

type CreateRoomRequest struct {
	Number RoomNumber `json:"number"`
	Status string     `json:"status"`
}
