package numbers

type Validator interface {
	Validate(input float64) error
}
