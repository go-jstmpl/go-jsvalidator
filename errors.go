package validator

import "fmt"

// InvalidFieldTypeError for Required Validate method
type InvalidFieldTypeError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func (e InvalidFieldTypeError) Error() string {
	return fmt.Sprintf("input struct have invalid field against required '%v'", e.Definition.Required)
}

// InvalidFieldTypeError for Required Validate method
type InvalidTypeError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("the type of argument for the Validate of the RequiredValidator should be struct or pointer struct `%v`", e.Definition.Required)
}
