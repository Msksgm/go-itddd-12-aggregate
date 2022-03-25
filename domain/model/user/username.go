package user

import "fmt"

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < 3 {
		return nil, fmt.Errorf("UserName is more than 3 characters.")
	}
	return &UserName{value: value}, nil
}
