package validator

import (
	"reflect"
	"testing"
)

type MaxLengthTestCase struct {
	Definition MaxLengthValidatorDefinition
	Expected   error
}

func TestMaxLength(t *testing.T) {
	tests := []MaxLengthTestCase{{
		Definition: MaxLengthValidatorDefinition{MaxLength: -1},
		Expected:   NoLengthError{},
	}}

	for _, test := range tests {
		_, err := NewMaxLengthValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type MaxLengthValidatorTestCase struct {
	Input    string
	Expected error
}

func TestMaxLengthValidator(t *testing.T) {
	definition := MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	validator, err := NewMaxLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []MaxLengthValidatorTestCase{{
		Input:    "あいうえ",
		Expected: nil,
	}, {
		Input:    "あいうえお",
		Expected: nil,
	}, {
		Input: "あいうえおか",
		Expected: &MaxLengthValidationError{
			Input:      "あいうえおか",
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
