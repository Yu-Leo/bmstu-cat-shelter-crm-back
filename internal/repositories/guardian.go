package repositories

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/errors"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"

	sqlite3 "github.com/mattn/go-sqlite3"
)

type GuardianRepository interface {
	Create(context.Context, models.CreateGuardianRequest) (models.GuardianId, error)
	GetList(context.Context) (*[]models.Guardian, error)
	Get(context.Context, models.GuardianId) (*models.Guardian, error)
	Delete(context.Context, models.GuardianId) error
	Update(context.Context, models.GuardianId, models.CreateGuardianRequest) error
}

type guardianRepository struct {
	storage *sqlitedb.Storage
}

func NewSqliteGuardianRepository(storage *sqlitedb.Storage) GuardianRepository {
	return &guardianRepository{
		storage: storage,
	}
}

func (r *guardianRepository) Create(ctx context.Context, rd models.CreateGuardianRequest) (_ models.GuardianId, err error) {
	q1 := `INSERT INTO people (photo_url, firstname, lastname, patronymic, phone)
VALUES (?, ?, ?, ?, ?) RETURNING people.person_id;`

	var personId int
	err = r.storage.DB.QueryRowContext(ctx, q1,
		rd.PhotoURL, rd.Firstname, rd.Lastname, rd.Patronymic, rd.Phone).Scan(&personId)
	if err != nil {
		if strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
			return -1, errors.PersonPhoneAlreadyExists
		}
		return -1, err
	}

	q2 := `INSERT INTO guardians (person_id)
VALUES (?) RETURNING guardians.guardian_id;`

	var guardianId models.GuardianId
	err = r.storage.DB.QueryRowContext(ctx, q2, personId).Scan(&guardianId)
	if err != nil {
		return -1, err
	}
	return guardianId, nil
}

func (r *guardianRepository) GetList(ctx context.Context) (guardiansList *[]models.Guardian, err error) {
	q := `SELECT guardian_id, g.person_id, p.photo_url, p.firstname, p.lastname, p.patronymic, p.phone
FROM guardians as g
JOIN people p on p.person_id = g.person_id;`
	objects := make([]models.Guardian, 0)

	rows, err := r.storage.DB.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		o := models.Guardian{}
		err = rows.Scan(&o.Id, &o.PersonId, &o.PhotoURL, &o.Firstname, &o.Lastname, &o.Patronymic, &o.Phone)
		if err != nil {
			return nil, err
		}
		objects = append(objects, o)
	}

	return &objects, nil
}

func (r *guardianRepository) Get(ctx context.Context, id models.GuardianId) (_ *models.Guardian, err error) {
	q := `SELECT guardian_id, g.person_id, p.photo_url, p.firstname, p.lastname, p.patronymic, p.phone
FROM guardians as g
JOIN people p on p.person_id = g.person_id
WHERE g.guardian_id = ?;`

	o := models.Guardian{}
	err = r.storage.DB.QueryRowContext(ctx, q, id).Scan(
		&o.Id, &o.PersonId, &o.PhotoURL, &o.Firstname, &o.Lastname, &o.Patronymic, &o.Phone)

	if err == sql.ErrNoRows {
		return nil, errors.GuardianNotFound
	}

	return &o, nil
}

func (r *guardianRepository) Delete(ctx context.Context, id models.GuardianId) (err error) {
	var personId int
	q1 := `SELECT person_id
FROM guardians
WHERE guardian_id = ?;`
	err = r.storage.DB.QueryRowContext(ctx, q1, id).Scan(&personId)
	if err == sql.ErrNoRows {
		return errors.GuardianNotFound
	}

	q2 := `DELETE
FROM guardians
WHERE guardian_id = ?;`
	_, err = r.storage.DB.ExecContext(ctx, q2, id)
	if err != nil {
		return err
	}

	q3 := `DELETE
FROM people
WHERE person_id = ?;`
	_, err = r.storage.DB.ExecContext(ctx, q3, personId)
	return err
}

func (r *guardianRepository) Update(ctx context.Context, id models.GuardianId, rd models.CreateGuardianRequest) (err error) {
	var personId int
	q1 := `SELECT person_id
FROM guardians
WHERE guardian_id = ?;`
	err = r.storage.DB.QueryRowContext(ctx, q1, id).Scan(&personId)
	if err == sql.ErrNoRows {
		return errors.GuardianNotFound
	}

	q2 := `UPDATE people
SET photo_url = ?, firstname = ?, lastname = ?, patronymic = ?, phone = ?
WHERE person_id = ?;`

	_, err = r.storage.DB.ExecContext(ctx, q2, rd.PhotoURL, rd.Firstname, rd.Lastname, rd.Patronymic,
		rd.Phone, personId)
	return err
}
