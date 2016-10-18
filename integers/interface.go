package integers

type Validator interface {
	Validate(input int) error
}
