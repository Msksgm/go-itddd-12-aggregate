package circle

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewCircleName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewCircleName("username")
		if err != nil {
			t.Fatal(err)
		}

		want := &CircleName{value: "username"}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(CircleName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail because value is less than 3 characters", func(t *testing.T) {
		_, err := NewCircleName("us")
		want := "CircleName is more than 3 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail because value is more than 20 characters", func(t *testing.T) {
		_, err := NewCircleName("usernameusernameusername")
		want := "CircleName is less than 20 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func Test_CircleNameEquals(t *testing.T) {
	circleName, err := NewCircleName("username")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("equal", func(t *testing.T) {
		otherCircleName, err := NewCircleName("username")
		if err != nil {
			t.Fatal(err)
		}
		if !circleName.Equals(*otherCircleName) {
			t.Errorf("circleName %v must be equal to otherCircleName %v", circleName, otherCircleName)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		otherCircleName, err := NewCircleName("otherusername")
		if err != nil {
			t.Fatal(err)
		}
		if circleName.Equals(*otherCircleName) {
			t.Errorf("circleName %v must not be equal to otherCircleName %v", circleName, otherCircleName)
		}
	})
}

func Test_CircleNameString(t *testing.T) {
	circleName, err := NewCircleName("username")
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("CircleName: [value: %s]", circleName.value)
	got := fmt.Sprint(circleName)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
