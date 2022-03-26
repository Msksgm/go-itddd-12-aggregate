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
