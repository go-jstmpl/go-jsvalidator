package strings

import (
	"fmt"
	"regexp"
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
	return fmt.Sprintf("input value '%s' does not match the regex pattern '%s'\n", p.Input, p.Definition.Pattern)
}

func NewPatternValidator(definition PatternValidatorDefinition) (PatternValidator, error) {
	if definition.Pattern == "" {
		return PatternValidator{}, EmptyError{"the pattern should not be empty"}
	}
	_, err := regexp.Compile(definition.Pattern)
	if err != nil {
		return PatternValidator{}, InvalidPatternError{fmt.Sprintf("invalid pattern: %s", definition.Pattern)}
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
