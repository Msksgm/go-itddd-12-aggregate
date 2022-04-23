package user

import (
	"fmt"
	"reflect"
)

type UserId struct {
	Value string
}

func NewUserId(value string) (*UserId, error) {
	return &UserId{Value: value}, nil
}

func (userId *UserId) Equals(other *UserId) bool {
	return reflect.DeepEqual(userId.Value, other.Value)
}

func (userId *UserId) String() string {
	return fmt.Sprintf("UserId [value: %s]", userId.Value)
}
