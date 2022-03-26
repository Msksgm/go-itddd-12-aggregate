package circle

import "github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"

type Circle struct {
	id      CircleId
	name    CircleName
	owner   user.User
	members []user.User
}

func NewCircle(id CircleId, name CircleName, owner user.User, members []user.User) (*Circle, error) {
	return &Circle{id: id, name: name, owner: owner, members: members}, nil
}
