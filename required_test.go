package validator

import (
	"reflect"
	"testing"

	"github.com/gocraft/dbr"
)

func TestRequired(t *testing.T) {
	type RequiredTestCase struct {
		Definition RequiredValidatorDefinition
		Expected   error
	}
	cases := []RequiredTestCase{
		{
			Definition: RequiredValidatorDefinition{Required: []string{}},
			Expected:   EmptyError{},
		},
		{
			Definition: RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}},
			Expected:   DuplicationError{},
		},
		{
			Definition: RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}},
			Expected:   DuplicationError{},
		},
		{
			Definition: RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}},
			Expected:   DuplicationError{},
		},
		{
			Definition: RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}},
			Expected:   DuplicationError{},
		},
	}
	for _, c := range cases {
		_, err := NewRequiredValidator(c.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(c.Expected) {
			t.Errorf("expected %v, but actual %v", reflect.TypeOf(c.Expected), reflect.TypeOf(err))
		}
	}
}

func TestRequiredValidator(t *testing.T) {
	definition := RequiredValidatorDefinition{
		Required: []string{"ID", "Addr"},
	}
	validator, err := NewRequiredValidator(definition)
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
			Expected: &InvalidFieldTypeError{
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
			Expected: &RequiredValidationError{
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
			Expected: &RequiredValidationError{
				Input:      &sample3,
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("%s: expected %+v, but actual %+v", c.Message, c.Expected, err)
		}
	}
}
