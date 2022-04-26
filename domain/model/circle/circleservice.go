package circle

type CircleService struct {
	circleRepository CircleRepositorier
}

func NewCircleService(circleRepository CircleRepositorier) (*CircleService, error) {
	return &CircleService{circleRepository: circleRepository}, nil
}

func (circleService *CircleService) Exists(circle *Circle) (bool, error) {
	circle, err := circleService.circleRepository.FindByCircleName(&circle.Name)
	if err != nil {
		return false, err
	}
	if circle == nil {
		return false, nil
	}
	return true, nil
}
