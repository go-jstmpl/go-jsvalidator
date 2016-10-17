package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestPattern(t *testing.T) {
	_, err := strings.NewPatternValidator(strings.PatternValidatorDefinition{Pattern: ""})
	_, ok := err.(strings.EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = strings.NewPatternValidator(strings.PatternValidatorDefinition{Pattern: "[a-z"})
	_, ok = err.(strings.InvalidPatternError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "InvalidPatternError")
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
