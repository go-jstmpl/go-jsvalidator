package bools

type Validator interface {
	Validate(input bool) error
}
