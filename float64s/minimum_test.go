package float64s_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jsvalidator/float64s"
)

func TestNewMinimumValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition float64s.MinimumValidatorDefinition
	}
	cases := []Case{
		{
			Message: "positive number and positive exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   1.0,
				Exclusive: true,
			},
		},
		{
			Message: "positive number and negative exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   1.0,
				Exclusive: false,
			},
		},
		{
			Message: "zero and positive exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: true,
			},
		},
		{
			Message: "zero number and negative exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   0.0,
				Exclusive: false,
			},
		},
		{
			Message: "negative number and positive exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   -1.0,
				Exclusive: true,
			},
		},
		{
			Message: "negative number and negative exclusive",
			Definition: float64s.MinimumValidatorDefinition{
				Minimum:   -1.0,
				Exclusive: false,
			},
		},
	}
	for _, c := range cases {
		_, err := float64s.NewMinimumValidator(c.Definition)
		if err != nil {
			t.Errorf("Test with %s: fail to NewMinimumValidator with error %v", c.Message, err)
		}
	}
}

type MinimumValidatorTestCase struct {
	Message string
	Input   float64
	Error   error
}

func TestValidateOfMinimumValidatorWithPositiveNumberAndPositiveExclusive(t *testing.T) {
	def := float64s.MinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: true,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   1.0,
			Error: &float64s.MinimumValidationError{
				Input:      1.0,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   0.9,
			Error: &float64s.MinimumValidationError{
				Input:      0.9,
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
	def := float64s.MinimumValidatorDefinition{
		Minimum:   1.0,
		Exclusive: false,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   1.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   1.0,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   0.9,
			Error: &float64s.MinimumValidationError{
				Input:      0.9,
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

func TestValidateOfMinimumValidatorWithZeroAndPositiveExclusive(t *testing.T) {
	def := float64s.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: true,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   0.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error: &float64s.MinimumValidationError{
				Input:      0.0,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -0.1,
			Error: &float64s.MinimumValidationError{
				Input:      -0.1,
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
	def := float64s.MinimumValidatorDefinition{
		Minimum:   0.0,
		Exclusive: false,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   0.1,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   0.0,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -0.1,
			Error: &float64s.MinimumValidationError{
				Input:      -0.1,
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

func TestValidateOfMinimumValidatorWithNegativeNumberAndPositiveExclusive(t *testing.T) {
	def := float64s.MinimumValidatorDefinition{
		Minimum:   -1.0,
		Exclusive: true,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Fatalf("Fail to NewMinimumValidator: %s", err)
	}

	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -0.9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -1.0,
			Error: &float64s.MinimumValidationError{
				Input:      -1.0,
				Definition: def,
			},
		},
		{
			Message: "less number",
			Input:   -1.1,
			Error: &float64s.MinimumValidationError{
				Input:      -1.1,
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
	def := float64s.MinimumValidatorDefinition{
		Minimum:   -1.0,
		Exclusive: false,
	}

	v, err := float64s.NewMinimumValidator(def)
	if err != nil {
		t.Error(err.Error())
	}
	cases := []MinimumValidatorTestCase{
		{
			Message: "greater number",
			Input:   -0.9,
			Error:   nil,
		},
		{
			Message: "same number",
			Input:   -1.0,
			Error:   nil,
		},
		{
			Message: "less number",
			Input:   -1.1,
			Error: &float64s.MinimumValidationError{
				Input:      -1.1,
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
