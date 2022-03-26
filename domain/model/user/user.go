package user

type User struct {
	id   UserId
	name UserName
}

func NewUser(userId UserId, userName UserName) (*User, error) {
	return &User{id: userId, name: userName}, nil
}
