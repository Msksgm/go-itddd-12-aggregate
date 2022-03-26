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
	want := &User{id: *userId, name: *userName}
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

	otherUser := &User{id: *userId, name: *changedUserName}
	user.ChangeName(*changedUserName)

	got := user.name
	want := otherUser.name
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
