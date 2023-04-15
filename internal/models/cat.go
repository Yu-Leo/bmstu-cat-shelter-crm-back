package models

import "time"

type Cat struct {
	Nickname                 string    `json:"nickname"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   string    `json:"gender"`
	Age                      int       `json:"age"`
	ChipNumber               string    `json:"chipNumber"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}

type CatChipNumber struct {
	ChipNumber string `json:"chipNumber"`
}

type CreateCatRequest struct {
	Nickname                 string    `json:"nickname" binding:"required"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   string    `json:"gender"`
	Age                      int       `json:"age"`
	ChipNumber               string    `json:"chipNumber"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}
