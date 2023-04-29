package repositories

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"

	sqlite3 "github.com/mattn/go-sqlite3"
)

type CatRepository interface {
	Create(context.Context, models.CreateCatRequest) (*models.CatChipNumber, error)
	GetCatsList(context.Context) (*[]models.Cat, error)
	GetCat(context.Context, models.CatChipNumber) (*models.Cat, error)
	DeleteCat(context.Context, models.CatChipNumber) error
}

type catRepository struct {
	storage *sqlitedb.Storage
}

func NewSqliteCatRepository(storage *sqlitedb.Storage) CatRepository {
	return &catRepository{
		storage: storage,
	}
}

func (cr *catRepository) Create(ctx context.Context, rd models.CreateCatRequest) (catId *models.CatChipNumber, err error) {
	q := `INSERT INTO cats (nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter)
VALUES (?,?, ?, ?, ?, ?) RETURNING cats.chip_number;`

	var chipNumber string

	err = cr.storage.DB.QueryRowContext(ctx, q,
		rd.Nickname, rd.PhotoUrl, rd.Gender, rd.Age, rd.ChipNumber, rd.DateOfAdmissionToShelter).Scan(&chipNumber)

	if err != nil {
		if strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
			return nil, apperror.ValidationError
		}
		return nil, err
	}
	return &models.CatChipNumber{ChipNumber: chipNumber}, nil
}

func (cr *catRepository) GetCatsList(ctx context.Context) (catsList *[]models.Cat, err error) {
	q := `SELECT nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter FROM cats;`
	answer := make([]models.Cat, 0)

	rows, err := cr.storage.DB.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cat := models.Cat{}
		err = rows.Scan(&cat.Nickname, &cat.PhotoUrl, &cat.Gender, &cat.Age, &cat.ChipNumber, &cat.DateOfAdmissionToShelter)
		if err != nil {
			return nil, err
		}
		answer = append(answer, cat)
	}

	return &answer, nil
}

func (cr *catRepository) GetCat(ctx context.Context, catChipNumber models.CatChipNumber) (*models.Cat, error) {
	q := `SELECT nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter FROM cats
		WHERE chip_number = ?;`
	cat := models.Cat{}
	err := cr.storage.DB.QueryRow(q, catChipNumber.ChipNumber).Scan(&cat.Nickname, &cat.PhotoUrl, &cat.Gender, &cat.Age, &cat.ChipNumber, &cat.DateOfAdmissionToShelter)

	if err == sql.ErrNoRows {
		return nil, apperror.CatNotFound
	}

	return &cat, nil
}

func (cr *catRepository) DeleteCat(ctx context.Context, catChipNumber models.CatChipNumber) error {
	q := `DELETE
FROM cats
WHERE chip_number = ?;`
	_, err := cr.storage.DB.Exec(q, catChipNumber.ChipNumber)
	return err
}
