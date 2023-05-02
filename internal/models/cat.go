package models

import "time"

type CatChipNumber struct {
	ChipNumber string `json:"chipNumber"`
}

type Cat struct {
	CatChipNumber
	Nickname                 string    `json:"nickname"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   string    `json:"gender"`
	Age                      int       `json:"age"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}

type CreateCatRequest struct {
	CatChipNumber
	Nickname                 string    `json:"nickname" binding:"required"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   string    `json:"gender"`
	Age                      int       `json:"age"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}
