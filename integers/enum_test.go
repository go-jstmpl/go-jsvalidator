package integers_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/integers"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition integers.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: integers.EnumValidatorDefinition{Enum: []int{}},
			Error:      integers.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: integers.EnumValidatorDefinition{Enum: []int{-10}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: integers.EnumValidatorDefinition{Enum: []int{-10, 0, 10}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: integers.EnumValidatorDefinition{Enum: []int{-10, 0, -10}},
			Error:      integers.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := integers.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: fail to NewEnumValidator with error %v", c.Message, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := integers.EnumValidatorDefinition{
		Enum: []int{-10, 0, 10},
	}
	v, err := integers.NewEnumValidator(def)
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
			Message: "a value exists at first in Enum",
			Input:   -10,
			Error:   nil,
		},
		{
			Message: "a value exists at second in Enum",
			Input:   0,
			Error:   nil,
		},
		{
			Message: "a value exists at end in Enum",
			Input:   10,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enum",
			Input:   -20,
			Error: &integers.EnumValidationError{
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
