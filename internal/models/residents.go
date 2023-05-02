package models

type Resident struct {
	Cat
	Booking        bool   `json:"booking"`
	Aggressiveness bool   `json:"aggressiveness"`
	VKAlbumUrl     string `json:"VKAlbumUrl"`
	GuardianId     int    `json:"guardianId"`
}

type CreateResidentRequest struct {
	CreateCatRequest
	Booking        bool   `json:"booking"`
	Aggressiveness bool   `json:"aggressiveness"`
	VKAlbumUrl     string `json:"VKAlbumUrl"`
	GuardianId     int    `json:"guardianId"`
}
