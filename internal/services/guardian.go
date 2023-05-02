package services

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
)

type GuardianService struct {
	repository repositories.GuardianRepository
}

func NewGuardianService(guardianRepository repositories.GuardianRepository) *GuardianService {
	return &GuardianService{
		repository: guardianRepository,
	}
}

func (s GuardianService) CreateGuardian(requestData models.CreateGuardianRequest) (*models.GuardianId, error) {
	return s.repository.Create(context.Background(), requestData)
}

func (s GuardianService) GetGuardiansList() (*[]models.Guardian, error) {
	return s.repository.GetList(context.Background())
}

func (s GuardianService) GetGuardian(id models.GuardianId) (*models.Guardian, error) {
	return s.repository.Get(context.Background(), id)
}

func (s GuardianService) DeleteGuardian(id models.GuardianId) error {
	return s.repository.Delete(context.Background(), id)
}

func (s GuardianService) UpdateGuardian(id models.GuardianId, requestData models.CreateGuardianRequest) error {
	return s.repository.Update(context.Background(), id, requestData)
}
