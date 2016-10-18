package bools

import (
	"errors"
	"fmt"
)

var (
	EnumDefinitionEmptyError       = errors.New("the Enumerate should have at least one element")
	EnumDefinitionDuplicationError = errors.New("the elements of Enumerate shouldn't be duplicated")
)

type EnumValidator struct {
	definition EnumValidatorDefinition
}

type EnumValidatorDefinition struct {
	Enumerate []bool `json:"enumerate"`
}

type EnumValidationError struct {
	Definition EnumValidatorDefinition `json:"definition"`
	Input      bool                    `json:"input"`
}

func (err EnumValidationError) Error() string {
	return fmt.Sprintf("input value %b doesn't exist in %v", err.Input, err.Definition.Enumerate)
}

func NewEnumValidator(def EnumValidatorDefinition) (EnumValidator, error) {
	len := len(def.Enumerate)
	if len == 0 {
		return EnumValidator{}, EnumDefinitionEmptyError
	}

	for i := 0; i < len-1; i++ {
		e := def.Enumerate[i]
		for j := i + 1; j < len; j++ {
			if def.Enumerate[j] == e {
				return EnumValidator{}, EnumDefinitionDuplicationError
			}
		}
	}

	return EnumValidator{def}, nil
}

func (v EnumValidator) Validate(input bool) error {
	for _, e := range v.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &EnumValidationError{
		v.definition,
		input,
	}
}
