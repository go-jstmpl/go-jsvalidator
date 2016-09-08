package validator

import (
	"reflect"
	"testing"
)

type IntMaximumValidatorTestCase struct {
	Input    int
	Expected error
}

func TestIntMaximumValidator(t *testing.T) {
	// Case exclusive is false
	definition := IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: false,
	}

	validator, err := NewIntMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []IntMaximumValidatorTestCase{{
		Input:    99,
		Expected: nil,
	}, {
		Input:    100,
		Expected: nil,
	}, {
		Input: 101,
		Expected: &IntMaximumValidationError{
			Input:      101,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: true,
	}

	validator, err = NewIntMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []IntMaximumValidatorTestCase{{
		Input:    99,
		Expected: nil,
	}, {
		Input: 100,
		Expected: &IntMaximumValidationError{
			Input:      100,
			Definition: definition,
		},
	}, {
		Input: 101,
		Expected: &IntMaximumValidationError{
			Input:      101,
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

type FloatMaximumValidatorTestCase struct {
	Input    float64
	Expected error
}

func TestFloatMaximumValidator(t *testing.T) {
	// Case exclusive is false
	definition := FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	validator, err := NewFloatMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []FloatMaximumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input:    1.0,
		Expected: nil,
	}, {
		Input: 1.1,
		Expected: &FloatMaximumValidationError{
			Input:      1.1,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	validator, err = NewFloatMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []FloatMaximumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input: 1.0,
		Expected: &FloatMaximumValidationError{
			Input:      1.0,
			Definition: definition,
		},
	}, {
		Input: 1.1,
		Expected: &FloatMaximumValidationError{
			Input:      1.1,
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
