package models

import "time"

type CatChipNumber string

type Cat struct {
	CatChipNumber            CatChipNumber `json:"chipNumber"`
	Nickname                 string        `json:"nickname"`
	PhotoUrl                 string        `json:"photoUrl"`
	Gender                   string        `json:"gender"`
	Age                      int           `json:"age"`
	DateOfAdmissionToShelter time.Time     `json:"dateOfAdmissionToShelter"`
}

type CreateCatRequest struct {
	CatChipNumber            CatChipNumber `json:"chipNumber"`
	Nickname                 string        `json:"nickname" binding:"required"`
	PhotoUrl                 string        `json:"photoUrl"`
	Gender                   string        `json:"gender"`
	Age                      int           `json:"age"`
	DateOfAdmissionToShelter time.Time     `json:"dateOfAdmissionToShelter"`
}
