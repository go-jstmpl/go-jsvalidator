package validator

import (
	"reflect"
	"testing"
)

type IntEnumTestCase struct {
	Definition IntEnumValidatorDefinition
	Expected   error
}

func TestIntEnum(t *testing.T) {
	tests := []IntEnumTestCase{{
		Definition: IntEnumValidatorDefinition{Enumerate: []int{}},
		Expected:   EmptyError{},
	}, {
		Definition: IntEnumValidatorDefinition{Enumerate: []int{0, 1, 0}},
		Expected:   DuplicationError{},
	}}
	for _, test := range tests {
		_, err := NewIntEnumValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type IntEnumValidatorTestCase struct {
	Input    int
	Expected error
}

func TestIntEnumvalidator(t *testing.T) {
	definition := IntEnumValidatorDefinition{
		Enumerate: []int{401, 402, 403},
	}
	validator, err := NewIntEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []IntEnumValidatorTestCase{{
		Input:    401,
		Expected: nil,
	}, {
		Input:    402,
		Expected: nil,
	}, {
		Input:    403,
		Expected: nil,
	}, {
		Input: 501,
		Expected: &IntEnumValidationError{
			Input:      501,
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

type FloatEnumTestCase struct {
	Definition FloatEnumValidatorDefinition
	Expected   error
}

func TestFloatEnum(t *testing.T) {
	tests := []FloatEnumTestCase{{
		Definition: FloatEnumValidatorDefinition{Enumerate: []float64{}},
		Expected:   EmptyError{},
	}, {
		Definition: FloatEnumValidatorDefinition{Enumerate: []float64{0.9, 1.0, 1.0}},
		Expected:   DuplicationError{},
	}}
	for _, test := range tests {
		_, err := NewFloatEnumValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type FloatEnumValidatorTestCase struct {
	Input    float64
	Expected error
}

func TestFloatEnumvalidator(t *testing.T) {
	definition := FloatEnumValidatorDefinition{
		Enumerate: []float64{0.9, 1.0, 1.1},
	}
	validator, err := NewFloatEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []FloatEnumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input:    1.0,
		Expected: nil,
	}, {
		Input:    1.1,
		Expected: nil,
	}, {
		Input: 1.5,
		Expected: &FloatEnumValidationError{
			Input:      1.5,
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

type StringEnumTestCase struct {
	Definition StringEnumValidatorDefinition
	Expected   error
}

func TestStringEnum(t *testing.T) {
	tests := []StringEnumTestCase{{
		Definition: StringEnumValidatorDefinition{Enumerate: []string{}},
		Expected:   EmptyError{},
	}, {
		Definition: StringEnumValidatorDefinition{Enumerate: []string{"foo", "bar", "foo"}},
		Expected:   DuplicationError{},
	}}
	for _, test := range tests {
		_, err := NewStringEnumValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type StringEnumValidatorTestCase struct {
	Input    string
	Expected error
}

func TestStringEnumvalidator(t *testing.T) {
	definition := StringEnumValidatorDefinition{
		Enumerate: []string{"foo", "bar", "baz"},
	}
	validator, err := NewStringEnumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []StringEnumValidatorTestCase{{
		Input:    "foo",
		Expected: nil,
	}, {
		Input:    "bar",
		Expected: nil,
	}, {
		Input:    "baz",
		Expected: nil,
	}, {
		Input: "qux",
		Expected: &StringEnumValidationError{
			Input:      "qux",
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
