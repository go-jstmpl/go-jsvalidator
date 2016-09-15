package validator

import "testing"

type IntMinItemsTestCase struct {
	Definition MinItemsValidatorDefinition
	Expected   error
}

func TestIntMinItems(t *testing.T) {
	tests := []IntMinItemsTestCase{{
		Definition: MinItemsValidatorDefinition{MinItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewMinItemsValidator(test.Definition)
		if err == nil {
			t.Errorf("expected:%v, actual:%v", test.Expected, err)
		}
	}
}

type MinItemsValidatorTestCase struct {
	Input    interface{}
	Expected error
}

func TestMinItemsValidator(t *testing.T) {
	definition := MinItemsValidatorDefinition{
		MinItems: 3,
	}
	validator, err := NewMinItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []MinItemsValidatorTestCase{{
		Input: []int{1},
		Expected: &MinItemsValidationError{
			Input:      []int{1},
			Definition: definition,
		},
	}, {
		Input: []int{1, 2},
		Expected: &MinItemsValidationError{
			Input:      []int{1, 2},
			Definition: definition,
		},
	}, {
		Input:    []int{1, 2, 3},
		Expected: nil,
	}, {
		Input:    []int{1, 2, 3, 4},
		Expected: nil,
	}, {
		Input: []string{"foo"},
		Expected: &MinItemsValidationError{
			Input: []string{"foo"},

			Definition: definition,
		},
	}, {
		Input:    []string{"foo", "bar", "baz"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar", "bas", "qux"},
		Expected: nil,
	}, {
		Input: []float64{1},
		Expected: &MinItemsValidationError{
			Input:      []float64{1},
			Definition: definition,
		},
	}, {
		Input:    []float64{1, 2, 3},
		Expected: nil,
	}, {
		Input:    []float64{1, 2, 3, 4},
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if test.Expected == nil && err != nil {
			t.Errorf("expected: %+v ,actual: %+v", test.Expected, err)
		}

		if test.Expected != nil && err == nil {
			t.Errorf("expected: %+v ,actual: %+v", test.Expected, err)
		}
	}
}
