package models

import "time"

type CatChipNumber string

type Cat struct {
	ChipNumber               CatChipNumber `json:"chipNumber"`
	Nickname                 string        `json:"nickname"`
	PhotoURL                 string        `json:"photoURL"`
	Gender                   string        `json:"gender"`
	Age                      int           `json:"age"`
	DateOfAdmissionToShelter time.Time     `json:"dateOfAdmissionToShelter"`
}

type CreateCatRequest struct {
	ChipNumber               CatChipNumber `json:"chipNumber"`
	Nickname                 string        `json:"nickname" binding:"required"`
	PhotoURL                 string        `json:"photoURL"`
	Gender                   string        `json:"gender"`
	Age                      int           `json:"age"`
	DateOfAdmissionToShelter time.Time     `json:"dateOfAdmissionToShelter"`
}
