package user

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	return &UserName{value: value}, nil
}
