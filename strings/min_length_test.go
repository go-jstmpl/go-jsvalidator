package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestMinLength(t *testing.T) {
	_, err := strings.NewMinLengthValidator(strings.MinLengthValidatorDefinition{MinLength: -1})
	_, ok := err.(strings.NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMinLengthValidator(t *testing.T) {
	definition := strings.MinLengthValidatorDefinition{
		MinLength: 5,
	}
	va, err := strings.NewMinLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MinLengthValidatorTestCase struct {
		Input    string
		Expected error
	}
	cases := []MinLengthValidatorTestCase{
		{
			Input:    "あいうえおか",
			Expected: nil,
		},
		{
			Input:    "あいうえお",
			Expected: nil,
		},
		{
			Input: "あいうえ",
			Expected: &strings.MinLengthValidationError{
				Input:      "あいうえ",
				Definition: definition,
			},
		},
		{
			Input:    "abcde",
			Expected: nil,
		},
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("expected %v, but actual %v", c.Expected, err)
		}
	}
}
