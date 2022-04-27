package persistence

import (
	"database/sql"

	"github.com/msksgm/go-itddd-12-aggregate/domain/model/circle"
	"github.com/msksgm/go-itddd-12-aggregate/domain/model/user"
)

type CircleRepository struct {
	db *sql.DB
}

func NewCircleRepository(db *sql.DB) (*CircleRepository, error) {
	return &CircleRepository{db: db}, nil
}

func (cr *CircleRepository) Save(circle *circle.Circle) (err error) {
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

	_, err = tx.Exec("INSERT INTO circles(id, owner_id, circlename) VALUES ($1, $2, $3)", circle.Id.Value, circle.Owner.UserId.Value, circle.Name.Value)
	if err != nil {
		return err
	}

	for _, member := range circle.Members {
		_, err = tx.Exec("INSERT INTO userCircles(user_id, circle_id) VALUES ($1, $2)", member.UserId.Value, circle.Id.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

type FindByCircleNameQueryError struct {
	CircleName string
	Message    string
	Err        error
}

func (err *FindByCircleNameQueryError) Error() string {
	return err.Message
}

func (cr *CircleRepository) FindByCircleName(circleName *circle.CircleName) (findCircle *circle.Circle, err error) {
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

	rows, err := tx.Query("SELECT c.id, c.circlename, c.owner_id, u.id, u.name from circles c JOIN userCircles uc ON c.id = uc.circle_id JOIN users u ON u.id = uc.user_id WHERE c.circlename = $1", circleName.Value)
	if err != nil {
		return nil, &FindByCircleNameQueryError{CircleName: circleName.Value, Message: "error is occured in circlerepository.FindByCircleName", Err: err}
	}
	defer rows.Close()

	findCircleId := &circle.CircleId{}
	findCircleName := &circle.CircleName{}
	ownerId := &user.UserId{}
	memberId := &user.UserId{}
	memberName := &user.UserName{}
	members := []user.User{}
	for rows.Next() {
		err := rows.Scan(&findCircleId.Value, &findCircleName.Value, &ownerId.Value, &memberId.Value, &memberName.Value)
		if err != nil {
			return nil, err
		}
		members = append(members, user.User{UserId: *memberId, Name: *memberName})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if ownerId == nil {
		return nil, nil
	}
	rows, err = tx.Query("SELECT id, name from users WHERE id = $1", ownerId.Value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ownerName := &user.UserName{}
	for rows.Next() {
		err := rows.Scan(&ownerId.Value, &ownerName.Value)
		if err != nil {
			return nil, err
		}
		owner := &user.User{UserId: *ownerId, Name: *ownerName}
		findCircle = &circle.Circle{Id: *findCircleId, Name: *findCircleName, Owner: *owner, Members: members}
	}

	return findCircle, nil
}
