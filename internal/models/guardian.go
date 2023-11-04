package models

type GuardianId int

type Guardian struct {
	Id GuardianId `json:"id"`
	Person
}

type CreateGuardianRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	PhotoURL   string `json:"photoURL"`
	Phone      string `json:"phone"`
}
