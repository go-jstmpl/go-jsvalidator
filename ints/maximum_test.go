package ints_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/ints"
)

func TestNewMaximumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition ints.MaximumValidatorDefinition
	}
	cases := []Case{
		Case{
			Message: "With positive number and positive exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   10,
				Exclusive: true,
			},
		},
		Case{
			Message: "With positive number and negative exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   10,
				Exclusive: false,
			},
		},
		Case{
			Message: "With zero and positive exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   0,
				Exclusive: true,
			},
		},
		Case{
			Message: "With zero number and negative exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   0,
				Exclusive: false,
			},
		},
		Case{
			Message: "With negative number and positive exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   -10,
				Exclusive: true,
			},
		},
		Case{
			Message: "With negative number and negative exclusive",
			Definition: ints.MaximumValidatorDefinition{
				Maximum:   -10,
				Exclusive: false,
			},
		},
	}
	for _, c := range cases {
		_, err := ints.NewMaximumValidator(c.Definition)
		if err != nil {
			t.Errorf("%s: %s", c.Message, err)
		}
	}
}

type MaximumValidatorTestCase struct {
	Message string
	Input   int
	Error   error
}

func TestValidateOfMaximumValidatorWithPositiveNumberAndPositiveExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   10,
		Exclusive: true,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error: &ints.MaximumValidationError{
				Input:      10,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   11,
			Error: &ints.MaximumValidationError{
				Input:      11,
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

func TestValidateOfMaximumValidatorWithPositiveNumberAndNegativeExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   10,
		Exclusive: false,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   10,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   11,
			Error: &ints.MaximumValidationError{
				Input:      11,
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

func TestValidateOfMaximumValidatorWithZeroAndPositiveExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   0,
		Exclusive: true,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0,
			Error: &ints.MaximumValidationError{
				Input:      0,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   1,
			Error: &ints.MaximumValidationError{
				Input:      1,
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

func TestValidateOfMaximumValidatorWithZeroAndNegativeExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   0,
		Exclusive: false,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   1,
			Error: &ints.MaximumValidationError{
				Input:      1,
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

func TestValidateOfMaximumValidatorWithNegativeNumberAndPositiveExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   -10,
		Exclusive: true,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error: &ints.MaximumValidationError{
				Input:      -10,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   -9,
			Error: &ints.MaximumValidationError{
				Input:      -9,
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

func TestValidateOfMaximumValidatorWithNegativeNumberAndNegativeExclusive(t *testing.T) {
	def := ints.MaximumValidatorDefinition{
		Maximum:   -10,
		Exclusive: false,
	}

	v, err := ints.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -11,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -10,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   -9,
			Error: &ints.MaximumValidationError{
				Input:      -9,
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
