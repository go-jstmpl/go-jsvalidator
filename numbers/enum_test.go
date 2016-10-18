package numbers_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/numbers"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition numbers.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: numbers.EnumValidatorDefinition{Enum: []float64{}},
			Error:      numbers.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: numbers.EnumValidatorDefinition{Enum: []float64{1.1}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: numbers.EnumValidatorDefinition{Enum: []float64{-1.1, 0, 1.1}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: numbers.EnumValidatorDefinition{Enum: []float64{-1.1, 0, -1.1}},
			Error:      numbers.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := numbers.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: fail to NewEnumValidator with error %v", c.Message, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := numbers.EnumValidatorDefinition{
		Enum: []float64{-1.1, 0, 1.1},
	}
	v, err := numbers.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator with error %v", err)
	}

	type Case struct {
		Message string
		Input   float64
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists at first in Enum",
			Input:   -1.1,
			Error:   nil,
		},
		{
			Message: "a value exists at second in Enum",
			Input:   0,
			Error:   nil,
		},
		{
			Message: "a value exists at end in Enum",
			Input:   1.1,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enum",
			Input:   1.2,
			Error: &numbers.EnumValidationError{
				Input:      1.2,
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
