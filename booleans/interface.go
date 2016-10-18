package booleans

type Validator interface {
	Validate(input bool) error
}
