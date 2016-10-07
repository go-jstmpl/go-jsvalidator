package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestPattern(t *testing.T) {
	_, err := validator.NewPatternValidator(validator.PatternValidatorDefinition{Pattern: ""})
	_, ok := err.(validator.EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = validator.NewPatternValidator(validator.PatternValidatorDefinition{Pattern: "[a-z"})
	_, ok = err.(validator.InvalidPatternError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "InvalidPatternError")
	}
}

func TestPatternValidator(t *testing.T) {
	definition := validator.PatternValidatorDefinition{
		Pattern: "^\\d{7}$",
	}
	va, err := validator.NewPatternValidator(definition)
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
			Expected: &validator.PatternValidationError{
				Input:      "abcdefg",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
