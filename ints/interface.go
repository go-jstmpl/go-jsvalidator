package ints

type Validator interface {
	Validate(input int) error
}
