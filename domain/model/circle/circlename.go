package circle

import (
	"fmt"
	"reflect"
)

type CircleName struct {
	Value string
}

func NewCircleName(value string) (*CircleName, error) {
	if len(value) < 3 {
		return nil, fmt.Errorf("CircleName is more than 3 characters.")
	}
	if len(value) > 20 {
		return nil, fmt.Errorf("CircleName is less than 20 characters.")
	}
	return &CircleName{Value: value}, nil
}

func (circleName *CircleName) Equals(other CircleName) bool {
	return reflect.DeepEqual(circleName.Value, other.Value)
}

func (circleName *CircleName) String() string {
	return fmt.Sprintf("CircleName: [value: %s]", circleName.Value)
}
