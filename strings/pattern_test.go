package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestNewPatternValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition strings.PatternValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "valid pattern",
			Definition: strings.PatternValidatorDefinition{Pattern: "[a-z]"},
			Error:      nil,
		},
		{
			Message:    "empty string",
			Definition: strings.PatternValidatorDefinition{Pattern: ""},
			Error:      strings.PatternDefinitionEmptyError,
		},
	}
	for _, c := range cases {
		_, err := strings.NewPatternValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewPatternValidator with error %v", c.Message, err)
		}
	}
}

func TestNewPatternValidatorWithInvalidPattern(t *testing.T) {
	_, err := strings.NewPatternValidator(strings.PatternValidatorDefinition{
		Pattern: "[a-z",
	})
	if _, ok := err.(strings.InvalidPatternError); !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}
}

func TestPatternValidator(t *testing.T) {
	definition := strings.PatternValidatorDefinition{
		Pattern: "^\\d{7}$",
	}
	va, err := strings.NewPatternValidator(definition)
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
			Expected: &strings.PatternValidationError{
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
