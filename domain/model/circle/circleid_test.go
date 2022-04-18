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
	want := &CircleId{Value: "id"}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(CircleId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func Test_CircleIdEquals(t *testing.T) {
	data := []struct {
		testname string
		id       string
		want     bool
	}{
		{"equal", "id", true},
		{"not equal", "otherId", false},
	}
	circleId, err := NewCircleId("id")
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range data {
		t.Run(d.testname, func(t *testing.T) {
			otherCircleId, err := NewCircleId(d.id)
			if err != nil {
				t.Fatal(err)
			}

			got := circleId.Equals(otherCircleId)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported()); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}

func Test_CircleIdString(t *testing.T) {
	circleId, err := NewCircleId("id")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(circleId)
	want := fmt.Sprintf("CircleId [value: %s]", circleId.Value)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
