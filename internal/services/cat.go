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

func (s CatService) CreateCat(requestData models.CreateCatRequest) (*models.CatId, error) {
	return s.repository.Create(context.Background(), requestData)
}

func (s CatService) GetCatsList() (*[]models.Cat, error) {
	return s.repository.GetCatsList(context.Background())
}

func (s CatService) GetCat(catId int) (*models.Cat, error) {
	return s.repository.GetCat(context.Background(), catId)
}
