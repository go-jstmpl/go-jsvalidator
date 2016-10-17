package ints_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/ints"
)

func TestNewMinimumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition ints.MinimumValidatorDefinition
	}
	cases := []Case{
		Case{
			Message: "With positive number and positive exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   10,
				Exclusive: true,
			},
		},
		Case{
			Message: "With positive number and negative exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   10,
				Exclusive: false,
			},
		},
		Case{
			Message: "With zero and positive exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: true,
			},
		},
		Case{
			Message: "With zero number and negative exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: false,
			},
		},
		Case{
			Message: "With negative number and positive exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   -10,
				Exclusive: true,
			},
		},
		Case{
			Message: "With negative number and negative exclusive",
			Definition: ints.MinimumValidatorDefinition{
				Minimum:   -10,
				Exclusive: false,
			},
		},
	}
	for _, c := range cases {
		_, err := ints.NewMinimumValidator(c.Definition)
		if err != nil {
			t.Errorf("%s: %s", c.Message, err)
		}
	}
}

type MinimumValidatorTestCase struct {
	Message string
	Input   int
	Error   error
}

func TestValidateOfMinimumValidatorWithPositiveNumberAndPositiveExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   10,
		Exclusive: true,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error: &ints.MinimumValidationError{
				Input:      10,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   9,
			Error: &ints.MinimumValidationError{
				Input:      9,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithPositiveNumberAndNegativeExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   10,
		Exclusive: false,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   9,
			Error: &ints.MinimumValidationError{
				Input:      9,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("expected %v, but actual %v", c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithZeroAndPositiveExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: true,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error: &ints.MinimumValidationError{
				Input:      0.0,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -1,
			Error: &ints.MinimumValidationError{
				Input:      -1,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithZeroAndNegativeExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: false,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -1,
			Error: &ints.MinimumValidationError{
				Input:      -1,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("expected %v, but actual %v", c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithNegativeNumberAndPositiveExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   -10,
		Exclusive: true,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error: &ints.MinimumValidationError{
				Input:      -10,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -11,
			Error: &ints.MinimumValidationError{
				Input:      -11,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("Test with %s: expected %v, but actual %v", c.Message, c.Error, err)
		}
	}
}

func TestValidateOfMinimumValidatorWithNegativeNumberAndNegativeExclusive(t *testing.T) {
	def := ints.MinimumValidatorDefinition{
		Minimum:   -10,
		Exclusive: false,
	}

	v, err := ints.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -11,
			Error: &ints.MinimumValidationError{
				Input:      -11,
				Definition: def,
			},
		},
	}

	for _, c := range cases {
		if err := v.Validate(c.Input); !reflect.DeepEqual(err, c.Error) {
			t.Errorf("expected '%s', but actual '%s'", c.Error, err)
		}
	}
}
