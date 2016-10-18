package strings_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/strings"
)

func TestMaxLength(t *testing.T) {
	_, err := strings.NewMaxLengthValidator(strings.MaxLengthValidatorDefinition{MaxLength: -1})
	_, ok := err.(strings.NoLengthError)
	if !ok {
		t.Errorf("Type of error expected %v, but not.", "NoLengthError")
	}
}

func TestMaxLengthValidator(t *testing.T) {
	definition := strings.MaxLengthValidatorDefinition{
		MaxLength: 5,
	}
	va, err := strings.NewMaxLengthValidator(definition)
	if err != nil {
		t.Error(err.Error())
	}

	type MaxLengthValidatorTestCase struct {
		Input    string
		Expected error
	}

	cases := []MaxLengthValidatorTestCase{
		{
			Input:    "あいうえ",
			Expected: nil,
		},
		{
			Input:    "あいうえお",
			Expected: nil,
		},
		{
			Input: "あいうえおか",
			Expected: &strings.MaxLengthValidationError{
				Input:      "あいうえおか",
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
