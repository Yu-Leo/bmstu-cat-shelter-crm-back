package sqlite

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/repositories"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"
)

type catRepository struct {
	storage *sqlitedb.Storage
}

func NewSqliteCatRepository(storage *sqlitedb.Storage) repositories.CatRepository {
	return &catRepository{
		storage: storage,
	}
}

func (cr *catRepository) Create(ctx context.Context, requestData models.CreateCatRequest) (catId *models.CatId, err error) {
	q := `INSERT INTO cats (name) VALUES (?) RETURNING id`

	var id int
	err = cr.storage.DB.QueryRowContext(ctx, q, requestData.Name).Scan(&id)

	if err != nil {
		return nil, err
	}
	return &models.CatId{Id: id}, nil
}
