package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestMaxLength(t *testing.T) {
	_, err := validator.NewMaxLengthValidator(validator.MaxLengthValidatorDefinition{MaxLength: -1})
	_, ok := err.(validator.NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMaxLengthValidator(t *testing.T) {
	definition := validator.MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	va, err := validator.NewMaxLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
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
			Expected: &validator.MaxLengthValidationError{
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
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
