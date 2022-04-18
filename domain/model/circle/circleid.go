package circle

import (
	"fmt"
	"reflect"
)

type CircleId struct {
	Value string
}

func NewCircleId(id string) (*CircleId, error) {
	return &CircleId{Value: id}, nil
}

func (circleId *CircleId) Equals(other *CircleId) bool {
	return reflect.DeepEqual(circleId.Value, other.Value)
}

func (circleId *CircleId) String() string {
	return fmt.Sprintf("CircleId [value: %s]", circleId.Value)
}
