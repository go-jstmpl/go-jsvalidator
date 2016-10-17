package strings

type Validator interface {
	Validate(input string) error
}
