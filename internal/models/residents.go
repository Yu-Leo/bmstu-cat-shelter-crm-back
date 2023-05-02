package models

import "time"

type Resident struct {
	Cat
	Booking        bool   `json:"booking"`
	Aggressiveness bool   `json:"aggressiveness"`
	VKAlbumUrl     string `json:"VKAlbumUrl"`
	GuardianId     int    `json:"guardianId"`
}

type CreateResidentRequest struct {
	Nickname                 string    `json:"nickname" binding:"required"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   string    `json:"gender"`
	Age                      int       `json:"age"`
	ChipNumber               string    `json:"chipNumber"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
	Booking                  bool      `json:"booking"`
	Aggressiveness           bool      `json:"aggressiveness"`
	VKAlbumUrl               string    `json:"VKAlbumUrl"`
	GuardianId               int       `json:"guardianId"`
}
