package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestMaxItems(t *testing.T) {
	_, err := validator.NewMaxItemsValidator(validator.MaxItemsValidatorDefinition{MaxItems: -1})
	_, ok := err.(validator.NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMaxItemsValidator(t *testing.T) {
	definition := validator.MaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	va, err := validator.NewMaxItemsValidator(definition)
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
			Expected: &validator.MaxItemsValidationError{
				Input:      []int{1, 2, 3, 4},
				Definition: definition,
			},
		},
		{
			Input: []string{"foo", "bar", "bas", "qux"},
			Expected: &validator.MaxItemsValidationError{
				Input:      []string{"foo", "bar", "bas", "qux"},
				Definition: definition,
			},
		},
		{
			Input: []float64{1, 2, 3, 4},
			Expected: &validator.MaxItemsValidationError{
				Input:      []float64{1, 2, 3, 4},
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %+v, but actual %+v", c.Expected, err)
		}
	}
}
