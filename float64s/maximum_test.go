package float64s_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/float64s"
)

func TestNewMaximumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition float64s.MaximumValidatorDefinition
	}
	cases := []Case{
		{
			Message: "positive number and positive exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   1.0,
				Exclusive: true,
			},
		},
		{
			Message: "positive number and negative exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   1.0,
				Exclusive: false,
			},
		},
		{
			Message: "zero and positive exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   0.0,
				Exclusive: true,
			},
		},
		{
			Message: "zero number and negative exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   0.0,
				Exclusive: false,
			},
		},
		{
			Message: "negative number and positive exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   -1.0,
				Exclusive: true,
			},
		},
		{
			Message: "negative number and negative exclusive",
			Definition: float64s.MaximumValidatorDefinition{
				Maximum:   -1.0,
				Exclusive: false,
			},
		},
	}
	for _, c := range cases {
		_, err := float64s.NewMaximumValidator(c.Definition)
		if err != nil {
			t.Errorf("Test with %s: fail to NewMaximumValidator with error %v", c.Message, err)
		}
	}
}

type MaximumValidatorTestCase struct {
	Message string
	Input   float64
	Error   error
}

func TestValidateOfMaximumValidatorWithPositiveNumberAndPositiveExclusive(t *testing.T) {
	def := float64s.MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: true,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   0.9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   1.0,
			Error: &float64s.MaximumValidationError{
				Input:      1.0,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   1.1,
			Error: &float64s.MaximumValidationError{
				Input:      1.1,
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
	def := float64s.MaximumValidatorDefinition{
		Maximum:   1.0,
		Exclusive: false,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   0.9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   1.0,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   1.1,
			Error: &float64s.MaximumValidationError{
				Input:      1.1,
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

func TestValidateOfMaximumValidatorWithZeroAndPositiveExclusive(t *testing.T) {
	def := float64s.MaximumValidatorDefinition{
		Maximum:   0.0,
		Exclusive: true,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -0.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error: &float64s.MaximumValidationError{
				Input:      0.0,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   0.1,
			Error: &float64s.MaximumValidationError{
				Input:      0.1,
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
	def := float64s.MaximumValidatorDefinition{
		Maximum:   0.0,
		Exclusive: false,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -0.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   0.1,
			Error: &float64s.MaximumValidationError{
				Input:      0.1,
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

func TestValidateOfMaximumValidatorWithNegativeNumberAndPositiveExclusive(t *testing.T) {
	def := float64s.MaximumValidatorDefinition{
		Maximum:   -1.0,
		Exclusive: true,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMaximumValidator: %s", err)
	}

	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -1.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -1.0,
			Error: &float64s.MaximumValidationError{
				Input:      -1.0,
				Definition: def,
			},
		},
		{
			Message: "greater number",
			Input:   -0.9,
			Error: &float64s.MaximumValidationError{
				Input:      -0.9,
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
	def := float64s.MaximumValidatorDefinition{
		Maximum:   -1.0,
		Exclusive: false,
	}

	v, err := float64s.NewMaximumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MaximumValidatorTestCase{
		{
			Message: "less number",
			Input:   -1.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -1.0,
			Error:   nil,
		},
		{
			Message: "greater number",
			Input:   -0.9,
			Error: &float64s.MaximumValidationError{
				Input:      -0.9,
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
