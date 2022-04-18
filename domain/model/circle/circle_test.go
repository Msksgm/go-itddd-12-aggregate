package circle

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"
	"github.com/google/go-cmp/cmp"
)

func Test_NewCircle(t *testing.T) {
	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		t.Fatal(err)
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		t.Fatal(err)
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
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

func Test_IsFull(t *testing.T) {
	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		t.Fatal(err)
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		t.Fatal(err)
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
	if err != nil {
		t.Fatal(err)
	}
	memberId, err := user.NewUserId("memberId")
	if err != nil {
		t.Fatal(err)
	}
	memberName, err := user.NewUserName("memberName")
	if err != nil {
		t.Fatal(err)
	}
	member, err := user.NewUser(*memberId, *memberName)
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
	t.Run("full", func(t *testing.T) {
		members := []user.User{*owner}

		for i := 0; i < 30; i++ {
			members = append(members, *member)
		}

		circle, err := NewCircle(*circleId, *circleName, *owner, members)
		if err != nil {
			t.Fatal(err)
		}
		if !circle.IsFull() {
			t.Error("circle must be less than 29 members")
		}
	})
	t.Run("not full", func(t *testing.T) {
		members := []user.User{*owner}

		for i := 0; i < 20; i++ {
			members = append(members, *member)
		}

		circle, err := NewCircle(*circleId, *circleName, *owner, members)
		if err != nil {
			t.Fatal(err)
		}
		if circle.IsFull() {
			t.Error("circle must be less than 29 members")
		}
	})
}

func Test_Join(t *testing.T) {
	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		t.Fatal(err)
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		t.Fatal(err)
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
	if err != nil {
		t.Fatal(err)
	}

	memberId, err := user.NewUserId("memberId")
	if err != nil {
		t.Fatal(err)
	}
	memberName, err := user.NewUserName("memberName")
	if err != nil {
		t.Fatal(err)
	}
	member, err := user.NewUser(*memberId, *memberName)
	if err != nil {
		t.Fatal(err)
	}

	circleId := &CircleId{Value: "circleId"}
	circleName := &CircleName{Value: "circlename"}

	members := []user.User{*owner}

	circle := &Circle{*circleId, *circleName, *owner, members}
	t.Run("success", func(t *testing.T) {
		wantMembers := append(members, *member)
		want := &Circle{
			id:      *circleId,
			name:    *circleName,
			owner:   *owner,
			members: wantMembers,
		}

		if err := circle.Join(member); err != nil {
			t.Error(err)
		}
		if diff := cmp.Diff(want, circle, cmp.AllowUnexported(CircleId{}, CircleName{}, user.User{}, user.UserId{}, user.UserName{}, Circle{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail", func(t *testing.T) {
		// add members to full amount
		for i := 0; i < 29; i++ {
			members = append(members, *member)
		}
		circle, err := NewCircle(*circleId, *circleName, *owner, members)
		if err != nil {
			t.Fatal(err)
		}

		var circleIsFullError *CircleIsFullError

		err = circle.Join(member)
		if !errors.As(err, &circleIsFullError) {
			t.Errorf("err type: %v, expect err type: %v", reflect.TypeOf(err), reflect.TypeOf(circleIsFullError))
		}
	})
}

func Test_ChangeMemberName(t *testing.T) {
	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		t.Fatal(err)
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		t.Fatal(err)
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
	if err != nil {
		t.Fatal(err)
	}

	circleId := &CircleId{Value: "circleId"}
	circleName := &CircleName{Value: "circlename"}

	memberId, err := user.NewUserId("memberId")
	if err != nil {
		t.Fatal(err)
	}
	memberName, err := user.NewUserName("memberName")
	if err != nil {
		t.Fatal(err)
	}
	member, err := user.NewUser(*memberId, *memberName)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("succes", func(t *testing.T) {
		members := []user.User{*owner, *member}

		circle, err := NewCircle(*circleId, *circleName, *owner, members)
		if err != nil {
			t.Fatal(err)
		}

		changedUserName, err := user.NewUserName("changedMemberName")
		if err != nil {
			t.Fatal(err)
		}

		chagedUserNameMember, err := user.NewUser(*memberId, *changedUserName)
		if err != nil {
			t.Fatal(err)
		}

		if err := circle.ChangeMemberName(memberId, changedUserName); err != nil {
			t.Error(err)
		}
		got := circle
		want := &Circle{
			id:      *circleId,
			name:    *circleName,
			owner:   *owner,
			members: []user.User{*owner, *chagedUserNameMember},
		}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(Circle{}, CircleId{}, CircleName{}, user.User{}, user.UserId{}, user.UserName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail", func(t *testing.T) {
		// member is not included in members
		members := []user.User{*owner}

		circle, err := NewCircle(*circleId, *circleName, *owner, members)
		if err != nil {
			t.Fatal(err)
		}

		changedUserName, err := user.NewUserName("changedMemberName")
		if err != nil {
			t.Fatal(err)
		}

		var memberIsNotFoundError *MemberIsNotFoundError
		err = circle.ChangeMemberName(memberId, changedUserName)
		if !errors.As(err, &memberIsNotFoundError) {
			t.Errorf("err type: %v, expect err type: %v", reflect.TypeOf(err), reflect.TypeOf(memberIsNotFoundError))
		}
	})
}
