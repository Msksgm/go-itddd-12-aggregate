package circle

import (
	"testing"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"
	"github.com/google/go-cmp/cmp"
)

func Test_NewCircle(t *testing.T) {
	userId, err := user.NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	userName, err := user.NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	owner, err := user.NewUser(*userId, *userName)
	if err != nil {
		t.Fatal(err)
	}

	circleId, err := NewCircleId("circleId")
	if err != nil {
		t.Fatal(err)
	}
	circleName, err := NewCircleName("circlename")
	if err != nil {
		t.Fatal(err)
	}

	members := []user.User{*owner}

	got, err := NewCircle(*circleId, *circleName, *owner, members)
	if err != nil {
		t.Fatal(err)
	}
	want := &Circle{
		id:      *circleId,
		name:    *circleName,
		owner:   *owner,
		members: members,
	}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(CircleId{}, CircleName{}, user.User{}, user.UserId{}, user.UserName{}, Circle{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
