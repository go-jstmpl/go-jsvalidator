package validator

import (
	"reflect"
	"testing"
)

func TestMinimumValidator(t *testing.T) {
	// Case exclusive is false
	definition := MinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: false,
	}

	validator, err := NewMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MinimumValidatorTestCase struct {
		Input    interface{}
		Expected error
	}
	cases := []MinimumValidatorTestCase{
		{
			Input:    101,
			Expected: nil,
		},
		{
			Input:    100,
			Expected: nil,
		},
		{
			Input: 99,
			Expected: &MinimumValidationError{
				Input:      99,
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

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	// Case exclusive is true
	definition = MinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: true,
	}

	validator, err = NewMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MinimumValidatorTestCase{
		{
			Input:    101,
			Expected: nil,
		},
		{
			Input: 100,
			Expected: &MinimumValidationError{
				Input:      100,
				Definition: definition,
			},
		},
		{
			Input: 99,
			Expected: &MinimumValidationError{
				Input:      99,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, actual %v", c.Expected, err)
		}
	}

	// Case exclusive is false
	definition = MinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: false,
	}

	validator, err = NewMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MinimumValidatorTestCase{
		{
			Input:    1.1,
			Expected: nil,
		},
		{
			Input:    1.0,
			Expected: nil,
		},
		{
			Input: 0.9,
			Expected: &MinimumValidationError{
				Input:      0.9,
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

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	// Case exclusive is true
	definition = MinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: true,
	}

	validator, err = NewMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []MinimumValidatorTestCase{
		{
			Input:    1.1,
			Expected: nil,
		},
		{
			Input: 1.0,
			Expected: &MinimumValidationError{
				Input:      1.0,
				Definition: definition,
			},
		},
		{
			Input: 0.9,
			Expected: &MinimumValidationError{
				Input:      0.9,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
