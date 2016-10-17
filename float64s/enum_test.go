package float64s_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/float64s"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition float64s.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: float64s.EnumValidatorDefinition{Enumerate: []float64{}},
			Error:      float64s.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: float64s.EnumValidatorDefinition{Enumerate: []float64{1.1}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: float64s.EnumValidatorDefinition{Enumerate: []float64{-1.1, 0, 1.1}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: float64s.EnumValidatorDefinition{Enumerate: []float64{-1.1, 0, -1.1}},
			Error:      float64s.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := float64s.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected '%s', but actual '%s'", c.Message, c.Error, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := float64s.EnumValidatorDefinition{
		Enumerate: []float64{-1.1, 0, 1.1},
	}
	v, err := float64s.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator: %s", err)
	}

	type Case struct {
		Message string
		Input   float64
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists at first in Enumerate",
			Input:   -1.1,
			Error:   nil,
		},
		{
			Message: "a value exists at second in Enumerate",
			Input:   0,
			Error:   nil,
		},
		{
			Message: "a value exists at end in Enumerate",
			Input:   1.1,
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enumerate",
			Input:   1.2,
			Error: &float64s.EnumValidationError{
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
