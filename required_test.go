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

var (
	stringValue = "string value"
	intValue    = 1
	floatValue  = 1.1
	boolValue   = true
	timeValue   = time.Now()

	falsyStringValue = ""
	falsyIntValue    = 0
	falsyFloatValue  = 0.0
	falsyBoolValue   = false
	falsyTimeValue   = time.Time{}

	nullableString = dbr.NewNullString(stringValue)
	nullableInt    = dbr.NewNullInt64(intValue)
	nullableFloat  = dbr.NewNullFloat64(floatValue)
	nullableBool   = dbr.NewNullBool(boolValue)
	nullableTime   = dbr.NewNullTime(timeValue)

	falsyNullableString = dbr.NewNullString(falsyStringValue)
	falsyNullableInt    = dbr.NewNullInt64(falsyIntValue)
	falsyNullableFloat  = dbr.NewNullFloat64(falsyFloatValue)
	falsyNullableBool   = dbr.NewNullBool(falsyBoolValue)
	falsyNullableTime   = dbr.NewNullTime(falsyTimeValue)

	nullNullableString = dbr.NewNullString(nil)
	nullNullableInt    = dbr.NewNullInt64(nil)
	nullNullableFloat  = dbr.NewNullFloat64(nil)
	nullNullableBool   = dbr.NewNullBool(nil)
	nullNullableTime   = dbr.NewNullTime(nil)
)

type Native struct {
	StringValue string
	IntValue    int
	FloatValue  float64
	BoolValue   bool
	TimeValue   time.Time
}

type PtrNative struct {
	StringValue *string
	IntValue    *int
	FloatValue  *float64
	BoolValue   *bool
	TimeValue   *time.Time
}

type Null struct {
	StringValue dbr.NullString
	IntValue    dbr.NullInt64
	FloatValue  dbr.NullFloat64
	BoolValue   dbr.NullBool
	TimeValue   dbr.NullTime
}

type PtrNull struct {
	StringValue *dbr.NullString
	IntValue    *dbr.NullInt64
	FloatValue  *dbr.NullFloat64
	BoolValue   *dbr.NullBool
	TimeValue   *dbr.NullTime
}

type RequiredValidatorTestCase struct {
	Message  string
	Input    interface{}
	Expected error
}

func TestValidateOfRequiredValidatorWithString(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidTypeError{
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
func TestValidateOfRequiredValidatorWithInt(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"IntValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidTypeError{
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
func TestValidateOfRequiredValidatorWithFloat(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"FloatValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidTypeError{
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
func TestValidateOfRequiredValidatorWithBool(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidTypeError{
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
func TestValidateOfRequiredValidatorWithTime(t *testing.T) {
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "nil",
			Input:   nil,
			Expected: &validator.InvalidTypeError{
				Input:      nil,
				Definition: definition,
			},
		},
		{
			Message: "non-struct",
			Input:   "foo",
			Expected: &validator.InvalidTypeError{
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of falsy value",
			Input: Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of non-pointer nullable type of null value",
			Input: Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer struct of pointer nullable type of value",
			Input: PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of falsy value",
			Input: PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer struct of pointer nullable type of null value",
			Input: PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
				TimeValue:   timeValue,
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
				TimeValue:   timeValue,
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
				TimeValue:   &timeValue,
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
				TimeValue:   &falsyTimeValue,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNative{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer struct of non-pointer nullable type of value",
			Input: &Null{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of falsy value",
			Input: &Null{
				StringValue: falsyNullableString,
				IntValue:    falsyNullableInt,
				FloatValue:  falsyNullableFloat,
				BoolValue:   falsyNullableBool,
				TimeValue:   falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of non-pointer nullable type of null value",
			Input: &Null{
				StringValue: nullNullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer struct of pointer nullable type of value",
			Input: &PtrNull{
				StringValue: &nullableString,
				IntValue:    &nullableInt,
				FloatValue:  &nullableFloat,
				BoolValue:   &nullableBool,
				TimeValue:   &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of falsy value",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
				IntValue:    &falsyNullableInt,
				FloatValue:  &falsyNullableFloat,
				BoolValue:   &falsyNullableBool,
				TimeValue:   &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer struct of pointer nullable type of null value",
			Input: &PtrNull{
				StringValue: &nullNullableString,
				IntValue:    &nullNullableInt,
				FloatValue:  &nullNullableFloat,
				BoolValue:   &nullNullableBool,
				TimeValue:   &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
					IntValue:    &nullNullableInt,
					FloatValue:  &nullNullableFloat,
					BoolValue:   &nullNullableBool,
					TimeValue:   &nullNullableTime,
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
				TimeValue:   nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
					IntValue:    nil,
					FloatValue:  nil,
					BoolValue:   nil,
					TimeValue:   nil,
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
