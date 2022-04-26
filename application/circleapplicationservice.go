package application

import (
	"fmt"
	"log"

	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/circle"
	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/user"
)

type CircleApplicationService struct {
	circleRepository circle.CircleRepositorier
	circleService    circle.CircleService
}

func NewCircleApplicationService(circleRepository circle.CircleRepositorier, circleService circle.CircleService) (*CircleApplicationService, error) {
	return &CircleApplicationService{circleRepository: circleRepository, circleService: circleService}, nil
}

func (cas *CircleApplicationService) Register(circleName string) error {
	newCircleId, err := circle.NewCircleId("test-circle-id")
	if err != nil {
		return nil
	}
	newCircleName, err := circle.NewCircleName(circleName)
	if err != nil {
		return nil
	}

	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		return nil
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		return nil
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
	if err != nil {
		return nil
	}

	memberId, err := user.NewUserId("memberId")
	if err != nil {
		return nil
	}
	memberName, err := user.NewUserName("memberName")
	if err != nil {
		return nil
	}
	member, err := user.NewUser(*memberId, *memberName)
	if err != nil {
		return nil
	}

	members := []user.User{*owner, *member}
	newCircle, err := circle.NewCircle(*newCircleId, *newCircleName, *owner, members)
	if err != nil {
		return nil
	}
	isCircleExists, err := cas.circleService.Exists(newCircle)
	if err != nil {
		return err
	}
	if isCircleExists {
		return fmt.Errorf("circleName of %s is already exists.", circleName)
	}

	if err := cas.circleRepository.Save(newCircle); err != nil {
		return err
	}
	log.Println("success fully saved")
	return nil
}
