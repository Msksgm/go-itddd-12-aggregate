package circle

import (
	"testing"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"
)

type ICircleRepositoryStub struct{}

func (cs ICircleRepositoryStub) FindByCircleName(circleName *CircleName) (*Circle, error) {
	ownerId, _ := user.NewUserId("ownerId")
	ownerName, _ := user.NewUserName("ownerName")
	owner, _ := user.NewUser(*ownerId, *ownerName)
	circleId := &CircleId{value: "circleId"}
	members := []user.User{*owner}
	expectedCircleName, _ := NewCircleName("circlename")
	expectedCircle := &Circle{
		id:      *circleId,
		name:    *expectedCircleName,
		owner:   *owner,
		members: members,
	}
	if circleName.Equals(*expectedCircleName) {
		return expectedCircle, nil
	}
	return nil, nil
}

func Test_Exists(t *testing.T) {
	circleService := CircleService{circleRepository: ICircleRepositoryStub{}}

	ownerId, _ := user.NewUserId("ownerId")
	ownerName, _ := user.NewUserName("ownerName")
	owner, _ := user.NewUser(*ownerId, *ownerName)
	circleId := &CircleId{value: "circleId"}
	members := []user.User{*owner}
	circleName, _ := NewCircleName("circlename")
	circle := &Circle{
		id:      *circleId,
		name:    *circleName,
		owner:   *owner,
		members: members,
	}
	isExists, err := circleService.Exists(circle)
	if err != nil {
		t.Fatal(err)
	}
	if !isExists {
		t.Errorf("isExists must be %v but %v", isExists, isExists)
	}
}
