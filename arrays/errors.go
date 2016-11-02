package arrays

type TypeError struct {
	Message string
}

func (e TypeError) Error() string {
	return e.Message
}
