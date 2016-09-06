package validator

import (
	"reflect"
	"testing"

	"github.com/gocraft/dbr"
)

type IntMaximumValidatorTestCase struct {
	Input    int
	Expected *IntMaximumValidationError
}

func TestIntMaximumValidator(t *testing.T) {
	// Case exclusive is false
	definition := IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: false,
	}

	validator, err := NewIntMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []IntMaximumValidatorTestCase{{
		Input:    99,
		Expected: nil,
	}, {
		Input:    100,
		Expected: nil,
	}, {
		Input: 101,
		Expected: &IntMaximumValidationError{
			Input:      101,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = IntMaximumValidatorDefinition{
		Maximum:   100,
		Exclusive: true,
	}

	validator, err = NewIntMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []IntMaximumValidatorTestCase{{
		Input:    99,
		Expected: nil,
	}, {
		Input: 100,
		Expected: &IntMaximumValidationError{
			Input:      100,
			Definition: definition,
		},
	}, {
		Input: 101,
		Expected: &IntMaximumValidationError{
			Input:      101,
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

type IntMinimumValidatorTestCase struct {
	Input    int
	Expected *IntMinimumValidationError
}

func TestIntMinimumValidator(t *testing.T) {
	// Case exclusive is false
	definition := IntMinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: false,
	}

	validator, err := NewIntMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []IntMinimumValidatorTestCase{{
		Input:    101,
		Expected: nil,
	}, {
		Input:    100,
		Expected: nil,
	}, {
		Input: 99,
		Expected: &IntMinimumValidationError{
			Input:      99,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = IntMinimumValidatorDefinition{
		Minimum:   100,
		Exclusive: true,
	}

	validator, err = NewIntMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []IntMinimumValidatorTestCase{{
		Input:    101,
		Expected: nil,
	}, {
		Input: 100,
		Expected: &IntMinimumValidationError{
			Input:      100,
			Definition: definition,
		},
	}, {
		Input: 99,
		Expected: &IntMinimumValidationError{
			Input:      99,
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

type FloatMaximumValidatorTestCase struct {
	Input    float64
	Expected *FloatMaximumValidationError
}

func TestFloatMaximumValidator(t *testing.T) {
	// Case exclusive is false
	definition := FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	validator, err := NewFloatMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []FloatMaximumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input:    1.0,
		Expected: nil,
	}, {
		Input: 1.1,
		Expected: &FloatMaximumValidationError{
			Input:      1.1,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = FloatMaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	validator, err = NewFloatMaximumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []FloatMaximumValidatorTestCase{{
		Input:    0.9,
		Expected: nil,
	}, {
		Input: 1.0,
		Expected: &FloatMaximumValidationError{
			Input:      1.0,
			Definition: definition,
		},
	}, {
		Input: 1.1,
		Expected: &FloatMaximumValidationError{
			Input:      1.1,
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

type FloatMinimumValidatorTestCase struct {
	Input    float64
	Expected *FloatMinimumValidationError
}

func TestFloatMinimumValidator(t *testing.T) {
	// Case exclusive is false
	definition := FloatMinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: false,
	}

	validator, err := NewFloatMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []FloatMinimumValidatorTestCase{{
		Input:    1.1,
		Expected: nil,
	}, {
		Input:    1.0,
		Expected: nil,
	}, {
		Input: 0.9,
		Expected: &FloatMinimumValidationError{
			Input:      0.9,
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	// Case exclusive is true
	definition = FloatMinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: true,
	}

	validator, err = NewFloatMinimumValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}
	tests = []FloatMinimumValidatorTestCase{{
		Input:    1.1,
		Expected: nil,
	}, {
		Input: 1.0,
		Expected: &FloatMinimumValidationError{
			Input:      1.0,
			Definition: definition,
		},
	}, {
		Input: 0.9,
		Expected: &FloatMinimumValidationError{
			Input:      0.9,
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

type MaxLengthTestCase struct {
	Definition MaxLengthValidatorDefinition
	Expected   error
}

func TestMaxLength(t *testing.T) {
	tests := []MaxLengthTestCase{{
		Definition: MaxLengthValidatorDefinition{MaxLength: -1},
		Expected:   NoLengthError{},
	}}

	for _, test := range tests {
		_, err := NewMaxLengthValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type MaxLengthValidatorTestCase struct {
	Input    string
	Expected *MaxLengthValidationError
}

func TestMaxLengthValidator(t *testing.T) {
	definition := MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	validator, err := NewMaxLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []MaxLengthValidatorTestCase{{
		Input:    "あいうえ",
		Expected: nil,
	}, {
		Input:    "あいうえお",
		Expected: nil,
	}, {
		Input: "あいうえおか",
		Expected: &MaxLengthValidationError{
			Input:      "あいうえおか",
			Definition: definition,
		},
	}, {
		Input:    "abcde",
		Expected: nil,
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
}

type PatternTestCase struct {
	Definition PatternValidatorDefinition
	Expected   error
}

func TestPattern(t *testing.T) {
	tests := []PatternTestCase{{
		Definition: PatternValidatorDefinition{Pattern: ""},
		Expected:   EmptyError{},
	}, {
		Definition: PatternValidatorDefinition{Pattern: "[a-z"},
		Expected:   InvalidPatternError{},
	}}
	for _, test := range tests {
		_, err := NewPatternValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type PatternValidatorTestCase struct {
	Input    string
	Expected *PatternValidationError
}

func TestPatternValidator(t *testing.T) {
	definition := PatternValidatorDefinition{
		Pattern: "^\\d{7}$",
	}
	validator, err := NewPatternValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []PatternValidatorTestCase{{
		Input:    "1234567",
		Expected: nil,
	}, {
		Input: "abcdefg",
		Expected: &PatternValidationError{
			Input:      "abcdefg",
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
	Expected *IntEnumValidationError
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
	Expected *FloatEnumValidationError
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
	Expected *StringEnumValidationError
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
	Expected *IntMaxItemsValidationError
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
	Expected *StringMaxItemsValidationError
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
	Expected *FloatMaxItemsValidationError
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

type FormatTestCase struct {
	Definition FormatValidatorDefinition
	Expected   error
}

func TestFormat(t *testing.T) {
	tests := []FormatTestCase{{
		Definition: FormatValidatorDefinition{Format: "password"},
		Expected:   InvalidFormatError{},
	}}

	for _, test := range tests {
		_, err := NewFormatValidator(test.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(test.Expected) {
			t.Errorf("expected:%v, actual:%v", reflect.TypeOf(test.Expected), reflect.TypeOf(err))
		}
	}
}

type FormatValidationTestCase struct {
	Input    string
	Expected *FormatValidationError
}

func TestFormatValidator(t *testing.T) {
	definition := FormatValidatorDefinition{
		Format: "date-time",
	}
	validator, err := NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []FormatValidationTestCase{{
		Input:    "2016-05-09T19:45:32Z",
		Expected: nil,
	}, {
		Input: "209385790284750",
		Expected: &FormatValidationError{
			Input:      "209385790284750",
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}
	definition = FormatValidatorDefinition{
		Format: "email",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests = []FormatValidationTestCase{{
		Input:    "foo@bar.com",
		Expected: nil,
	}, {
		Input: "foobar.com",
		Expected: &FormatValidationError{
			Input:      "foobar.com",
			Definition: definition,
		},
	}, {
		Input: "foo@bar",
		Expected: &FormatValidationError{
			Input:      "foo@bar",
			Definition: definition,
		},
	}, {
		Input: "foo@bar.",
		Expected: &FormatValidationError{
			Input:      "foo@bar.",
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	definition = FormatValidatorDefinition{
		Format: "hostname",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests = []FormatValidationTestCase{{
		Input:    "example",
		Expected: nil,
	}, {
		Input:    "example.com",
		Expected: nil,
	}, {
		Input:    "example.example.com",
		Expected: nil,
	}, {
		Input:    "example-example.com",
		Expected: nil,
	}, {
		Input: "example@example.com",
		Expected: &FormatValidationError{
			Input:      "example@example.com",
			Definition: definition,
		},
	}, {
		Input: "example,com",
		Expected: &FormatValidationError{
			Input:      "example,com",
			Definition: definition,
		},
	}, {
		Input: "example..com",
		Expected: &FormatValidationError{
			Input:      "example..com",
			Definition: definition,
		},
	}, {
		Input: ".example.com",
		Expected: &FormatValidationError{
			Input:      ".example.com",
			Definition: definition,
		},
	}}

	for _, test := range tests {
		err := validator.Validate(test.Input)
		if !reflect.DeepEqual(err, test.Expected) {
			t.Errorf("expected:%v ,actual:%v", test.Expected, err)
		}
	}

	definition = FormatValidatorDefinition{
		Format: "uri",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	tests = []FormatValidationTestCase{{
		Input:    "https://example.com",
		Expected: nil,
	}, {
		Input:    "http://example.com",
		Expected: nil,
	}, {
		Input:    "https://example.com/foo/bar",
		Expected: nil,
	}, {
		Input:    "ftp://example.com",
		Expected: nil,
	}, {
		Input: "foobar.com",
		Expected: &FormatValidationError{
			Input:      "foobar.com",
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
