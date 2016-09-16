package validator

import (
	"reflect"
	"testing"
)

type EnumTestCase struct {
	Definition EnumValidatorDefinition
	Expected   error
}

func TestEnum(t *testing.T) {
	tests := []EnumTestCase{
		{
			Definition: EnumValidatorDefinition{Enumerate: []int{}},
			Expected:   EmptyError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []int{0, 1, 0}},
			Expected:   DuplicationError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []float64{}},
			Expected:   EmptyError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []float64{0.9, 1.0, 1.0}},
			Expected:   DuplicationError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []string{}},
			Expected:   EmptyError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []string{"foo", "bar", "foo"}},
			Expected:   DuplicationError{},
		},
		{
			Definition: EnumValidatorDefinition{Enumerate: []bool{true, false}},
			Expected:   TypeError{},
		},
	}
	for _, test := range tests {
		_, err := NewEnumValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
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
	tests := []EnumValidatorTestCase{
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
	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("\nexpected: %v \n  actual: %v", test.Expected, err)
		}
	}

	definition = EnumValidatorDefinition{
		Enumerate: []float64{0.9, 1.0, 1.1},
	}
	validator, err = NewEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []EnumValidatorTestCase{
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
	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	definition = EnumValidatorDefinition{
		Enumerate: []string{"foo", "bar", "baz"},
	}
	validator, err = NewEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []EnumValidatorTestCase{
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
	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
