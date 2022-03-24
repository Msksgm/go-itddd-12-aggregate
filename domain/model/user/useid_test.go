package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserId(t *testing.T) {
	got, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}

	want := &UserId{value: "id"}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
