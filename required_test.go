package validator_test

import (
	"database/sql"
	"reflect"
	"testing"

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
	type Native struct {
		StringValue string
		IntValue    int
		FloatValue  float64
		BoolValue   bool
	}

	type PtrNative struct {
		StringValue *string
		IntValue    *int
		FloatValue  *float64
		BoolValue   *bool
	}

	type Null struct {
		StringValue dbr.NullString
		IntValue    dbr.NullInt64
		FloatValue  dbr.NullFloat64
		BoolValue   dbr.NullBool
	}

	type PtrNull struct {
		StringValue *dbr.NullString
		IntValue    *dbr.NullInt64
		FloatValue  *dbr.NullFloat64
		BoolValue   *dbr.NullBool
	}

	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	stringValue := "string value"
	intValue := 1
	floatValue := 1.1
	boolValue := true

	falsyStringValue := ""
	falsyIntValue := 0
	falsyFloatValue := 0.0
	falsyBoolValue := false

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue", "IntValue", "BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidFieldTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidFieldTypeError{
				Input:      "foo",
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of non-pointer native type of value",
			Input: Native{
				StringValue: "value",
				IntValue:    1,
				FloatValue:  1.1,
				BoolValue:   true,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer native type of falsy value",
			Input: Native{
				StringValue: "",
				IntValue:    0,
				FloatValue:  0.0,
				BoolValue:   false,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer native type of value",
			Input: PtrNative{
				StringValue: &stringValue,
				IntValue:    &intValue,
				FloatValue:  &floatValue,
				BoolValue:   &boolValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer native type of falsy value",
			Input: PtrNative{
				StringValue: &falsyStringValue,
				IntValue:    &falsyIntValue,
				FloatValue:  &falsyFloatValue,
				BoolValue:   &falsyBoolValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer native type of nil",
			Input: PtrNative{
				StringValue: nil,
				IntValue:    nil,
				FloatValue:  nil,
				BoolValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of value",
			Input: Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "value",
						Valid:  true,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 1,
						Valid:   true,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  true,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  true,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: true,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   true,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   false,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: false,
					},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: dbr.NullString{
						NullString: sql.NullString{
							String: "",
							Valid:  false,
						},
					},
					IntValue: dbr.NullInt64{
						NullInt64: sql.NullInt64{
							Int64: 0,
							Valid: false,
						},
					},
					FloatValue: dbr.NullFloat64{
						NullFloat64: sql.NullFloat64{
							Float64: 0,
							Valid:   false,
						},
					},
					BoolValue: dbr.NullBool{
						NullBool: sql.NullBool{
							Bool:  false,
							Valid: false,
						},
					},
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "value",
						Valid:  true,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 1.0,
						Valid:   true,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  true,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  true,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: true,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0.0,
						Valid:   true,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   false,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: false,
					},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &dbr.NullString{
						NullString: sql.NullString{
							String: "",
							Valid:  false,
						},
					},
					IntValue: &dbr.NullInt64{
						NullInt64: sql.NullInt64{
							Int64: 0,
							Valid: false,
						},
					},
					FloatValue: &dbr.NullFloat64{
						NullFloat64: sql.NullFloat64{
							Float64: 0,
							Valid:   false,
						},
					},
					BoolValue: &dbr.NullBool{
						NullBool: sql.NullBool{
							Bool:  false,
							Valid: false,
						},
					},
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of nil",
			Input: PtrNull{
				StringValue: nil,
				IntValue:    nil,
				FloatValue:  nil,
				BoolValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer native type of value",
			Input: &Native{
				StringValue: "value",
				IntValue:    1,
				FloatValue:  1.1,
				BoolValue:   true,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer native type of falsy value",
			Input: &Native{
				StringValue: "",
				IntValue:    0,
				FloatValue:  0.0,
				BoolValue:   false,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer native type of value",
			Input: &PtrNative{
				StringValue: &stringValue,
				IntValue:    &intValue,
				FloatValue:  &floatValue,
				BoolValue:   &boolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer native type of falsy value",
			Input: &PtrNative{
				StringValue: &falsyStringValue,
				IntValue:    &falsyIntValue,
				FloatValue:  &falsyFloatValue,
				BoolValue:   &falsyBoolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer native type of nil",
			Input: &PtrNative{
				StringValue: nil,
				IntValue:    nil,
				FloatValue:  nil,
				BoolValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "value",
						Valid:  true,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 1,
						Valid:   true,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  true,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  true,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: true,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   true,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				IntValue: dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
				},
				FloatValue: dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   false,
					},
				},
				BoolValue: dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: false,
					},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: dbr.NullString{
						NullString: sql.NullString{
							String: "",
							Valid:  false,
						},
					},
					IntValue: dbr.NullInt64{
						NullInt64: sql.NullInt64{
							Int64: 0,
							Valid: false,
						},
					},
					FloatValue: dbr.NullFloat64{
						NullFloat64: sql.NullFloat64{
							Float64: 0,
							Valid:   false,
						},
					},
					BoolValue: dbr.NullBool{
						NullBool: sql.NullBool{
							Bool:  false,
							Valid: false,
						},
					},
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "value",
						Valid:  true,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 1,
						Valid:   true,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  true,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  true,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: true,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   true,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: true,
					},
				},
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &dbr.NullString{
					NullString: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				IntValue: &dbr.NullInt64{
					NullInt64: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
				},
				FloatValue: &dbr.NullFloat64{
					NullFloat64: sql.NullFloat64{
						Float64: 0,
						Valid:   false,
					},
				},
				BoolValue: &dbr.NullBool{
					NullBool: sql.NullBool{
						Bool:  false,
						Valid: false,
					},
				},
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &dbr.NullString{
						NullString: sql.NullString{
							String: "",
							Valid:  false,
						},
					},
					IntValue: &dbr.NullInt64{
						NullInt64: sql.NullInt64{
							Int64: 0,
							Valid: false,
						},
					},
					FloatValue: &dbr.NullFloat64{
						NullFloat64: sql.NullFloat64{
							Float64: 0,
							Valid:   false,
						},
					},
					BoolValue: &dbr.NullBool{
						NullBool: sql.NullBool{
							Bool:  false,
							Valid: false,
						},
					},
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of nil",
			Input: &PtrNull{
				StringValue: nil,
				IntValue:    nil,
				FloatValue:  nil,
				BoolValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
				},
				Definition: definition,
			},
		},
	}

	va, err := validator.NewRequiredValidator(definition)
	if err != nil {
		t.Fatal("Fail to create new required validator:", err)
	}

	for _, c := range cases {
		err := va.Validate(c.Input)
		if !reflect.DeepEqual(err, c.Expected) {
			t.Errorf("Test with %s: expected %+v, but actual %+v", c.Message, c.Expected, err)
		}
	}
}
