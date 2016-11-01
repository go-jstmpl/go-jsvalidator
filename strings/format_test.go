package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestNewFormatValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition strings.FormatValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "date-time format",
			Definition: strings.FormatValidatorDefinition{Format: "date-time"},
			Error:      nil,
		},
		{
			Message:    "email format",
			Definition: strings.FormatValidatorDefinition{Format: "email"},
			Error:      nil,
		},
		{
			Message:    "hostname format",
			Definition: strings.FormatValidatorDefinition{Format: "hostname"},
			Error:      nil,
		},
		{
			Message:    "uri format",
			Definition: strings.FormatValidatorDefinition{Format: "uri"},
			Error:      nil,
		},
		{
			Message:    "password-0Aa format",
			Definition: strings.FormatValidatorDefinition{Format: "password-0Aa"},
			Error:      nil,
		},
		{
			Message:    "empty string",
			Definition: strings.FormatValidatorDefinition{Format: ""},
			Error:      strings.FormatDefinitionInvalidFormatError,
		},
		{
			Message:    "unsupported format",
			Definition: strings.FormatValidatorDefinition{Format: "fax-number"},
			Error:      strings.FormatDefinitionInvalidFormatError,
		},
	}
	for _, c := range cases {
		_, err := strings.NewFormatValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewFormatValidator with error %v", c.Message, err)
		}
	}
}

func TestFormatValidator(t *testing.T) {
	definition := strings.FormatValidatorDefinition{
		Format: "date-time",
	}
	va, err := strings.NewFormatValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type FormatValidationTestCase struct {
		Input    string
		Expected error
	}
	cases := []FormatValidationTestCase{
		{
			Input:    "2016-10-07T16:58:37Z",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37+09:00",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37-09:00",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091Z",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091+09:00",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091-09:00",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091232123Z",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091232123+09:00",
			Expected: nil,
		},
		{
			Input:    "2016-10-07T16:58:37.091232123-09:00",
			Expected: nil,
		},
		{
			Input: "209385790284750",
			Expected: &strings.FormatValidationError{
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
	definition = strings.FormatValidatorDefinition{
		Format: "email",
	}
	va, err = strings.NewFormatValidator(definition)
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
			Expected: &strings.FormatValidationError{
				Input:      "foobar.com",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar",
			Expected: &strings.FormatValidationError{
				Input:      "foo@bar",
				Definition: definition,
			},
		},
		{
			Input: "foo@bar.",
			Expected: &strings.FormatValidationError{
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

	definition = strings.FormatValidatorDefinition{
		Format: "hostname",
	}
	va, err = strings.NewFormatValidator(definition)
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
			Expected: &strings.FormatValidationError{
				Input:      "example@example.com",
				Definition: definition,
			},
		},
		{
			Input: "example,com",
			Expected: &strings.FormatValidationError{
				Input:      "example,com",
				Definition: definition,
			},
		},
		{
			Input: "example..com",
			Expected: &strings.FormatValidationError{
				Input:      "example..com",
				Definition: definition,
			},
		},
		{
			Input: ".example.com",
			Expected: &strings.FormatValidationError{
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

	definition = strings.FormatValidatorDefinition{
		Format: "uri",
	}
	va, err = strings.NewFormatValidator(definition)
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
			Expected: &strings.FormatValidationError{
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

	definition = strings.FormatValidatorDefinition{
		Format: "password-0Aa",
	}
	va, err = strings.NewFormatValidator(definition)
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
			Input:    "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			Expected: nil,
		},
		{
			Input: "aA0!あ",
			Expected: &strings.FormatValidationError{
				Input:      "aA0!あ",
				Definition: definition,
			},
		},
		{
			Input: "12345678",
			Expected: &strings.FormatValidationError{
				Input:      "12345678",
				Definition: definition,
			},
		},
		{
			Input: "password",
			Expected: &strings.FormatValidationError{
				Input:      "password",
				Definition: definition,
			},
		},
		{
			Input: "PASSWORD",
			Expected: &strings.FormatValidationError{
				Input:      "PASSWORD",
				Definition: definition,
			},
		},
		{
			Input: "Password",
			Expected: &strings.FormatValidationError{
				Input:      "Password",
				Definition: definition,
			},
		},
		{
			Input: "password123",
			Expected: &strings.FormatValidationError{
				Input:      "password123",
				Definition: definition,
			},
		},
		{
			Input: "PASSWORD123",
			Expected: &strings.FormatValidationError{
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
