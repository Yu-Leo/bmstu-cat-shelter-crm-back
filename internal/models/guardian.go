package models

type Guardian struct {
	GuardianId int `json:"id"`
	Person
}

type GuardianId struct {
	Id int `json:"id"`
}

type CreateGuardianRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	PhotoUrl   string `json:"photoUrl"`
	Phone      string `json:"phone"`
}
