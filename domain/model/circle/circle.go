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

func (circle *Circle) IsFull() bool {
	return len(circle.members) >= 29
}

type CircleIsFullError struct {
	Member  user.User
	Message string
}

func (cife *CircleIsFullError) Error() string {
	return cife.Message
}

func (circle *Circle) Join(newMember *user.User) error {
	if circle.IsFull() {
		return &CircleIsFullError{Member: *newMember, Message: "cannnot join member because the circle is full"}
	}
	circle.members = append(circle.members, *newMember)
	return nil
}
