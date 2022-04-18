package circle

import "github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"

type Circle struct {
	Id      CircleId
	Name    CircleName
	Owner   user.User
	Members []user.User
}

func NewCircle(id CircleId, name CircleName, owner user.User, members []user.User) (*Circle, error) {
	return &Circle{Id: id, Name: name, Owner: owner, Members: members}, nil
}

func (circle *Circle) IsFull() bool {
	return len(circle.Members) >= 29
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
		return &CircleIsFullError{CircleId: circle.Id, Message: "cannnot join member because the circle is full"}
	}
	circle.Members = append(circle.Members, *newMember)
	return nil
}

type MemberIsNotFoundError struct {
	MemberId user.UserId
	Message  string
}

func (minfe *MemberIsNotFoundError) Error() string {
	return minfe.Message
}

func (circle *Circle) ChangeMemberName(memberId *user.UserId, changedUserName *user.UserName) error {
	for i, member := range circle.Members {
		if member.Id().Equals(memberId) {
			circle.Members[i].ChangeName(*changedUserName)
			return nil
		}
	}
	return &MemberIsNotFoundError{MemberId: *memberId, Message: "member is not found"}
}
