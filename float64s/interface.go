package float64s

type Validator interface {
	Validate(input float64) error
}
