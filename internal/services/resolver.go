package services

import "github.com/sirupsen/logrus"

type Resolver struct {
	GuardianService *GuardianService
	ResidentService *ResidentService
	RoomService     *RoomService
	Logger          *logrus.Logger
}

func NewResolver(guardianService *GuardianService, residentService *ResidentService,
	roomService *RoomService, logger *logrus.Logger) *Resolver {

	return &Resolver{
		GuardianService: guardianService,
		ResidentService: residentService,
		RoomService:     roomService,
		Logger:          logger,
	}
}
