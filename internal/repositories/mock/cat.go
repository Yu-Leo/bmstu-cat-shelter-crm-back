package mock

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
)

type catRepository struct {
	data map[int]models.Cat
}

func NewMockCatRepository() repositories.CatRepository {
	return &catRepository{
		data: map[int]models.Cat{},
	}
}

func (ur *catRepository) Create(ctx context.Context, requestData models.CreateCatRequest) (catId *models.CatId, err error) {
	a := len(ur.data)
	ur.data[a] = models.Cat{Id: a, Name: requestData.Name}
	return &models.CatId{Id: a}, nil
}
