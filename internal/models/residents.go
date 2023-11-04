package models

type Resident struct {
	Cat
	Booking        bool       `json:"booking"`
	Aggressiveness bool       `json:"aggressiveness"`
	VKAlbumURL     string     `json:"VKAlbumURL"`
	GuardianId     GuardianId `json:"guardianId"`
	RoomNumber     RoomNumber `json:"roomNumber"`
}

type CreateResidentRequest struct {
	CreateCatRequest
	Booking        bool       `json:"booking"`
	Aggressiveness bool       `json:"aggressiveness"`
	VKAlbumURL     string     `json:"VKAlbumURL"`
	GuardianId     GuardianId `json:"guardianId"`
	RoomNumber     RoomNumber `json:"roomNumber"`
}
