package validator

import "fmt"

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

// InvalidFieldTypeError for Required Validate method
type InvalidFieldTypeError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func (e InvalidFieldTypeError) Error() string {
	return fmt.Sprintf("input struct have invalid field against required '%v'", e.Definition.Required)
}
