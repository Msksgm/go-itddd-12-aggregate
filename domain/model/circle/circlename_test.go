package circle

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewCircleName(t *testing.T) {
	data := []struct {
		testname string
		userName string
		want     *CircleName
		errMsg   string
	}{
		{"success", "username", &CircleName{value: "username"}, ""},
		{"fail because value is less than 3 characters", "us", nil, "CircleName is more than 3 characters."},
		{"fail because value is more than 20 characters", "usernameusernameusername", nil, "CircleName is less than 20 characters."},
	}
	for _, d := range data {
		t.Run("success", func(t *testing.T) {
			got, err := NewCircleName(d.userName)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported(CircleName{})); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
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
