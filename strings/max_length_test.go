package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestNewMaxLengthValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition strings.MaxLengthValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "negative numver",
			Definition: strings.MaxLengthValidatorDefinition{MaxLength: -1},
			Error:      strings.MaxLengthDefinitionNoLengthError,
		},
		{
			Message:    "zero",
			Definition: strings.MaxLengthValidatorDefinition{MaxLength: 0},
			Error:      nil,
		},
		{
			Message:    "positive numver",
			Definition: strings.MaxLengthValidatorDefinition{MaxLength: 1},
			Error:      nil,
		},
	}
	for _, c := range cases {
		_, err := strings.NewMaxLengthValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewMaxLengthValidator with error %v", c.Message, err)
		}
	}
}

func TestValidateOfMaxLengthValidator(t *testing.T) {
	definition := strings.MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	va, err := strings.NewMaxLengthValidator(definition)
	if err != nil {
		t.Error(&err)
	}

	type MaxLengthValidatorTestCase struct {
		Input    string
		Expected error
	}

	cases := []MaxLengthValidatorTestCase{
		{
			Input:    "あいうえ",
			Expected: nil,
		},
		{
			Input:    "あいうえお",
			Expected: nil,
		},
		{
			Input: "あいうえおか",
			Expected: &strings.MaxLengthValidationError{
				Input:      "あいうえおか",
				Definition: definition,
			},
		},
		{
			Input:    "abcde",
			Expected: nil,
		},
	}

	for _, c := range cases {
		if err := va.Validate(c.Input); !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
