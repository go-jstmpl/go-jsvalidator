package arrays_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/arrays"
)

func TestNewMaxItemsValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition arrays.MaxItemsValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "positive length",
			Definition: arrays.MaxItemsValidatorDefinition{MaxItems: 1},
			Error:      nil,
		},
		{
			Message:    "zero length",
			Definition: arrays.MaxItemsValidatorDefinition{MaxItems: 0},
			Error:      nil,
		},
		{
			Message:    "negative length",
			Definition: arrays.MaxItemsValidatorDefinition{MaxItems: -1},
			Error:      arrays.MaxItemsDefinitionNoLengthError,
		},
	}

	for _, c := range cases {
		if _, err := arrays.NewMaxItemsValidator(c.Definition); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("%s: Error is expected '%v', but actual '%v'", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMaxItemsValidator(t *testing.T) {
	def := arrays.MaxItemsValidatorDefinition{
		MaxItems: 2,
	}
	v, err := arrays.NewMaxItemsValidator(def)
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
			Error: &arrays.MaxItemsValidationError{
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
			Error: &arrays.MaxItemsValidationError{
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
			Error: &arrays.MaxItemsValidationError{
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
			Error: &arrays.MaxItemsValidationError{
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
