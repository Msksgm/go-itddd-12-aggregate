package circle

type CircleId struct {
	value string
}

func NewCircleId(id string) (*CircleId, error) {
	return &CircleId{value: id}, nil
}
