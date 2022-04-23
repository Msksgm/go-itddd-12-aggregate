package user

type User struct {
	UserId UserId
	Name   UserName
}

func NewUser(userId UserId, userName UserName) (*User, error) {
	return &User{UserId: userId, Name: userName}, nil
}

func (user *User) ChangeName(userName UserName) {
	user.Name = userName
}

func (user *User) Id() *UserId {
	return &user.UserId
}
