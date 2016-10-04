package validator

import (
	"reflect"
	"testing"
)

func TestEnum(t *testing.T) {
	_, err := NewEnumValidator(EnumValidatorDefinition{Enumerate: []int{}})
	_, ok := err.(EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []int{0, 1, 0}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []float64{}})
	_, ok = err.(EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []float64{0.9, 1.0, 1.0}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []string{}})
	_, ok = err.(EmptyError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "EmptyError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []string{"foo", "bar", "foo"}})
	_, ok = err.(DuplicationError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "DuplicationError")
	}

	_, err = NewEnumValidator(EnumValidatorDefinition{Enumerate: []bool{true, false}})
	_, ok = err.(TypeError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "TypeError")
	}
}

type EnumValidatorTestCase struct {
	Input    interface{}
	Expected error
}

func TestEnumvalidator(t *testing.T) {
	definition := EnumValidatorDefinition{
		Enumerate: []int{401, 402, 403},
	}
	validator, err := NewEnumValidator(definition)
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
			Expected: &EnumValidationError{
				Input:      501,
				Definition: definition,
			},
		},
		{
			Input: 10.1,
			Expected: TypeError{
				message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: TypeError{
				message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = EnumValidatorDefinition{
		Enumerate: []float64{0.9, 1.0, 1.1},
	}
	validator, err = NewEnumValidator(definition)
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
			Expected: &EnumValidationError{
				Input:      1.5,
				Definition: definition,
			},
		},
		{
			Input: "hoge",
			Expected: TypeError{
				message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: TypeError{
				message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = EnumValidatorDefinition{
		Enumerate: []string{"foo", "bar", "baz"},
	}
	validator, err = NewEnumValidator(definition)
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
			Expected: &EnumValidationError{
				Input:      "qux",
				Definition: definition,
			},
		},
		{
			Input: 10,
			Expected: TypeError{
				message: "input and element of enumerate should be same type",
			},
		},
		{
			Input: true,
			Expected: TypeError{
				message: "input should be same type as element of enumerate",
			},
		},
	}
	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
