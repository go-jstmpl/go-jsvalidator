package validator

import (
	"reflect"
	"testing"
)

func TestMinLength(t *testing.T) {
	_, err := NewMinLengthValidator(MinLengthValidatorDefinition{MinLength: -1})
	_, ok := err.(NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
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
