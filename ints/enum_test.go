package ints_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/ints"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition ints.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: ints.EnumValidatorDefinition{Enumerate: []int{}},
			Error:      ints.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: ints.EnumValidatorDefinition{Enumerate: []int{-10}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: ints.EnumValidatorDefinition{Enumerate: []int{-10, 0, 10}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: ints.EnumValidatorDefinition{Enumerate: []int{-10, 0, -10}},
			Error:      ints.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := ints.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: fail to NewEnumValidator with error %v", c.Message, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := ints.EnumValidatorDefinition{
		Enumerate: []int{-10, 0, 10},
	}
	v, err := ints.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator with error %v", err)
	}

	type Case struct {
		Message string
		Input   int
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists at first in Enumerate",
			Input:   -10,
			Error:   nil,
		},
		{
			Message: "a value exists at second in Enumerate",
			Input:   0,
			Error:   nil,
		},
		{
			Message: "a value exists at end in Enumerate",
			Input:   10,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enumerate",
			Input:   -20,
			Error: &ints.EnumValidationError{
				Input:      -20,
				Definition: def,
			},
		},
	}
	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}
