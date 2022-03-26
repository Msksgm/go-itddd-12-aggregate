package circle

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewCircleId(t *testing.T) {
	got, err := NewCircleId("id")
	if err != nil {
		t.Fatal(err)
	}
	want := &CircleId{value: "id"}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(CircleId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func Test_CircleIdEquals(t *testing.T) {
	circleId, err := NewCircleId("id")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("equal", func(t *testing.T) {
		otherCircleId, err := NewCircleId("id")
		if err != nil {
			t.Fatal(err)
		}

		if !circleId.Equals(otherCircleId) {
			t.Errorf("circleId: %v must be equal to otherCircleId: %v", circleId, otherCircleId)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		otherCircleId, err := NewCircleId("otherId")
		if err != nil {
			t.Fatal(err)
		}

		if circleId.Equals(otherCircleId) {
			t.Errorf("circleId: %v must be equal to otherCircleId: %v", circleId, otherCircleId)
		}
	})
}

func Test_CircleIdString(t *testing.T) {
	circleId, err := NewCircleId("id")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(circleId)
	want := fmt.Sprintf("CircleId [value: %s]", circleId.value)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
