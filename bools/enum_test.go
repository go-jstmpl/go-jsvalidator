package bools_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/bools"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition bools.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: bools.EnumValidatorDefinition{Enumerate: []bool{}},
			Error:      bools.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: bools.EnumValidatorDefinition{Enumerate: []bool{true}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: bools.EnumValidatorDefinition{Enumerate: []bool{true, false}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: bools.EnumValidatorDefinition{Enumerate: []bool{true, false, true}},
			Error:      bools.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := bools.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected '%s', but actual '%s'", c.Message, c.Error, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := bools.EnumValidatorDefinition{
		Enumerate: []bool{true},
	}
	v, err := bools.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator: %s", err)
	}

	type Case struct {
		Message string
		Input   bool
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists in Enumerate",
			Input:   true,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enumerate",
			Input:   false,
			Error: &bools.EnumValidationError{
				Input:      false,
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
