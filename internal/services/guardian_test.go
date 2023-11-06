package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories/mocks"
)

func TestCreateGuardian(t *testing.T) {
	guardianRequest := models.CreateGuardianRequest{Firstname: "L", Lastname: "Y"}

	guardianRepo := mocks.NewGuardianRepository(t)
	guardianRepo.On("Create", context.Background(), guardianRequest).Return(models.GuardianId(1), nil)

	guardianService := NewGuardianService(guardianRepo)

	guardianId, err := guardianService.CreateGuardian(guardianRequest)

	assert.Equal(t, models.GuardianId(1), guardianId)
	assert.NoError(t, err)
}
