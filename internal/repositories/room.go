package repositories

import (
	"context"
	"strings"

	"github.com/mattn/go-sqlite3"

	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/errors"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/internal/models"
	"github.com/Yu-Leo/bmstu-cat-shelter-crm-back/pkg/sqlitedb"
)

type RoomRepository interface {
	Create(context.Context, models.CreateRoomRequest) (models.RoomNumber, error)
	GetList(context.Context) (*[]models.Room, error)
	Get(context.Context, models.RoomNumber) (*models.Room, error)
	Delete(context.Context, models.RoomNumber) error
	Update(context.Context, models.RoomNumber, models.CreateRoomRequest) error
}

type roomRepository struct {
	storage *sqlitedb.Storage
}

func NewSqliteRoomRepository(storage *sqlitedb.Storage) RoomRepository {
	return &roomRepository{
		storage: storage,
	}
}

func (r *roomRepository) Create(ctx context.Context, rd models.CreateRoomRequest) (roomNumber models.RoomNumber, err error) {
	q := `INSERT INTO rooms (number, status)
VALUES (?, ?) RETURNING rooms.number;`

	var number models.RoomNumber

	err = r.storage.DB.QueryRowContext(ctx, q, rd.Number, rd.Status).Scan(&number)

	if err != nil {
		if strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
			return "", errors.RoomNumberAlreadyExists
		}
		return "", err
	}
	return number, nil
}

func (r *roomRepository) GetList(ctx context.Context) (roomsList *[]models.Room, err error) {
	return nil, nil
}

func (r *roomRepository) Get(ctx context.Context, roomNumber models.RoomNumber) (*models.Room, error) {
	return nil, nil
}

func (r *roomRepository) Delete(ctx context.Context, roomNumber models.RoomNumber) error {
	return nil
}

func (r *roomRepository) Update(ctx context.Context, roomNumber models.RoomNumber, rd models.CreateRoomRequest) error {
	return nil
}
