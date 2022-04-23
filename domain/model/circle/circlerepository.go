package circle

type CircleRepositorier interface {
	Save(circle *Circle) error
	FindByCircleName(circleName *CircleName) (*Circle, error)
}
