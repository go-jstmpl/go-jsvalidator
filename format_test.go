package validator_test

import (
	"reflect"
	"testing"

	validator "github.com/go-jstmpl/go-jsvalidator"
)

func TestFormat(t *testing.T) {
	_, err := validator.NewFormatValidator(validator.FormatValidatorDefinition{Format: "password"})
	_, ok := err.(validator.InvalidFormatError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "InvalidFormatError")
	}
}

func TestFormatValidator(t *testing.T) {
	definition := validator.FormatValidatorDefinition{
		Format: "date-time",
	}
	va, err := validator.NewFormatValidator(definition)
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
			Expected: &validator.FormatValidationError{
				Input:      "209385790284750",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected:%v ,actual:%v", c.Expected, err)
		}
	}
	definition = validator.FormatValidatorDefinition{
		Format: "email",
	}
	va, err = validator.NewFormatValidator(definition)
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
			Expected: &validator.FormatValidationError{
				Input:      "foobar.com",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar",
			Expected: &validator.FormatValidationError{
				Input:      "foo@bar",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar.",
			Expected: &validator.FormatValidationError{
				Input:      "foo@bar.",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = validator.FormatValidatorDefinition{
		Format: "hostname",
	}
	va, err = validator.NewFormatValidator(definition)
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
			Expected: &validator.FormatValidationError{
				Input:      "example@example.com",
				Definition: definition,
			},
		},
		{
			Input: "example,com",
			Expected: &validator.FormatValidationError{
				Input:      "example,com",
				Definition: definition,
			},
		},
		{
			Input: "example..com",
			Expected: &validator.FormatValidationError{
				Input:      "example..com",
				Definition: definition,
			},
		},
		{
			Input: ".example.com",
			Expected: &validator.FormatValidationError{
				Input:      ".example.com",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = validator.FormatValidatorDefinition{
		Format: "uri",
	}
	va, err = validator.NewFormatValidator(definition)
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
			Expected: &validator.FormatValidationError{
				Input:      "foobar.com",
				Definition: definition,
			},
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}

	definition = validator.FormatValidatorDefinition{
		Format: "password-0Aa",
	}
	va, err = validator.NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	cases = []FormatValidationTestCase{
		{
			Input:    "0Aa",
			Expected: nil,
		},
		{
			Input:    "0aA",
			Expected: nil,
		},
		{
			Input:    "A0a",
			Expected: nil,
		},
		{
			Input:    "Aa0",
			Expected: nil,
		},
		{
			Input:    "a0A",
			Expected: nil,
		},
		{
			Input:    "aA0",
			Expected: nil,
		},
		{
			Input:    "!0Aa",
			Expected: nil,
		},
		{
			Input:    "!0aA",
			Expected: nil,
		},
		{
			Input:    "!A0a",
			Expected: nil,
		},
		{
			Input:    "!Aa0",
			Expected: nil,
		},
		{
			Input:    "!a0A",
			Expected: nil,
		},
		{
			Input:    "!aA0",
			Expected: nil,
		},
		{
			Input:    "0!Aa",
			Expected: nil,
		},
		{
			Input:    "0!aA",
			Expected: nil,
		},
		{
			Input:    "0A!a",
			Expected: nil,
		},
		{
			Input:    "0Aa!",
			Expected: nil,
		},
		{
			Input:    "0a!A",
			Expected: nil,
		},
		{
			Input:    "0aA!",
			Expected: nil,
		},
		{
			Input:    "A!0a",
			Expected: nil,
		},
		{
			Input:    "A!a0",
			Expected: nil,
		},
		{
			Input:    "A0!a",
			Expected: nil,
		},
		{
			Input:    "A0a!",
			Expected: nil,
		},
		{
			Input:    "Aa!0",
			Expected: nil,
		},
		{
			Input:    "Aa0!",
			Expected: nil,
		},
		{
			Input:    "a!0A",
			Expected: nil,
		},
		{
			Input:    "a!A0",
			Expected: nil,
		},
		{
			Input:    "a0!A",
			Expected: nil,
		},
		{
			Input:    "a0A!",
			Expected: nil,
		},
		{
			Input:    "aA!0",
			Expected: nil,
		},
		{
			Input:    "aA0!",
			Expected: nil,
		},
		{
			Input:    "AZaz09!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`",
			Expected: nil,
		},
		{
			Input: "aA0!あ",
			Expected: &validator.FormatValidationError{
				Input:      "aA0!あ",
				Definition: definition,
			},
		},
		{
			Input: "12345678",
			Expected: &validator.FormatValidationError{
				Input:      "12345678",
				Definition: definition,
			},
		},
		{
			Input: "password",
			Expected: &validator.FormatValidationError{
				Input:      "password",
				Definition: definition,
			},
		},
		{
			Input: "PASSWORD",
			Expected: &validator.FormatValidationError{
				Input:      "PASSWORD",
				Definition: definition,
			},
		},
		{
			Input: "Password",
			Expected: &validator.FormatValidationError{
				Input:      "Password",
				Definition: definition,
			},
		},
		{
			Input: "password123",
			Expected: &validator.FormatValidationError{
				Input:      "password123",
				Definition: definition,
			},
		},
		{
			Input: "PASSWORD123",
			Expected: &validator.FormatValidationError{
				Input:      "PASSWORD123",
				Definition: definition,
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
