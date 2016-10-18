package strings

// EmptyError for Constructor methods
type EmptyError struct {
	message string
}

func (e EmptyError) Error() string {
	return e.message
}

// DuplicationError for Constructor methods
type DuplicationError struct {
	message string
}

func (e DuplicationError) Error() string {
	return e.message
}

// InvalidPatternError for Constructor methods
type InvalidPatternError struct {
	message string
}

func (e InvalidPatternError) Error() string {
	return e.message
}

// InvalidFormatError for Constructor methods
type InvalidFormatError struct {
	message string
}

func (e InvalidFormatError) Error() string {
	return e.message
}

type NoLengthError struct {
	message string
}

func (e NoLengthError) Error() string {
	return e.message
}
