package circle

import (
	"fmt"
	"reflect"
)

type CircleId struct {
	value string
}

func NewCircleId(id string) (*CircleId, error) {
	return &CircleId{value: id}, nil
}

func (circleId *CircleId) Equals(other *CircleId) bool {
	return reflect.DeepEqual(circleId.value, other.value)
}

func (circleId *CircleId) String() string {
	return fmt.Sprintf("CircleId [value: %s]", circleId.value)
}
