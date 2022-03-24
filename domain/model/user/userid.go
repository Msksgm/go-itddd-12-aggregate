package user

import (
	"fmt"
	"reflect"
)

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	return &UserId{value: value}, nil
}

func (userId *UserId) Equals(other *UserId) bool {
	return reflect.DeepEqual(userId.value, other.value)
}

func (userId *UserId) String() string {
	return fmt.Sprintf("UserId [value: %s]", userId.value)
}
