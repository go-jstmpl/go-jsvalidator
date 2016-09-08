package validator

import (
	"reflect"
	"testing"
)

type MinLengthTestCase struct {
	Definition MinLengthValidatorDefinition
	Expected   error
}

func TestMinLength(t *testing.T) {
	tests := []MinLengthTestCase{{
		Definition: MinLengthValidatorDefinition{MinLength: -1},
		Expected:   NoLengthError{},
	}}

	for _, test := range tests {
		_, err := NewMinLengthValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type MinLengthValidatorTestCase struct {
	Input    string
	Expected error
}

func TestMinLengthValidator(t *testing.T) {
	definition := MinLengthValidatorDefinition{
		MinLength: 5,
	}
	validator, err := NewMinLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []MinLengthValidatorTestCase{{
		Input:    "あいうえおか",
		Expected: nil,
	}, {
		Input:    "あいうえお",
		Expected: nil,
	}, {
		Input: "あいうえ",
		Expected: &MinLengthValidationError{
			Input:      "あいうえ",
			Definition: definition,
		},
	}, {
		Input:    "abcde",
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
