package models

type Cat struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CatId struct {
	Id int `json:"catId"`
}

type CreateCatRequest struct {
	Name string `json:"name" binding:"required"`
}
