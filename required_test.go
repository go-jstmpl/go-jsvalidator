package validator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-jstmpl/go-jsvalidator"
	"github.com/gocraft/dbr"
)

func TestNewRequiredValidator(t *testing.T) {
	type Case struct {
		Message    string
		Definition validator.RequiredValidatorDefinition
		Error      error
	}
	cases := []Case{
		{
			Message:    "single element",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo"}},
			Error:      nil,
		},
		{
			Message:    "multi elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar"}},
			Error:      nil,
		},
		{
			Message:    "empty slice",
			Definition: validator.RequiredValidatorDefinition{Required: []string{}},
			Error:      validator.RequiredDefinitionEmptyError,
		},
		{
			Message:    "duplicated elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and second",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "bar"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at first and end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "bar", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated elements at second end end",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"bar", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
		{
			Message:    "duplicated all elements",
			Definition: validator.RequiredValidatorDefinition{Required: []string{"foo", "foo", "foo"}},
			Error:      validator.RequiredDefinitionDuplicationError,
		},
	}
	for _, c := range cases {
		_, err := validator.NewRequiredValidator(c.Definition)
		if err != c.Error {
			t.Errorf("Test with %s: fail to NewPatternValidator with error %v", c.Message, err)
		}
	}
}

func TestValidateOfRequiredValidator(t *testing.T) {
	type Types struct {
		NullableStringValue dbr.NullString
		NullableIntValue    dbr.NullInt64
		NullableFloatValue  dbr.NullFloat64
		NullableBoolValue   dbr.NullBool
		NullableTimeValue   dbr.NullTime
		StringValue         string
		IntValue            int
		FloatValue          float64
		BoolValue           bool
		TimeValue           time.Time
	}

	type Input struct {
		Types      Types
		Definition validator.RequiredValidatorDefinition
	}

	type Case struct {
		Message  string
		Input    Input
		Expected error
	}

	cases := []Case{
		{
			Message: "complete struct against required",
			Input: Input{
				Types: Types{
					NullableStringValue: dbr.NewNullString("value"),
					NullableIntValue:    dbr.NewNullInt64(1),
					NullableFloatValue:  dbr.NewNullFloat64(1.1),
					NullableBoolValue:   dbr.NewNullBool(true),
					NullableTimeValue:   dbr.NewNullTime("2009-11-10 23:00:00"),
					StringValue:         "value",
					IntValue:            1,
					FloatValue:          1.1,
					BoolValue:           true,
					TimeValue:           time.Date(2009, 11, 10, 23, 00, 0, 0, time.UTC),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{
						"NullableStringValue",
						"NullableIntValue",
						"NullableFloatValue",
						"NullableBoolValue",
						"NullableTimeValue",
						"StringValue",
						"IntValue",
						"FloatValue",
						"BoolValue",
						"TimeValue",
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "NullableStringValue is missing",
			Input: Input{
				Types: Types{
					NullableStringValue: dbr.NewNullString(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableStringValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableStringValue: dbr.NewNullString(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableStringValue"},
				},
			},
		},
		{
			Message: "NullableIntValue is missing",
			Input: Input{
				Types: Types{
					NullableIntValue: dbr.NewNullInt64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableIntValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableIntValue: dbr.NewNullInt64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableIntValue"},
				},
			},
		},
		{
			Message: "NullableFloatValue is missing",
			Input: Input{
				Types: Types{
					NullableFloatValue: dbr.NewNullFloat64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableFloatValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableFloatValue: dbr.NewNullFloat64(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableFloatValue"},
				},
			},
		},
		{
			Message: "NullableBoolValue is missing",
			Input: Input{
				Types: Types{
					NullableBoolValue: dbr.NewNullBool(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableBoolValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableBoolValue: dbr.NewNullBool(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableBoolValue"},
				},
			},
		},
		{
			Message: "NullableTimeValue is missing",
			Input: Input{
				Types: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
		},

		{
			Message: "NullableTimeValue is missing",
			Input: Input{
				Types: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Types{
					NullableTimeValue: dbr.NewNullTime(nil),
				},
				Definition: validator.RequiredValidatorDefinition{
					Required: []string{"NullableTimeValue"},
				},
			},
		},
	}

	for _, c := range cases {
		definition := c.Input.Definition
		va, err := validator.NewRequiredValidator(definition)
		if err != nil {
			t.Errorf("test with %s: fail to create new required validator: %s", c.Message, err)
			continue
		}

		err = va.Validate(c.Input.Types)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("test with %s: expected %+v, but actual %+v", c.Message, c.Expected, err)
		}
	}
}

func TestConvertToConcreteValue(t *testing.T) {
	// Output expected is always Kind of non Ptr
	type Case struct {
		Message string
		Input   reflect.Value
	}
	var (
		stringValue = "string"
		intValue    = 1
		floatValue  = 1.1
		boolValue   = true
		structValue = time.Now()
	)

	cases := []Case{
		{
			Message: "kind of string",
			Input:   reflect.ValueOf(stringValue),
		},
		{
			Message: "kind of int",
			Input:   reflect.ValueOf(intValue),
		},
		{
			Message: "kind of float",
			Input:   reflect.ValueOf(floatValue),
		},
		{
			Message: "kind of bool",
			Input:   reflect.ValueOf(boolValue),
		},
		{
			Message: "kind of struct",
			Input:   reflect.ValueOf(structValue),
		},
		{
			Message: "kind of pointer of string",
			Input:   reflect.ValueOf(&stringValue),
		},
		{
			Message: "kind of pointer of int",
			Input:   reflect.ValueOf(&intValue),
		},
		{
			Message: "kind of pointer of float",
			Input:   reflect.ValueOf(&floatValue),
		},
		{
			Message: "kind of pointer of bool",
			Input:   reflect.ValueOf(&boolValue),
		},
		{
			Message: "kind of pointer of struct",
			Input:   reflect.ValueOf(&structValue),
		},
	}

	for _, c := range cases {
		v, ok := validator.ConvertToConcreteValue(c.Input)
		if !ok {
			t.Errorf("test with %s: fail to convert to concrete value %v", c.Message, c.Input)
		}
		if v.Kind() == reflect.Ptr {
			t.Errorf("test with  %s: expected non Ptr but not", c.Message)
		}
	}
}
