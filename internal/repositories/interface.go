package repositories

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
)

type CatRepository interface {
	Create(context.Context, models.CreateCatRequest) (*models.CatChipNumber, error)
	GetCatsList(context.Context) (*[]models.Cat, error)
	GetCat(context.Context, models.CatChipNumber) (*models.Cat, error)
}
