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
	CircleId CircleId
	Message  string
}

func (cife *CircleIsFullError) Error() string {
	return cife.Message
}

func (circle *Circle) Join(newMember *user.User) error {
	if circle.IsFull() {
		return &CircleIsFullError{CircleId: circle.id, Message: "cannnot join member because the circle is full"}
	}
	circle.members = append(circle.members, *newMember)
	return nil
}

func (circle *Circle) ChangeMemberName(memberId *user.UserId, changedUserName *user.UserName) error {
	for i, member := range circle.members {
		if member.Id().Equals(memberId) {
			circle.members[i].ChangeName(*changedUserName)
			return nil
		}
	}
	return nil
}
