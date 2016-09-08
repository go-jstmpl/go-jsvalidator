package validator

import (
	"reflect"
	"testing"
)

type IntMinItemsTestCase struct {
	Definition IntMinItemsValidatorDefinition
	Expected   error
}

func TestIntMinItems(t *testing.T) {
	tests := []IntMinItemsTestCase{{
		Definition: IntMinItemsValidatorDefinition{MinItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewIntMinItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type IntMinItemsValidatorTestCase struct {
	Input    []int
	Expected error
}

func TestIntMinItemsValidator(t *testing.T) {
	definition := IntMinItemsValidatorDefinition{
		MinItems: 3,
	}
	validator, err := NewIntMinItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []IntMinItemsValidatorTestCase{{
		Input: []int{1},
		Expected: &IntMinItemsValidationError{
			Input:      []int{1},
			Definition: definition,
		},
	}, {
		Input: []int{1, 2},
		Expected: &IntMinItemsValidationError{
			Input:      []int{1, 2},
			Definition: definition,
		},
	}, {
		Input:    []int{1, 2, 3},
		Expected: nil,
	}, {
		Input:    []int{1, 2, 3, 4},
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}

type StringMinItemsTestCase struct {
	Definition StringMinItemsValidatorDefinition
	Expected   error
}

func TestStringMinItems(t *testing.T) {
	tests := []StringMinItemsTestCase{{
		Definition: StringMinItemsValidatorDefinition{MinItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewStringMinItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type StringMinItemsValidatorTestCase struct {
	Input    []string
	Expected error
}

func TestStringMinItemsValidator(t *testing.T) {
	definition := StringMinItemsValidatorDefinition{
		MinItems: 3,
	}
	validator, err := NewStringMinItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []StringMinItemsValidatorTestCase{{
		Input: []string{"foo"},
		Expected: &StringMinItemsValidationError{
			Input: []string{"foo"},

			Definition: definition,
		},
	}, {
		Input:    []string{"foo", "bar", "baz"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar", "bas", "qux"},
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}

type FloatMinItemsTestCase struct {
	Definition FloatMinItemsValidatorDefinition
	Expected   error
}

func TestFloatMinItems(t *testing.T) {
	tests := []FloatMinItemsTestCase{{
		Definition: FloatMinItemsValidatorDefinition{MinItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewFloatMinItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type FloatMinItemsValidatorTestCase struct {
	Input    []float64
	Expected error
}

func TestFloatMinItemsValidator(t *testing.T) {
	definition := FloatMinItemsValidatorDefinition{
		MinItems: 3,
	}
	validator, err := NewFloatMinItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []FloatMinItemsValidatorTestCase{{
		Input: []float64{1},
		Expected: &FloatMinItemsValidationError{
			Input:      []float64{1},
			Definition: definition,
		},
	}, {
		Input:    []float64{1, 2, 3},
		Expected: nil,
	}, {
		Input:    []float64{1, 2, 3, 4},
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}
