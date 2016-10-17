package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestMaximumValidator(t *testing.T) {
	// Case type int and exclusive is false
	definition := validator.MaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: false,
	}

	va, err := validator.NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MaximumValidatorTestCase struct {
		Input    interface{}
		Expected error
	}
	cases := []MaximumValidatorTestCase{
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
			Expected: &validator.MaximumValidationError{
				Input:      101,
				Definition: definition,
			},
		},
		{
			Input: 10.1,
			Expected: validator.TypeError{
				Message: "input and maximum should be same type",
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	// Case type int and exclusive is true
	definition = validator.MaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: true,
	}

	va, err = validator.NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MaximumValidatorTestCase{
		{
			Input:    99,
			Expected: nil,
		},
		{
			Input: 100,
			Expected: &validator.MaximumValidationError{
				Input:      100,
				Definition: definition,
			},
		},
		{
			Input: 101,
			Expected: &validator.MaximumValidationError{
				Input:      101,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	// Case type float64 and exclusive is false
	definition = validator.MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	va, err = validator.NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MaximumValidatorTestCase{
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
			Expected: &validator.MaximumValidationError{
				Input:      1.1,
				Definition: definition,
			},
		},
		{
			Input: 10,
			Expected: validator.TypeError{
				Message: "input and maximum should be same type",
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	// Case type float64 and exclusive is true
	definition = validator.MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	va, err = validator.NewMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MaximumValidatorTestCase{
		{
			Input:    0.9,
			Expected: nil,
		},
		{
			Input: 1.0,
			Expected: &validator.MaximumValidationError{
				Input:      1.0,
				Definition: definition,
			},
		},
		{
			Input: 1.1,
			Expected: &validator.MaximumValidationError{
				Input:      1.1,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
