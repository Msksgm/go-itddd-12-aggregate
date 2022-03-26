package circle

import (
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
