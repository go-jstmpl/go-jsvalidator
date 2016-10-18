package slices

type NoLengthError struct{}

func (e NoLengthError) Error() string {
	return "the value of MaxItems should be greater than or equal to 0"
}

type TypeError struct {
	Message string
}

func (e TypeError) Error() string {
	return e.Message
}
