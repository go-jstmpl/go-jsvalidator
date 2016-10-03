package validator

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	type FormatTestCase struct {
		Definition FormatValidatorDefinition
		Expected   error
	}
	cases := []FormatTestCase{
		{
			Definition: FormatValidatorDefinition{Format: "password"},
			Expected:   InvalidFormatError{},
		},
	}

	for _, c := range cases {
		_, err := NewFormatValidator(c.Definition)
		if reflect.TypeOf(err) != reflect.TypeOf(c.Expected) {
			t.Errorf("expected %v, but actual %v", reflect.TypeOf(c.Expected), reflect.TypeOf(err))
		}
	}
}

func TestFormatValidator(t *testing.T) {
	definition := FormatValidatorDefinition{
		Format: "date-time",
	}
	validator, err := NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type FormatValidationTestCase struct {
		Input    string
		Expected error
	}
	cases := []FormatValidationTestCase{
		{
			Input:    "2016-05-09T19:45:32Z",
			Expected: nil,
		},
		{
			Input: "209385790284750",
			Expected: &FormatValidationError{
				Input:      "209385790284750",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected:%v ,actual:%v", c.Expected, err)
		}
	}
	definition = FormatValidatorDefinition{
		Format: "email",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	cases = []FormatValidationTestCase{
		{
			Input:    "foo@bar.com",
			Expected: nil,
		},
		{
			Input: "foobar.com",
			Expected: &FormatValidationError{
				Input:      "foobar.com",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar",
			Expected: &FormatValidationError{
				Input:      "foo@bar",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar.",
			Expected: &FormatValidationError{
				Input:      "foo@bar.",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = FormatValidatorDefinition{
		Format: "hostname",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	cases = []FormatValidationTestCase{
		{
			Input:    "example",
			Expected: nil,
		},
		{
			Input:    "example.com",
			Expected: nil,
		},
		{
			Input:    "example.example.com",
			Expected: nil,
		},
		{
			Input:    "example-example.com",
			Expected: nil,
		},
		{
			Input: "example@example.com",
			Expected: &FormatValidationError{
				Input:      "example@example.com",
				Definition: definition,
			},
		},
		{
			Input: "example,com",
			Expected: &FormatValidationError{
				Input:      "example,com",
				Definition: definition,
			},
		},
		{
			Input: "example..com",
			Expected: &FormatValidationError{
				Input:      "example..com",
				Definition: definition,
			},
		},
		{
			Input: ".example.com",
			Expected: &FormatValidationError{
				Input:      ".example.com",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := validator.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = FormatValidatorDefinition{
		Format: "uri",
	}
	validator, err = NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	cases = []FormatValidationTestCase{
		{
			Input:    "https://example.com",
			Expected: nil,
		},
		{
			Input:    "http://example.com",
			Expected: nil,
		},
		{
			Input:    "https://example.com/foo/bar",
			Expected: nil,
		},
		{
			Input:    "ftp://example.com",
			Expected: nil,
		},
		{
			Input: "foobar.com",
			Expected: &FormatValidationError{
				Input:      "foobar.com",
				Definition: definition,
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
