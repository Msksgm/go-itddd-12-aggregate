package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserName(t *testing.T) {
	got, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}

	want := &UserName{value: "username"}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
