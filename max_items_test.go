package validator

import (
	"reflect"
	"testing"
)

func TestMaxItems(t *testing.T) {
	type MaxItemsTestCase struct {
		Definition MaxItemsValidatorDefinition
		Expected   error
	}
	cases := []MaxItemsTestCase{{
		Definition: MaxItemsValidatorDefinition{MaxItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, c := range cases {
		_, err := NewMaxItemsValidator(c.Definition)
		if err == nil {
			t.Errorf("expected %+v, but actual %+v", c.Expected, err)
		}
	}
}

func TestMaxItemsValidator(t *testing.T) {
	definition := MaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	validator, err := NewMaxItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MaxItemsValidatorTestCase struct {
		Input    interface{}
		Expected error
	}
	cases := []MaxItemsValidatorTestCase{
		{
			Input:    []int{1},
			Expected: nil,
		},
		{
			Input:    []int{1, 2},
			Expected: nil,
		},
		{
			Input:    []int{1, 2, 3},
			Expected: nil,
		},
		{
			Input:    []string{"foo"},
			Expected: nil,
		},
		{
			Input:    []string{"foo", "bar"},
			Expected: nil,
		},
		{
			Input:    []string{"foo", "bar", "baz"},
			Expected: nil,
		},

		{
			Input:    []float64{1},
			Expected: nil,
		},
		{
			Input:    []float64{1, 2},
			Expected: nil,
		},
		{
			Input:    []float64{1, 2, 3},
			Expected: nil,
		},
		{
			Input: []int{1, 2, 3, 4},
			Expected: &MaxItemsValidationError{
				Input:      []int{1, 2, 3, 4},
				Definition: definition,
			},
		},
		{
			Input: []string{"foo", "bar", "bas", "qux"},
			Expected: &MaxItemsValidationError{
				Input:      []string{"foo", "bar", "bas", "qux"},
				Definition: definition,
			},
		},
		{
			Input: []float64{1, 2, 3, 4},
			Expected: &MaxItemsValidationError{
				Input:      []float64{1, 2, 3, 4},
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %+v, but actual %+v", c.Expected, err)
		}
	}
}
