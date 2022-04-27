package persistence

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/msksgm/go-itddd-12-aggregate/domain/model/circle"
	"github.com/msksgm/go-itddd-12-aggregate/domain/model/user"
)

func Test_Save(t *testing.T) {
	owner := user.User{UserId: user.UserId{Value: "ownerId"}, Name: user.UserName{Value: "ownerName"}}
	member := user.User{UserId: user.UserId{Value: "memberId"}, Name: user.UserName{Value: "memberName"}}
	saveCircle := &circle.Circle{
		Id:      circle.CircleId{Value: "circleId"},
		Name:    circle.CircleName{Value: "circleName"},
		Owner:   owner,
		Members: []user.User{owner, member},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	circleRepository, err := NewCircleRepository(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO circles").
			WithArgs("circleId", "ownerId", "circleName").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO userCircles").
			WithArgs("ownerId", "circleId").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO userCircles").
			WithArgs("memberId", "circleId").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		got := circleRepository.Save(saveCircle)
		if got != nil {
			t.Errorf("got must be nil, but %v", got)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_FindByCircleName(t *testing.T) {
	circleName, _ := circle.NewCircleName("circleName")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	circleRepository, err := NewCircleRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT c.id, c.circlename, c.owner_id, u.id, u.name from circles c JOIN userCircles uc ON c.id = uc.circle_id JOIN users u ON u.id = uc.user_id WHERE c.circlename = $1`)).
			WithArgs("circleName").
			WillReturnRows(mock.NewRows([]string{"circleId", "circleName", "ownerId", "userId", "userName"}).
				AddRow("circleId", "circleName", "ownerId", "ownerId", "ownerName").
				AddRow("circleId", "circleName", "ownerId", "userId", "userName"))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name from users WHERE id = $1`)).
			WithArgs("ownerId").
			WillReturnRows(mock.NewRows([]string{"ownerId", "ownerName"}).AddRow("ownerId", "ownerName"))
		mock.ExpectCommit()

		got, err := circleRepository.FindByCircleName(circleName)
		if err != nil {
			t.Error(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		wantOwner := user.User{UserId: user.UserId{Value: "ownerId"}, Name: user.UserName{Value: "ownerName"}}
		want := &circle.Circle{
			Id:      circle.CircleId{Value: "circleId"},
			Name:    circle.CircleName{Value: "circleName"},
			Owner:   wantOwner,
			Members: []user.User{wantOwner, {UserId: user.UserId{Value: "userId"}, Name: user.UserName{Value: "userName"}}},
		}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported()); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
}
