package circle

type CircleRepositorier interface {
	FindByCircleName(circleName *CircleName) (*Circle, error)
}
