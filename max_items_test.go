package validator

import "testing"

type MaxItemsTestCase struct {
	Definition MaxItemsValidatorDefinition
	Expected   error
}

func TestMaxItems(t *testing.T) {
	tests := []MaxItemsTestCase{{
		Definition: MaxItemsValidatorDefinition{MaxItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewMaxItemsValidator(test.Definition)
		if err == nil {
			t.Errorf("expected:%v, actual:%v", test.Expected, err)
		}
	}
}

type MaxItemsValidatorTestCase struct {
	Input    interface{}
	Expected error
}

func TestMaxItemsValidator(t *testing.T) {
	definition := MaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	validator, err := NewMaxItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []MaxItemsValidatorTestCase{{
		Input:    []int{1},
		Expected: nil,
	}, {
		Input:    []int{1, 2},
		Expected: nil,
	}, {
		Input:    []int{1, 2, 3},
		Expected: nil,
	}, {
		Input: []int{1, 2, 3, 4},
		Expected: &MaxItemsValidationError{
			Input:      []int{1, 2, 3, 4},
			Definition: definition,
		},
	}, {
		Input:    []string{"foo"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar", "baz"},
		Expected: nil,
	}, {
		Input: []string{"foo", "bar", "bas", "qux"},
		Expected: &MaxItemsValidationError{
			Input:      []string{"foo", "bar", "bas", "qux"},
			Definition: definition,
		},
	}, {
		Input:    []float64{1},
		Expected: nil,
	}, {
		Input:    []float64{1, 2},
		Expected: nil,
	}, {
		Input:    []float64{1, 2, 3},
		Expected: nil,
	}, {
		Input: []float64{1, 2, 3, 4},
		Expected: &MaxItemsValidationError{
			Input:      []float64{1, 2, 3, 4},
			Definition: definition,
		},
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
