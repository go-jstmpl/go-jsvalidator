package slices_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/slices"
)

func TestNewMinItemsValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition slices.MinItemsValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "positive length",
			Definition: slices.MinItemsValidatorDefinition{MinItems: 1},
			Error:      nil,
		},
		{
			Message:    "zero length",
			Definition: slices.MinItemsValidatorDefinition{MinItems: 0},
			Error:      nil,
		},
		{
			Message:    "negative length",
			Definition: slices.MinItemsValidatorDefinition{MinItems: -1},
			Error:      &slices.NoLengthError{},
		},
	}

	for _, c := range cases {
		if _, err := slices.NewMinItemsValidator(c.Definition); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("%s: Error is expected '%v', but actual '%v'", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinItemsValidator(t *testing.T) {
	def := slices.MinItemsValidatorDefinition{
		MinItems: 2,
	}
	v, err := slices.NewMinItemsValidator(def)
	if err != nil {
		t.Fatalf("Fail to construct: %s", err)
	}

	type Case struct {
		Message string
		Input   interface{}
		Error   error
	}
	type Foo struct{}
	cases := []Case{
		{
			Message: "zero length of int slice",
			Input:   []int{},
			Error: &slices.MinItemsValidationError{
				Input:      []int{},
				Definition: def,
			},
		},
		{
			Message: "less length of int slice",
			Input:   []int{1},
			Error: &slices.MinItemsValidationError{
				Input:      []int{1},
				Definition: def,
			},
		},
		{
			Message: "same length of int slice",
			Input:   []int{1, 2},
			Error:   nil,
		},
		{
			Message: "greater length of int slice",
			Input:   []int{1, 2, 3},
			Error:   nil,
		},
		{
			Message: "zero length of string slice",
			Input:   []string{},
			Error: &slices.MinItemsValidationError{
				Input:      []string{},
				Definition: def,
			},
		},
		{
			Message: "less length of string slice",
			Input:   []string{"foo"},
			Error: &slices.MinItemsValidationError{
				Input:      []string{"foo"},
				Definition: def,
			},
		},
		{
			Message: "same length of string slice",
			Input:   []string{"foo", "bar"},
			Error:   nil,
		},
		{
			Message: "greater length of string slice",
			Input:   []string{"foo", "bar", "baz"},
			Error:   nil,
		},
		{
			Message: "zero length of float64 slice",
			Input:   []float64{},
			Error: &slices.MinItemsValidationError{
				Input:      []float64{},
				Definition: def,
			},
		},
		{
			Message: "less length of float64 slice",
			Input:   []float64{1},
			Error: &slices.MinItemsValidationError{
				Input:      []float64{1},
				Definition: def,
			},
		},
		{
			Message: "same length of float64 slice",
			Input:   []float64{1, 2},
			Error:   nil,
		},
		{
			Message: "greater length of float64 slice",
			Input:   []float64{1, 2, 3},
			Error:   nil,
		},
		{
			Message: "zero length of struct slice",
			Input:   []Foo{},
			Error: &slices.MinItemsValidationError{
				Input:      []Foo{},
				Definition: def,
			},
		},
		{
			Message: "less length of struct slice",
			Input:   []Foo{{}},
			Error: &slices.MinItemsValidationError{
				Input:      []Foo{{}},
				Definition: def,
			},
		},
		{
			Message: "same length of struct slice",
			Input:   []Foo{{}, {}},
			Error:   nil,
		},
		{
			Message: "greater length of struct slice",
			Input:   []Foo{{}, {}, {}},
			Error:   nil,
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("%s: expected %+v, but actual %+v", c.Message, c.Error, err)
		}
	}
}
