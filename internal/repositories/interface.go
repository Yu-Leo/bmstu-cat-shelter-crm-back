package repositories

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
)

type CatRepository interface {
	Create(context.Context, models.CreateCatRequest) (*models.CatId, error)
	GetCatsList(context.Context) (*[]models.Cat, error)
	GetCat(context.Context, int) (*models.Cat, error)
}
