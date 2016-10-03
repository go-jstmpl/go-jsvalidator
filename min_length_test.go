package validator

import (
	"reflect"
	"testing"
)

func TestMinLength(t *testing.T) {
	type MinLengthTestCase struct {
		Definition MinLengthValidatorDefinition
		Expected   error
	}
	cases := []MinLengthTestCase{
		{
			Definition: MinLengthValidatorDefinition{MinLength: -1},
			Expected:   NoLengthError{},
		},
	}

	for _, c := range cases {
		_, err := NewMinLengthValidator(c.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(c.Expected) {
			t.Errorf("expected %v, but actual %v", reflect.TypeOf(c.Expected), reflect.TypeOf(err))
		}
	}
}

func TestMinLengthValidator(t *testing.T) {
	definition := MinLengthValidatorDefinition{
		MinLength: 5,
	}
	validator, err := NewMinLengthValidator(definition)
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
			Expected: &MinLengthValidationError{
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
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
