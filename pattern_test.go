package validator

import (
	"reflect"
	"testing"
)

type PatternTestCase struct {
	Definition PatternValidatorDefinition
	Expected   error
}

func TestPattern(t *testing.T) {
	tests := []PatternTestCase{{
		Definition: PatternValidatorDefinition{Pattern: ""},
		Expected:   EmptyError{},
	}, {
		Definition: PatternValidatorDefinition{Pattern: "[a-z"},
		Expected:   InvalidPatternError{},
	}}
	for _, test := range tests {
		_, err := NewPatternValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type PatternValidatorTestCase struct {
	Input    string
	Expected error
}

func TestPatternValidator(t *testing.T) {
	definition := PatternValidatorDefinition{
		Pattern: "^\\d{7}$",
	}
	validator, err := NewPatternValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []PatternValidatorTestCase{{
		Input:    "1234567",
		Expected: nil,
	}, {
		Input: "abcdefg",
		Expected: &PatternValidationError{
			Input:      "abcdefg",
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
