package circle

import (
	"testing"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"
)

type CircleRepositorierStub struct {
	save             func(circle Circle) error
	findByCircleName func(circleName CircleName) (*Circle, error)
}

func (crs CircleRepositorierStub) Save(circle *Circle) error {
	return crs.save(*circle)
}

func (crs CircleRepositorierStub) FindByCircleName(circleName *CircleName) (*Circle, error) {
	return crs.findByCircleName(*circleName)
}

func Test_Exists(t *testing.T) {
	ownerId, _ := user.NewUserId("ownerId")
	ownerName, _ := user.NewUserName("ownerName")
	owner, _ := user.NewUser(*ownerId, *ownerName)
	circleId := &CircleId{Value: "circleId"}
	circleName, _ := NewCircleName("circlename")
	members := []user.User{*owner}
	data := []struct {
		testname         string
		findByCircleName func(circleName CircleName) (*Circle, error)
		want             bool
		circle           *Circle
		testErrMsg       string
	}{
		{
			"exists",
			func(circleName CircleName) (*Circle, error) {
				return &Circle{Id: *circleId, Name: CircleName{Value: "circlename"}, Owner: *owner, Members: members}, nil
			},
			true,
			&Circle{Id: *circleId, Name: *circleName, Owner: *owner, Members: members},
			"userService.Exists must be true but false",
		},
		{
			"not exists",
			func(circleName CircleName) (*Circle, error) {
				return nil, nil
			},
			false,
			&Circle{Id: *circleId, Name: *circleName, Owner: *owner, Members: members},
			"userService.Exists must be false but true",
		},
	}
	circleServie := CircleService{}

	for _, d := range data {
		t.Run(d.testname, func(t *testing.T) {
			circleServie.circleRepository = &CircleRepositorierStub{findByCircleName: d.findByCircleName}
			got, err := circleServie.Exists(d.circle)
			if err != nil {
				t.Fatal(err)
			}
			if got != d.want {
				t.Errorf("got %v, want %v", got, d.want)
			}
		})
	}
}
