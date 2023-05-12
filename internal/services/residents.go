package services

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
)

type ResidentService struct {
	repository repositories.ResidentRepository
}

func NewResidentService(residentRepository repositories.ResidentRepository) *ResidentService {
	return &ResidentService{
		repository: residentRepository,
	}
}

func (s ResidentService) CreateResident(requestData models.CreateResidentRequest) (models.CatChipNumber, error) {
	return s.repository.Create(context.Background(), requestData)
}

func (s ResidentService) GetResidentsList() (*[]models.Resident, error) {
	return s.repository.GetList(context.Background())
}

func (s ResidentService) GetResident(catChipNumber models.CatChipNumber) (*models.Resident, error) {
	return s.repository.Get(context.Background(), catChipNumber)
}

func (s ResidentService) DeleteResident(catChipNumber models.CatChipNumber) error {
	return s.repository.Delete(context.Background(), catChipNumber)
}

func (s ResidentService) UpdateResident(catChipNumber models.CatChipNumber, requestData models.CreateResidentRequest) error {
	return s.repository.Update(context.Background(), catChipNumber, requestData)
}
