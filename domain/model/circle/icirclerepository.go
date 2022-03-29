package circle

type ICircleRepository interface {
	FindByCircleName(circleName *CircleName) (*Circle, error)
}
