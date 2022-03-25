package user

import (
	"fmt"
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

func Test_UserIdEquals(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("equal", func(t *testing.T) {
		otherUserId, err := NewUserId("id")
		if err != nil {
			t.Fatal(err)
		}

		if !userId.Equals(otherUserId) {
			t.Errorf("userId: %v must be equal to otherUserId: %v", userId, otherUserId)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		otherUserId, err := NewUserId("otherId")
		if err != nil {
			t.Fatal(err)
		}

		if userId.Equals(otherUserId) {
			t.Errorf("userId: %v must be equal to otherUserId: %v", userId, otherUserId)
		}
	})
}

func Test_UserIdString(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(userId)
	want := fmt.Sprintf("UserId [value: %s]", userId.value)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
