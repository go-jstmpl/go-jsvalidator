package validator

import (
	"reflect"
	"testing"
)

type IntMaxItemsTestCase struct {
	Definition IntMaxItemsValidatorDefinition
	Expected   error
}

func TestIntMaxItems(t *testing.T) {
	tests := []IntMaxItemsTestCase{{
		Definition: IntMaxItemsValidatorDefinition{MaxItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewIntMaxItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type IntMaxItemsValidatorTestCase struct {
	Input    []int
	Expected error
}

func TestIntMaxItemsValidator(t *testing.T) {
	definition := IntMaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	validator, err := NewIntMaxItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []IntMaxItemsValidatorTestCase{{
		Input:    []int{1},
		Expected: nil,
	}, {
		Input:    []int{1, 2},
		Expected: nil,
	}, {
		Input:    []int{1, 2, 3},
		Expected: nil,
	}, {
		Input: []int{1, 2, 3, 4},
		Expected: &IntMaxItemsValidationError{
			Input:      []int{1, 2, 3, 4},
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

type StringMaxItemsTestCase struct {
	Definition StringMaxItemsValidatorDefinition
	Expected   error
}

func TestStringMaxItems(t *testing.T) {
	tests := []StringMaxItemsTestCase{{
		Definition: StringMaxItemsValidatorDefinition{MaxItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewStringMaxItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type StringMaxItemsValidatorTestCase struct {
	Input    []string
	Expected error
}

func TestStringMaxItemsValidator(t *testing.T) {
	definition := StringMaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	validator, err := NewStringMaxItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []StringMaxItemsValidatorTestCase{{
		Input:    []string{"foo"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar"},
		Expected: nil,
	}, {
		Input:    []string{"foo", "bar", "baz"},
		Expected: nil,
	}, {
		Input: []string{"foo", "bar", "bas", "qux"},
		Expected: &StringMaxItemsValidationError{
			Input: []string{"foo", "bar", "bas", "qux"},

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

type FloatMaxItemsTestCase struct {
	Definition FloatMaxItemsValidatorDefinition
	Expected   error
}

func TestFloatMaxItems(t *testing.T) {
	tests := []FloatMaxItemsTestCase{{
		Definition: FloatMaxItemsValidatorDefinition{MaxItems: -1},
		Expected:   NoLengthError{},
	}}
	for _, test := range tests {
		_, err := NewFloatMaxItemsValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type FloatMaxItemsValidatorTestCase struct {
	Input    []float64
	Expected error
}

func TestFloatMaxItemsValidator(t *testing.T) {
	definition := FloatMaxItemsValidatorDefinition{
		MaxItems: 3,
	}
	validator, err := NewFloatMaxItemsValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []FloatMaxItemsValidatorTestCase{{
		Input:    []float64{1},
		Expected: nil,
	}, {
		Input:    []float64{1, 2},
		Expected: nil,
	}, {
		Input:    []float64{1, 2, 3},
		Expected: nil,
	}, {
		Input: []float64{1, 2, 3, 4},
		Expected: &FloatMaxItemsValidationError{
			Input:      []float64{1, 2, 3, 4},
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
