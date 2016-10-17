package slices_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/slices"
)

func TestNewMaxItemsValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition slices.MaxItemsValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "positive length",
			Definition: slices.MaxItemsValidatorDefinition{MaxItems: 1},
			Error:      nil,
		},
		{
			Message:    "zero length",
			Definition: slices.MaxItemsValidatorDefinition{MaxItems: 0},
			Error:      nil,
		},
		{
			Message:    "negative length",
			Definition: slices.MaxItemsValidatorDefinition{MaxItems: -1},
			Error:      &slices.NoLengthError{},
		},
	}

	for _, c := range cases {
		if _, err := slices.NewMaxItemsValidator(c.Definition); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("%s: Error is expected '%v', but actual '%v'", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMaxItemsValidator(t *testing.T) {
	def := slices.MaxItemsValidatorDefinition{
		MaxItems: 2,
	}
	v, err := slices.NewMaxItemsValidator(def)
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
			Error:   nil,
		},
		{
			Message: "less length of int slice",
			Input:   []int{1},
			Error:   nil,
		},
		{
			Message: "same length of int slice",
			Input:   []int{1, 2},
			Error:   nil,
		},
		{
			Message: "greater length of int slice",
			Input:   []int{1, 2, 3},
			Error: &slices.MaxItemsValidationError{
				Input:      []int{1, 2, 3},
				Definition: def,
			},
		},
		{
			Message: "zero length of string slice",
			Input:   []string{},
			Error:   nil,
		},
		{
			Message: "less length of string slice",
			Input:   []string{"foo"},
			Error:   nil,
		},
		{
			Message: "same length of string slice",
			Input:   []string{"foo", "bar"},
			Error:   nil,
		},
		{
			Message: "greater length of string slice",
			Input:   []string{"foo", "bar", "baz"},
			Error: &slices.MaxItemsValidationError{
				Input:      []string{"foo", "bar", "baz"},
				Definition: def,
			},
		},
		{
			Message: "zero length of float64 slice",
			Input:   []float64{},
			Error:   nil,
		},
		{
			Message: "less length of float64 slice",
			Input:   []float64{1},
			Error:   nil,
		},
		{
			Message: "same length of float64 slice",
			Input:   []float64{1, 2},
			Error:   nil,
		},
		{
			Message: "greater length of float64 slice",
			Input:   []float64{1, 2, 3},
			Error: &slices.MaxItemsValidationError{
				Input:      []float64{1, 2, 3},
				Definition: def,
			},
		},
		{
			Message: "zero length of struct slice",
			Input:   []Foo{},
			Error:   nil,
		},
		{
			Message: "less length of struct slice",
			Input:   []Foo{{}},
			Error:   nil,
		},
		{
			Message: "same length of struct slice",
			Input:   []Foo{{}, {}},
			Error:   nil,
		},
		{
			Message: "greater length of struct slice",
			Input:   []Foo{{}, {}, {}},
			Error: &slices.MaxItemsValidationError{
				Input:      []Foo{{}, {}, {}},
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("%s: expected %+v, but actual %+v", c.Message, c.Error, err)
		}
	}
}
