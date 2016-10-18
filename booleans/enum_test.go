package booleans_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/booleans"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition booleans.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: booleans.EnumValidatorDefinition{Enum: []bool{}},
			Error:      booleans.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: booleans.EnumValidatorDefinition{Enum: []bool{true}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: booleans.EnumValidatorDefinition{Enum: []bool{true, false}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: booleans.EnumValidatorDefinition{Enum: []bool{true, false, true}},
			Error:      booleans.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := booleans.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: fail to NewEnumValidator with error %v", c.Message, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := booleans.EnumValidatorDefinition{
		Enum: []bool{true},
	}
	v, err := booleans.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator with error %v", err)
	}

	type Case struct {
		Message string
		Input   bool
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists in Enum",
			Input:   true,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enum",
			Input:   false,
			Error: &booleans.EnumValidationError{
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
