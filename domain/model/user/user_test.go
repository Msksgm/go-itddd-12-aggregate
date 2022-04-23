package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUser(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	got, err := NewUser(*userId, *userName)
	if err != nil {
		t.Fatal(err)
	}
	want := &User{UserId: *userId, Name: *userName}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(User{}, UserId{}, UserName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func Test_ChangeUserName(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	user, err := NewUser(*userId, *userName)
	if err != nil {
		t.Fatal(err)
	}

	changedUserName, err := NewUserName("changedUserName")
	if err != nil {
		t.Fatal(err)
	}

	otherUser := &User{UserId: *userId, Name: *changedUserName}
	user.ChangeName(*changedUserName)

	got := user.Name
	want := otherUser.Name
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
