package user

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	return &UserId{value: value}, nil
}
