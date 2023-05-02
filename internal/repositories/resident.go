package repositories

import (
	"context"
	"database/sql"
	"strings"

	"github.com/mattn/go-sqlite3"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/apperror"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"
)

type ResidentRepository interface {
	Create(context.Context, models.CreateResidentRequest) (models.CatChipNumber, error)
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

func (r *residentRepository) Create(ctx context.Context, rd models.CreateResidentRequest) (_ models.CatChipNumber, err error) {
	q1 := `INSERT INTO cats (nickname, photo_url, gender, age, chip_number, date_of_admission_to_shelter)
VALUES (?,?, ?, ?, ?, ?) RETURNING cats.chip_number;`

	var chipNumber models.CatChipNumber

	err = r.storage.DB.QueryRowContext(ctx, q1,
		rd.Nickname, rd.PhotoUrl, rd.Gender, rd.Age, rd.CatChipNumber, rd.DateOfAdmissionToShelter).Scan(&chipNumber)

	if err != nil {
		if strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
			return "", apperror.CatChipNumberAlreadyExists
		}
		return "", err
	}

	q2 := `INSERT INTO residents (cat_chip_number, booking, aggressiveness, vk_album_url, guardian_id)
VALUES (?, ?, ?, ?, ?); `

	_, err = r.storage.DB.ExecContext(ctx, q2,
		chipNumber, rd.Booking, rd.Aggressiveness, rd.VKAlbumUrl, rd.GuardianId)

	if err != nil {
		if strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
			return "", apperror.CatChipNumberAlreadyExists
		}
		return "", err
	}
	return chipNumber, nil
}

func (r *residentRepository) GetList(ctx context.Context) (list *[]models.Resident, err error) {
	q := `SELECT c.chip_number, c.nickname, c.photo_url, c.gender,
       c.age, c.date_of_admission_to_shelter, r.booking, r.aggressiveness, r.vk_album_url, r.guardian_id
FROM residents as r
JOIN cats c on c.chip_number = r.cat_chip_number;`
	answer := make([]models.Resident, 0)

	rows, err := r.storage.DB.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ru := models.Resident{}
		err = rows.Scan(&ru.CatChipNumber, &ru.Nickname, &ru.PhotoUrl, &ru.Gender, &ru.Age,
			&ru.DateOfAdmissionToShelter, &ru.Booking, &ru.Aggressiveness, &ru.VKAlbumUrl, &ru.GuardianId)
		if err != nil {
			return nil, err
		}
		answer = append(answer, ru)
	}

	return &answer, nil
}

func (r *residentRepository) Get(ctx context.Context, catChipNumber models.CatChipNumber) (_ *models.Resident, err error) {
	q := `SELECT c.chip_number, c.nickname, c.photo_url, c.gender,
       c.age, c.date_of_admission_to_shelter, r.booking, r.aggressiveness, r.vk_album_url, r.guardian_id
FROM residents as r
JOIN cats c on c.chip_number = r.cat_chip_number
WHERE r.cat_chip_number = ?;`

	ru := models.Resident{}
	err = r.storage.DB.QueryRow(q, catChipNumber).Scan(&ru.CatChipNumber, &ru.Nickname, &ru.PhotoUrl, &ru.Gender, &ru.Age,
		&ru.DateOfAdmissionToShelter, &ru.Booking, &ru.Aggressiveness, &ru.VKAlbumUrl, &ru.GuardianId)
	if err == sql.ErrNoRows {
		return nil, apperror.ResidentNotFound
	}

	return &ru, nil
}

func (r *residentRepository) Delete(ctx context.Context, catChipNumber models.CatChipNumber) (err error) {
	q1 := `DELETE
FROM residents
WHERE cat_chip_number = ?;`
	_, err = r.storage.DB.Exec(q1, catChipNumber)
	if err != nil {
		return err
	}

	q2 := `DELETE
FROM cats
WHERE chip_number = ?;`
	_, err = r.storage.DB.Exec(q2, catChipNumber)
	return err
}

func (r *residentRepository) Update(ctx context.Context, catChipNumber models.CatChipNumber, rd models.CreateResidentRequest) (err error) {
	q1 := `UPDATE residents
SET booking = ?, aggressiveness = ?, vk_album_url = ?, guardian_id = ?
WHERE cat_chip_number = ?;`

	_, err = r.storage.DB.Exec(q1, rd.CatChipNumber, rd.Booking, rd.Aggressiveness, rd.VKAlbumUrl, rd.GuardianId, catChipNumber)
	if err != nil {
		return err
	}

	q2 := `UPDATE cats
SET nickname = ?, photo_url = ?, gender = ?, age = ?, chip_number = ?, date_of_admission_to_shelter = ? 
WHERE chip_number = ?;`

	_, err = r.storage.DB.Exec(q2, rd.Nickname, rd.PhotoUrl, rd.Gender, rd.Age, rd.CatChipNumber, rd.DateOfAdmissionToShelter, catChipNumber)
	return err
}
