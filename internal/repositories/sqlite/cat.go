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

func (cr *catRepository) Create(ctx context.Context, rd models.CreateCatRequest) (catId *models.CatId, err error) {
	q := `INSERT INTO cats (nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter)
VALUES (?,?, ?, ?, ?, ?) RETURNING cats.id;`

	var id int
	err = cr.storage.DB.QueryRowContext(ctx, q,
		rd.Nickname, rd.PhotoUrl, rd.Gender, rd.Age, rd.ChipNumber, rd.DateOfAdmissionToShelter).Scan(&id)

	if err != nil {
		return nil, err
	}
	return &models.CatId{Id: id}, nil
}

func (cr *catRepository) GetCatsList(ctx context.Context) (catsList *[]models.Cat, err error) {
	q := `SELECT id, nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter FROM cats;`
	answer := make([]models.Cat, 0)

	rows, err := cr.storage.DB.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cat := models.Cat{}
		err = rows.Scan(&cat.Id, &cat.Nickname, &cat.PhotoUrl, &cat.Gender, &cat.Age, &cat.ChipNumber, &cat.DateOfAdmissionToShelter)
		if err != nil {
			return nil, err
		}
		answer = append(answer, cat)
	}

	return &answer, nil
}
