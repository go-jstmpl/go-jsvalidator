package validator

import (
	"github.com/gocraft/dbr"
	"reflect"
	"testing"
)

type RequiredTestCase struct {
	Definition RequiredValidatorDefinition
	Expected   error
}

func TestRequired(t *testing.T) {
	tests := []RequiredTestCase{{
		Definition: RequiredValidatorDefinition{Required: []string{}},
		Expected:   EmptyError{},
	}, {
		Definition: RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}},
		Expected:   DuplicationError{},
	}, {
		Definition: RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}},
		Expected:   DuplicationError{},
	}, {
		Definition: RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}},
		Expected:   DuplicationError{},
	}, {
		Definition: RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}},
		Expected:   DuplicationError{},
	}}
	for _, test := range tests {
		_, err := NewRequiredValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type Sample struct {
	ID   dbr.NullInt64
	Name dbr.NullString
	Addr dbr.NullString
}

type RequiredValidatorTestCase struct {
	Input    *Sample
	Expected error
}

func TestRequiredValidator(t *testing.T) {
	definition := RequiredValidatorDefinition{
		Required: []string{"ID", "Addr"},
	}
	validator, err := NewRequiredValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	sample1 := &Sample{
		ID:   dbr.NewNullInt64(1),
		Name: dbr.NewNullString("MyName"),
		Addr: dbr.NewNullString("foo@bar.com"),
	}
	sample2 := &Sample{
		ID:   dbr.NewNullInt64(2),
		Name: dbr.NewNullString(nil),
		Addr: dbr.NewNullString("foo@bar.com"),
	}
	sample3 := &Sample{
		ID:   dbr.NewNullInt64(nil),
		Name: dbr.NewNullString("hi"),
		Addr: dbr.NewNullString("foo@bar.com"),
	}
	tests := []RequiredValidatorTestCase{{
		Input:    sample1,
		Expected: nil,
	}, {
		Input:    sample2,
		Expected: nil,
	}, {
		Input: sample3,
		Expected: &RequiredValidationError{
			Input:      sample3,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
