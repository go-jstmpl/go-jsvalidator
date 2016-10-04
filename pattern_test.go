package validator

import (
	"reflect"
	"testing"
)

func TestPattern(t *testing.T) {
	_, err := NewPatternValidator(PatternValidatorDefinition{Pattern: ""})
	_, ok := err.(EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = NewPatternValidator(PatternValidatorDefinition{Pattern: "[a-z"})
	_, ok = err.(InvalidPatternError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "InvalidPatternError")
	}
}

func TestPatternValidator(t *testing.T) {
	definition := PatternValidatorDefinition{
		Pattern: "^\\d{7}$",
	}
	validator, err := NewPatternValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type PatternValidatorTestCase struct {
		Input    string
		Expected error
	}
	cases := []PatternValidatorTestCase{
		{
			Input:    "1234567",
			Expected: nil,
		},
		{
			Input: "abcdefg",
			Expected: &PatternValidationError{
				Input:      "abcdefg",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
