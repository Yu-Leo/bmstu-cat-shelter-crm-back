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
	PhotoUrl   string `json:"photoUrl"`
	Phone      string `json:"phone"`
}
