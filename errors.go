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
