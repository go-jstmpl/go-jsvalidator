package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestNewEnumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition strings.EnumValidatorDefinition
		Error      error
	}

	cases := []Case{
		{
			Message:    "empty slice",
			Definition: strings.EnumValidatorDefinition{Enumerate: []string{}},
			Error:      strings.EnumDefinitionEmptyError,
		},
		{
			Message:    "single element",
			Definition: strings.EnumValidatorDefinition{Enumerate: []string{"foo"}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: strings.EnumValidatorDefinition{Enumerate: []string{"foo", "bar"}},
			Error:      nil,
		},
		{
			Message:    "duplicated elements",
			Definition: strings.EnumValidatorDefinition{Enumerate: []string{"foo", "bar", "foo"}},
			Error:      strings.EnumDefinitionDuplicationError,
		},
	}

	for _, c := range cases {
		_, err := strings.NewEnumValidator(c.Definition)
		if !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestEnumvalidator(t *testing.T) {
	def := strings.EnumValidatorDefinition{
		Enumerate: []string{"foo", "bar", "baz"},
	}
	v, err := strings.NewEnumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewEnumValidator with error %v", err)
	}

	type Case struct {
		Message string
		Input   string
		Error   error
	}
	cases := []Case{
		{
			Message: "a value exists at first in Enumerate",
			Input:   "foo",
			Error:   nil,
		},
		{
			Message: "a value exists at second in Enumerate",
			Input:   "bar",
			Error:   nil,
		},
		{
			Message: "a value exists at end in Enumerate",
			Input:   "baz",
			Error:   nil,
		},
		{
			Message: "a value doesn't exist in Enumerate",
			Input:   "qux",
			Error: &strings.EnumValidationError{
				Input:      "qux",
				Definition: def,
			},
		},
		{
			Message: "empty value",
			Input:   "",
			Error: &strings.EnumValidationError{
				Input:      "",
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
