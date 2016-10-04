package validator

import (
	"reflect"
	"testing"

	"github.com/gocraft/dbr"
)

func TestRequired(t *testing.T) {
	_, err := NewRequiredValidator(RequiredValidatorDefinition{Required: []string{}})
	_, ok := err.(EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = NewRequiredValidator(RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewRequiredValidator(RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewRequiredValidator(RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewRequiredValidator(RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
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
