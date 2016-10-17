package validator_test

import (
	"reflect"
	"testing"

	validator "github.com/go-jstmpl/go-jsvalidator"
)

type IntMinimumValidatorTestCase struct {
	Input    int
	Expected error
}

func TestIntMinimumValidator(t *testing.T) {
	// Case exclusive is false
	definition := validator.IntMinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: false,
	}

	va, err := validator.NewIntMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []IntMinimumValidatorTestCase{{
		Input:    101,
		Expected: nil,
	}, {
		Input:    100,
		Expected: nil,
	}, {
		Input: 99,
		Expected: &validator.IntMinimumValidationError{
			Input:      99,
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
	definition = validator.IntMinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: true,
	}

	va, err = validator.NewIntMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []IntMinimumValidatorTestCase{{
		Input:    101,
		Expected: nil,
	}, {
		Input: 100,
		Expected: &validator.IntMinimumValidationError{
			Input:      100,
			Definition: definition,
		},
	}, {
		Input: 99,
		Expected: &validator.IntMinimumValidationError{
			Input:      99,
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

type FloatMinimumValidatorTestCase struct {
	Input    float64
	Expected error
}

func TestFloatMinimumValidator(t *testing.T) {
	// Case exclusive is false
	definition := validator.FloatMinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: false,
	}

	va, err := validator.NewFloatMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []FloatMinimumValidatorTestCase{{
		Input:    1.1,
		Expected: nil,
	}, {
		Input:    1.0,
		Expected: nil,
	}, {
		Input: 0.9,
		Expected: &validator.FloatMinimumValidationError{
			Input:      0.9,
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
	definition = validator.FloatMinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: true,
	}

	va, err = validator.NewFloatMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []FloatMinimumValidatorTestCase{{
		Input:    1.1,
		Expected: nil,
	}, {
		Input: 1.0,
		Expected: &validator.FloatMinimumValidationError{
			Input:      1.0,
			Definition: definition,
		},
	}, {
		Input: 0.9,
		Expected: &validator.FloatMinimumValidationError{
			Input:      0.9,
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
