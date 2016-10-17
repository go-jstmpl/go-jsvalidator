package slices

type Validator interface {
	Validate(input []interface{})
}
