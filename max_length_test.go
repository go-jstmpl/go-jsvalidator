package validator

import (
	"reflect"
	"testing"
)

func TestMaxLength(t *testing.T) {
	type MaxLengthTestCase struct {
		Definition MaxLengthValidatorDefinition
		Expected   error
	}

	cases := []MaxLengthTestCase{
		{
			Definition: MaxLengthValidatorDefinition{MaxLength: -1},
			Expected:   NoLengthError{},
		},
	}

	for _, c := range cases {
		_, err := NewMaxLengthValidator(c.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(c.Expected) {
			t.Errorf("expected %v, but actual %v", reflect.TypeOf(c.Expected), reflect.TypeOf(err))
		}
	}
}

func TestMaxLengthValidator(t *testing.T) {
	definition := MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	validator, err := NewMaxLengthValidator(definition)
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
			Expected: &MaxLengthValidationError{
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
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
