package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewUserName("username")
		if err != nil {
			t.Fatal(err)
		}

		want := &UserName{value: "username"}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail", func(t *testing.T) {
		_, err := NewUserName("us")
		want := "UserName is more than 3 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
