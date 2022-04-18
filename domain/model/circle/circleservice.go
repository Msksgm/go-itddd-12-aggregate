package circle

type CircleService struct {
	circleRepository CircleRepositorier
}

func (circleService *CircleService) NewCircleService(circleRepository CircleRepositorier) error {
	circleService.circleRepository = circleRepository
	return nil
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
