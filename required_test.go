package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
	"github.com/gocraft/dbr"
)

func TestNewRequiredValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition validator.RequiredValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "single element",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo"}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar"}},
			Error:      nil,
		},
		{
			Message:    "empty slice",
			Definition: validator.RequiredValidatorDefinition{Required: []string{}},
			Error:      validator.RequiredDefinitionEmptyError,
		},
		{
			Message:    "duplicated elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and second",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at second end end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated all elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
	}
	for _, c := range cases {
		_, err := validator.NewRequiredValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewPatternValidator with error %v", c.Message, err)
		}
	}
}

func TestValidateOfRequiredValidator(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"ID", "Addr"},
	}
	va, err := validator.NewRequiredValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type Sample struct {
		ID   dbr.NullInt64
		Name dbr.NullString
		Addr dbr.NullString
	}
	sample1 := Sample{
		ID:   dbr.NewNullInt64(1),
		Name: dbr.NewNullString("MyName"),
		Addr: dbr.NewNullString("foo@bar.com"),
	}
	sample2 := Sample{
		ID:   dbr.NewNullInt64(2),
		Name: dbr.NewNullString(nil),
		Addr: dbr.NewNullString("foo@bar.com"),
	}
	sample3 := Sample{
		ID:   dbr.NewNullInt64(nil),
		Name: dbr.NewNullString("hi"),
		Addr: dbr.NewNullString("foo@bar.com"),
	}

	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}
	cases := []RequiredValidatorTestCase{
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidFieldTypeError{
				Input:      "foo",
				Definition: definition,
			},
		},
		{
			Message:  "non-pointer of sample1",
			Input:    sample1,
			Expected: nil,
		},
		{
			Message:  "non-pointer of sample2",
			Input:    sample2,
			Expected: nil,
		},
		{
			Message: "non-pointer of sample3",
			Input:   sample3,
			Expected: &validator.RequiredValidationError{
				Input:      sample3,
				Definition: definition,
			},
		},
		{
			Message:  "pointer of sample1",
			Input:    &sample1,
			Expected: nil,
		},
		{
			Message:  "pointer of sample2",
			Input:    &sample2,
			Expected: nil,
		},
		{
			Message: "pointer of sample3",
			Input:   &sample3,
			Expected: &validator.RequiredValidationError{
				Input:      &sample3,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("%s: expected %+v, but actual %+v", c.Message, c.Expected, err)
		}
	}
}
