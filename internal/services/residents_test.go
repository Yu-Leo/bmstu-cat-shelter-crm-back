package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories/mocks"
)

func TestCreateResident(t *testing.T) {
	residentRequest := models.CreateResidentRequest{Booking: false}
	chipNumber := models.CatChipNumber("123456789012345")

	residentRepo := mocks.NewResidentRepository(t)
	residentRepo.On("Create", context.Background(), residentRequest).Return(chipNumber, nil)

	residentService := NewResidentService(residentRepo)

	residentChipNumber, err := residentService.CreateResident(residentRequest)

	assert.Equal(t, chipNumber, residentChipNumber)
	assert.NoError(t, err)
}
