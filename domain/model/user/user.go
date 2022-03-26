package user

type User struct {
	id   UserId
	name UserName
}

func NewUser(userId UserId, userName UserName) (*User, error) {
	return &User{id: userId, name: userName}, nil
}

func (user *User) ChangeName(userName UserName) {
	user.name = userName
}

func (user *User) Id() *UserId {
	return &user.id
}
