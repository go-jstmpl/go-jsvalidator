package strings

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	PatternDefinitionEmptyError          = errors.New("the pattern should not be empty")
	PatternDefinitionInvalidPatternError = errors.New("the pattern should not be invalid")
)

type PatternValidator struct {
	definition PatternValidatorDefinition
}

type PatternValidatorDefinition struct {
	Pattern string `json:"pattern"`
}

type PatternValidationError struct {
	Definition PatternValidatorDefinition `json:"definition"`
	Input      string                     `json:"input"`
}

func (p PatternValidationError) Error() string {
	return fmt.Sprintf("input value '%s' does not match the regex pattern '%s'", p.Input, p.Definition.Pattern)
}

func NewPatternValidator(definition PatternValidatorDefinition) (PatternValidator, error) {
	if definition.Pattern == "" {
		return PatternValidator{}, PatternDefinitionEmptyError
	}
	_, err := regexp.Compile(definition.Pattern)
	if err != nil {
		return PatternValidator{}, PatternDefinitionInvalidPatternError
	}
	return PatternValidator{definition}, nil
}

func (p PatternValidator) Validate(input string) error {
	ok, err := regexp.MatchString(p.definition.Pattern, input)
	if err == nil && ok == true {
		return nil
	}
	return &PatternValidationError{
		p.definition,
		input,
	}
}
