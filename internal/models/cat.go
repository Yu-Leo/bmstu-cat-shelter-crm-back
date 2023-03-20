package models

import "time"

type Cat struct {
	Id                       int       `json:"id"`
	Nickname                 string    `json:"nickname"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   bool      `json:"gender"`
	Age                      int       `json:"age"`
	ChipNumber               string    `json:"chipNumber"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}

type CatId struct {
	Id int `json:"catId"`
}

type CreateCatRequest struct {
	Nickname                 string    `json:"nickname" binding:"required"`
	PhotoUrl                 string    `json:"photoUrl"`
	Gender                   bool      `json:"gender"`
	Age                      int       `json:"age"`
	ChipNumber               string    `json:"chipNumber"`
	DateOfAdmissionToShelter time.Time `json:"dateOfAdmissionToShelter"`
}
