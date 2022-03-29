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

func (cr *CircleRepository) FindByCircleName(ctx context.Context, circleName *circle.CircleName) (*circle.Circle, error) {
	return &circle.Circle{}, nil
}
