package models

type Person struct {
	PersonId   int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	PhotoURL   string `json:"photoURL"`
	Phone      string `json:"phone"`
}
