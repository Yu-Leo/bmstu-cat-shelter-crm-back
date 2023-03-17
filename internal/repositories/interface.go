package repositories

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
)

type CatRepository interface {
	Create(context.Context, models.CreateCatRequest) (*models.CatId, error)
}
