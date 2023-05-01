package models

type Person struct {
	PersonId   int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	PhotoUrl   string `json:"photoUrl"`
	Phone      string `json:"phone"`
}
