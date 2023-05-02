package repositories

import (
	"context"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"
)

type ResidentRepository interface {
	Create(context.Context, models.CreateResidentRequest) (*models.CatChipNumber, error)
	GetList(context.Context) (*[]models.Resident, error)
	Get(context.Context, models.CatChipNumber) (*models.Resident, error)
	Delete(context.Context, models.CatChipNumber) error
	Update(context.Context, models.CatChipNumber, models.CreateResidentRequest) error
}

type residentRepository struct {
	storage *sqlitedb.Storage
}

func NewSqliteResidentRepository(storage *sqlitedb.Storage) ResidentRepository {
	return &residentRepository{
		storage: storage,
	}
}

func (r *residentRepository) Create(ctx context.Context, rd models.CreateResidentRequest) (_ *models.CatChipNumber, err error) {
	return nil, nil
}

func (r *residentRepository) GetList(ctx context.Context) (list *[]models.Resident, err error) {
	return nil, nil
}

func (r *residentRepository) Get(ctx context.Context, catChipNumber models.CatChipNumber) (_ *models.Resident, err error) {
	return nil, nil
}

func (r *residentRepository) Delete(ctx context.Context, catChipNumber models.CatChipNumber) (err error) {
	return nil
}

func (r *residentRepository) Update(ctx context.Context, catChipNumber models.CatChipNumber, rd models.CreateResidentRequest) (err error) {
	return nil
}
