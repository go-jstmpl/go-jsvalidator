package validator_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator"
)

func TestEnum(t *testing.T) {
	_, err := validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []int{}})
	_, ok := err.(validator.EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []int{0, 1, 0}})
	_, ok = err.(validator.DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []float64{}})
	_, ok = err.(validator.EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []float64{0.9, 1.0, 1.0}})
	_, ok = err.(validator.DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []string{}})
	_, ok = err.(validator.EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []string{"foo", "bar", "foo"}})
	_, ok = err.(validator.DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = validator.NewEnumValidator(validator.EnumValidatorDefinition{Enumerate: []bool{true, false}})
	_, ok = err.(validator.TypeError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "TypeError")
	}
}

type EnumValidatorTestCase struct {
	Input    interface{}
	Expected error
}

func TestEnumvalidator(t *testing.T) {
	definition := validator.EnumValidatorDefinition{
		Enumerate: []int{401, 402, 403},
	}
	va, err := validator.NewEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []EnumValidatorTestCase{
		{
			Input:    401,
			Expected: nil,
		},
		{
			Input:    402,
			Expected: nil,
		},
		{
			Input:    403,
			Expected: nil,
		},
		{
			Input: 501,
			Expected: &validator.EnumValidationError{
				Input:      501,
				Definition: definition,
			},
		},
		{
			Input: 10.1,
			Expected: validator.TypeError{
				Message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: validator.TypeError{
				Message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = validator.EnumValidatorDefinition{
		Enumerate: []float64{0.9, 1.0, 1.1},
	}
	va, err = validator.NewEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []EnumValidatorTestCase{
		{
			Input:    0.9,
			Expected: nil,
		},
		{
			Input:    1.0,
			Expected: nil,
		},
		{
			Input:    1.1,
			Expected: nil,
		},
		{
			Input: 1.5,
			Expected: &validator.EnumValidationError{
				Input:      1.5,
				Definition: definition,
			},
		},
		{
			Input: "hoge",
			Expected: validator.TypeError{
				Message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: validator.TypeError{
				Message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = validator.EnumValidatorDefinition{
		Enumerate: []string{"foo", "bar", "baz"},
	}
	va, err = validator.NewEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	cases = []EnumValidatorTestCase{
		{
			Input:    "foo",
			Expected: nil,
		},
		{
			Input:    "bar",
			Expected: nil,
		},
		{
			Input:    "baz",
			Expected: nil,
		},
		{
			Input: "qux",
			Expected: &validator.EnumValidationError{
				Input:      "qux",
				Definition: definition,
			},
		},
		{
			Input: 10,
			Expected: validator.TypeError{
				Message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: validator.TypeError{
				Message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
