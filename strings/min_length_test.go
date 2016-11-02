package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestNewMinLengthValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition strings.MinLengthValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "negative numver",
			Definition: strings.MinLengthValidatorDefinition{MinLength: -1},
			Error:      strings.MinLengthDefinitionNoLengthError,
		},
		{
			Message:    "zero",
			Definition: strings.MinLengthValidatorDefinition{MinLength: 0},
			Error:      nil,
		},
		{
			Message:    "positive numver",
			Definition: strings.MinLengthValidatorDefinition{MinLength: 1},
			Error:      nil,
		},
	}
	for _, c := range cases {
		_, err := strings.NewMinLengthValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewMinLengthValidator with error %v", c.Message, err)
		}
	}
}

func TestMinLengthValidator(t *testing.T) {
	definition := strings.MinLengthValidatorDefinition{
		MinLength: 5,
	}
	va, err := strings.NewMinLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MinLengthValidatorTestCase struct {
		Input    string
		Expected error
	}
	cases := []MinLengthValidatorTestCase{
		{
			Input:    "あいうえおか",
			Expected: nil,
		},
		{
			Input:    "あいうえお",
			Expected: nil,
		},
		{
			Input: "あいうえ",
			Expected: &strings.MinLengthValidationError{
				Input:      "あいうえ",
				Definition: definition,
			},
		},
		{
			Input:    "abcde",
			Expected: nil,
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
