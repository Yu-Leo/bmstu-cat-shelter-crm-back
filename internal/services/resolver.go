package services

import "github.com/sirupsen/logrus"

type Resolver struct {
	CatService      *CatService
	GuardianService *GuardianService
	ResidentService *ResidentService
	Logger          *logrus.Logger
}

func NewResolver(catService *CatService,
	guardianService *GuardianService,
	residentService *ResidentService, logger *logrus.Logger) *Resolver {

	return &Resolver{
		CatService:      catService,
		GuardianService: guardianService,
		ResidentService: residentService,
		Logger:          logger,
	}
}
