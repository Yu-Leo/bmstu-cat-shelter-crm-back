package services

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
)

type CatService struct {
	repository repositories.CatRepository
}

func NewCatService(catRepository repositories.CatRepository) *CatService {
	return &CatService{
		repository: catRepository,
	}
}

func (s CatService) CreateCat(requestData models.CreateCatRequest) (*models.CatChipNumber, error) {
	return s.repository.Create(context.Background(), requestData)
}

func (s CatService) GetCatsList() (*[]models.Cat, error) {
	return s.repository.GetCatsList(context.Background())
}

func (s CatService) GetCat(catChipNumber models.CatChipNumber) (*models.Cat, error) {
	return s.repository.GetCat(context.Background(), catChipNumber)
}

func (s CatService) DeleteCat(catChipNumber models.CatChipNumber) error {
	return s.repository.DeleteCat(context.Background(), catChipNumber)
}

func (s CatService) UpdateCat(catChipNumber models.CatChipNumber, requestData models.CreateCatRequest) error {
	return s.repository.UpdateCat(context.Background(), catChipNumber, requestData)
}
