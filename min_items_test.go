package validator

import (
	"reflect"
	"testing"
)

func TestMinItems(t *testing.T) {
	_, err := NewMinItemsValidator(MinItemsValidatorDefinition{MinItems: -1})
	_, ok := err.(NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMinItemsValidator(t *testing.T) {
	definition := MinItemsValidatorDefinition{
		MinItems: 3,
	}
	validator, err := NewMinItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MinItemsValidatorTestCase struct {
		Input    interface{}
		Expected error
	}
	cases := []MinItemsValidatorTestCase{
		{
			Input: []int{1},
			Expected: &MinItemsValidationError{
				Input:      []int{1},
				Definition: definition,
			},
		},
		{
			Input: []int{1, 2},
			Expected: &MinItemsValidationError{
				Input:      []int{1, 2},
				Definition: definition,
			},
		},
		{
			Input:    []int{1, 2, 3},
			Expected: nil,
		},
		{
			Input:    []int{1, 2, 3, 4},
			Expected: nil,
		},
		{
			Input: []string{"foo"},
			Expected: &MinItemsValidationError{
				Input: []string{"foo"},

				Definition: definition,
			},
		},
		{
			Input:    []string{"foo", "bar", "baz"},
			Expected: nil,
		},
		{
			Input:    []string{"foo", "bar", "bas", "qux"},
			Expected: nil,
		},
		{
			Input: []float64{1},
			Expected: &MinItemsValidationError{
				Input:      []float64{1},
				Definition: definition,
			},
		},
		{
			Input:    []float64{1, 2, 3},
			Expected: nil,
		},
		{
			Input:    []float64{1, 2, 3, 4},
			Expected: nil,
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
