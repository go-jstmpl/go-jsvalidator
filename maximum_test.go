package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

type IntMaximumValidatorTestCase struct {
	Input    int
	Expected error
}

func TestIntMaximumValidator(t *testing.T) {
	// Case exclusive is false
	definition := validator.IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: false,
	}

	va, err := validator.NewIntMaximumValidator(definition)
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
		Expected: &validator.IntMaximumValidationError{
			Input:      101,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := va.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = validator.IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: true,
	}

	va, err = validator.NewIntMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []IntMaximumValidatorTestCase{{
		Input:    99,
		Expected: nil,
	}, {
		Input: 100,
		Expected: &validator.IntMaximumValidationError{
			Input:      100,
			Definition: definition,
		},
	}, {
		Input: 101,
		Expected: &validator.IntMaximumValidationError{
			Input:      101,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := va.Validate(test.Input)
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
	definition := validator.FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	va, err := validator.NewFloatMaximumValidator(definition)
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
		Expected: &validator.FloatMaximumValidationError{
			Input:      1.1,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := va.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = validator.FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	va, err = validator.NewFloatMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []FloatMaximumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input: 1.0,
		Expected: &validator.FloatMaximumValidationError{
			Input:      1.0,
			Definition: definition,
		},
	}, {
		Input: 1.1,
		Expected: &validator.FloatMaximumValidationError{
			Input:      1.1,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := va.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
