package user

import (
	"fmt"
	"reflect"
)

type UserName struct {
	Value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < 3 {
		return nil, fmt.Errorf("UserName is more than 3 characters.")
	}
	if len(value) > 20 {
		return nil, fmt.Errorf("UserName is less than 20 characters.")
	}
	return &UserName{Value: value}, nil
}

func (userName *UserName) Equals(other UserName) bool {
	return reflect.DeepEqual(userName.Value, other.Value)
}

func (userName *UserName) String() string {
	return fmt.Sprintf("UserName: [value: %s]", userName.Value)
}
