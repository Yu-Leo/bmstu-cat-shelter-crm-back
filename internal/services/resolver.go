package services

import "github.com/sirupsen/logrus"

type Resolver struct {
	CatService      *CatService
	GuardianService *GuardianService
	ResidentService *ResidentService
	RoomService     *RoomService
	Logger          *logrus.Logger
}

func NewResolver(catService *CatService,
	guardianService *GuardianService,
	residentService *ResidentService, roomService *RoomService, logger *logrus.Logger) *Resolver {

	return &Resolver{
		CatService:      catService,
		GuardianService: guardianService,
		ResidentService: residentService,
		RoomService:     roomService,
		Logger:          logger,
	}
}
