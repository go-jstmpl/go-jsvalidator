package validator

import (
	"reflect"
	"testing"
)

func TestPattern(t *testing.T) {
	type PatternTestCase struct {
		Definition PatternValidatorDefinition
		Expected   error
	}
	cases := []PatternTestCase{
		{
			Definition: PatternValidatorDefinition{Pattern: ""},
			Expected:   EmptyError{},
		},
		{
			Definition: PatternValidatorDefinition{Pattern: "[a-z"},
			Expected:   InvalidPatternError{},
		},
	}
	for _, c := range cases {
		_, err := NewPatternValidator(c.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(c.Expected) {
			t.Errorf("expected %v, but actual %v", reflect.TypeOf(c.Expected), reflect.TypeOf(err))
		}
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
