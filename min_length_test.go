package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestMinLength(t *testing.T) {
	_, err := validator.NewMinLengthValidator(validator.MinLengthValidatorDefinition{MinLength: -1})
	_, ok := err.(validator.NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMinLengthValidator(t *testing.T) {
	definition := validator.MinLengthValidatorDefinition{
		MinLength: 5,
	}
	va, err := validator.NewMinLengthValidator(definition)
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
			Expected: &validator.MinLengthValidationError{
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
