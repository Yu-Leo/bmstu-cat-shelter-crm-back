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
