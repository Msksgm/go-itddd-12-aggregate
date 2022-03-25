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
	t.Run("fail because value is less than 3 characters", func(t *testing.T) {
		_, err := NewUserName("us")
		want := "UserName is more than 3 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail because value is more than 20 characters", func(t *testing.T) {
		_, err := NewUserName("usernameusernameusername")
		want := "UserName is less than 20 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func Test_UserNameEquals(t *testing.T) {
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	otherUserName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	if !userName.Equals(*otherUserName) {
		t.Errorf("userName %v must be equal to otherUserName %v", userName, otherUserName)
	}
}
