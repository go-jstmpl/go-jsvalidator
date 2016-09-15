package validator

import (
	"reflect"
	"testing"
)

type MaximumValidatorTestCase struct {
	Input    interface{}
	Expected error
}

func TestMaximumValidator(t *testing.T) {
	// Case type int and exclusive is false
	definition := MaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: false,
	}

	validator, err := NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []MaximumValidatorTestCase{
		{
			Input:    99,
			Expected: nil,
		},
		{
			Input:    100,
			Expected: nil,
		},
		{
			Input: 101,
			Expected: &MaximumValidationError{
				Input:      101,
				Definition: definition,
			},
		},
		{
			Input: 10.1,
			Expected: TypeError{
				message: "input and maximum should be same type",
			},
		},
	}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case type int and exclusive is true
	definition = MaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: true,
	}

	validator, err = NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []MaximumValidatorTestCase{
		{
			Input:    99,
			Expected: nil,
		},
		{
			Input: 100,
			Expected: &MaximumValidationError{
				Input:      100,
				Definition: definition,
			},
		},
		{
			Input: 101,
			Expected: &MaximumValidationError{
				Input:      101,
				Definition: definition,
			},
		},
	}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case type float64 and exclusive is false
	definition = MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	validator, err = NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []MaximumValidatorTestCase{
		{
			Input:    0.9,
			Expected: nil,
		},
		{
			Input:    1.0,
			Expected: nil,
		},
		{
			Input: 1.1,
			Expected: &MaximumValidationError{
				Input:      1.1,
				Definition: definition,
			},
		},
		{
			Input: 10,
			Expected: TypeError{
				message: "input and maximum should be same type",
			},
		},
	}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case type float64 and exclusive is true
	definition = MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	validator, err = NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []MaximumValidatorTestCase{
		{
			Input:    0.9,
			Expected: nil,
		},
		{
			Input: 1.0,
			Expected: &MaximumValidationError{
				Input:      1.0,
				Definition: definition,
			},
		},
		{
			Input: 1.1,
			Expected: &MaximumValidationError{
				Input:      1.1,
				Definition: definition,
			},
		},
	}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
