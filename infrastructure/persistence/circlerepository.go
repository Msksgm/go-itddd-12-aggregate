package persistence

import (
	"context"
	"database/sql"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/circle"
)

type CircleRepository struct {
	db *sql.DB
}

func NewCircleRepository(db *sql.DB) (*CircleRepository, error) {
	return &CircleRepository{db: db}, nil
}

type FindByCircleNameQueryError struct {
	CircleName string
	Message    string
	Err        error
}

func (err *FindByCircleNameQueryError) Error() string {
	return err.Message
}

func (cr *CircleRepository) FindByCircleName(ctx context.Context, circleName *circle.CircleName) (circle *circle.Circle, err error) {
	tx, err := cr.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM circles WHERE circlename = $1", circleName.Value())
	if err != nil {
		return nil, &FindByCircleNameQueryError{CircleName: circleName.Value(), Message: "error is occured in circlerepository.FindByCircleName", Err: err}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan()
		if err != nil {
			return nil, err
		}
	}
	return circle, nil
}
