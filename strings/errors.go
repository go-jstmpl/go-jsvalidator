package strings

// InvalidPatternError for Constructor methods
type InvalidPatternError struct {
	message string
}

func (e InvalidPatternError) Error() string {
	return e.message
}
