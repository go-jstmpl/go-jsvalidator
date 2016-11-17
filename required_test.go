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

func TestValidateOfRequiredValidatorWithPrimitiveString(t *testing.T) {
	type Primitive struct {
		StringValue string
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Primitive struct with value of string",
			Input: Primitive{
				StringValue: stringValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Primitive struct with value of falsy string",
			Input: Primitive{
				StringValue: falsyStringValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of string",
			Input: &Primitive{
				StringValue: stringValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of falsy string",
			Input: &Primitive{
				StringValue: falsyStringValue,
			},
			Expected: nil,
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
func TestValidateOfRequiredValidatorWithPrimitiveInt(t *testing.T) {
	type Primitive struct {
		IntValue int
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"IntValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Primitive struct with value of int",
			Input: Primitive{
				IntValue: intValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Primitive struct with value of falsy int",
			Input: Primitive{
				IntValue: falsyIntValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of int",
			Input: &Primitive{
				IntValue: intValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of falsy int",
			Input: &Primitive{
				IntValue: falsyIntValue,
			},
			Expected: nil,
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
func TestValidateOfRequiredValidatorWithPrimitiveFloat(t *testing.T) {
	type Primitive struct {
		FloatValue float64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"FloatValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Primitive struct with value of float",
			Input: Primitive{
				FloatValue: floatValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Primitive struct with value of falsy float",
			Input: Primitive{
				FloatValue: falsyFloatValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of float",
			Input: &Primitive{
				FloatValue: floatValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of falsy float",
			Input: &Primitive{
				FloatValue: falsyFloatValue,
			},
			Expected: nil,
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
func TestValidateOfRequiredValidatorWithPrimitiveBool(t *testing.T) {
	type Primitive struct {
		BoolValue bool
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Primitive struct with value of bool",
			Input: Primitive{
				BoolValue: boolValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Primitive struct with value of falsy bool",
			Input: Primitive{
				BoolValue: falsyBoolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of bool",
			Input: &Primitive{
				BoolValue: boolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of falsy bool",
			Input: &Primitive{
				BoolValue: falsyBoolValue,
			},
			Expected: nil,
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
func TestValidateOfRequiredValidatorWithPrimitiveTime(t *testing.T) {
	type Primitive struct {
		TimeValue time.Time
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}
	definition := validator.RequiredValidatorDefinition{
		Required: []string{"TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Primitive struct with value of time",
			Input: Primitive{
				TimeValue: timeValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Primitive struct with value of falsy time",
			Input: Primitive{
				TimeValue: falsyTimeValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of time",
			Input: &Primitive{
				TimeValue: timeValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer Primitive struct with value of falsy time",
			Input: &Primitive{
				TimeValue: falsyTimeValue,
			},
			Expected: nil,
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

func TestValidateOfRequiredValidatorWithPtrPrimitiveString(t *testing.T) {
	type PtrPrimitive struct {
		StringValue *string
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrPrimitive struct with value of string",
			Input: PtrPrimitive{
				StringValue: &stringValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with value of falsy string",
			Input: PtrPrimitive{
				StringValue: &falsyStringValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with nil",
			Input: PtrPrimitive{
				StringValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrPrimitive{
					StringValue: nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrPrimitive struct with value of string",
			Input: &PtrPrimitive{
				StringValue: &stringValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with value of falsy string",
			Input: &PtrPrimitive{
				StringValue: &falsyStringValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with nil",
			Input: &PtrPrimitive{
				StringValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrPrimitive{
					StringValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrPrimitiveInt(t *testing.T) {
	type PtrPrimitive struct {
		IntValue *int
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"IntValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrPrimitive struct with value of int",
			Input: PtrPrimitive{
				IntValue: &intValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with value of falsy int",
			Input: PtrPrimitive{
				IntValue: &falsyIntValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with nil",
			Input: PtrPrimitive{
				IntValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrPrimitive{
					IntValue: nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrPrimitive struct with value of int",
			Input: &PtrPrimitive{
				IntValue: &intValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with value of falsy int",
			Input: &PtrPrimitive{
				IntValue: &falsyIntValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with nil",
			Input: &PtrPrimitive{
				IntValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrPrimitive{
					IntValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrPrimitiveFloat(t *testing.T) {
	type PtrPrimitive struct {
		FloatValue *float64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"FloatValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrPrimitive struct with value of float",
			Input: PtrPrimitive{
				FloatValue: &floatValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with value of falsy float",
			Input: PtrPrimitive{
				FloatValue: &falsyFloatValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with nil",
			Input: PtrPrimitive{
				FloatValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrPrimitive{
					FloatValue: nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrPrimitive struct with value of float",
			Input: &PtrPrimitive{
				FloatValue: &floatValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with value of falsy float",
			Input: &PtrPrimitive{
				FloatValue: &falsyFloatValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with nil",
			Input: &PtrPrimitive{
				FloatValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrPrimitive{
					FloatValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrPrimitiveBoll(t *testing.T) {
	type PtrPrimitive struct {
		BoolValue *bool
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrPrimitive struct with value of bool",
			Input: PtrPrimitive{
				BoolValue: &boolValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with value of falsy bool",
			Input: PtrPrimitive{
				BoolValue: &falsyBoolValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with nil",
			Input: PtrPrimitive{
				BoolValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrPrimitive{
					BoolValue: nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrPrimitive struct with value of bool",
			Input: &PtrPrimitive{
				BoolValue: &boolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with value of falsy bool",
			Input: &PtrPrimitive{
				BoolValue: &falsyBoolValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with nil",
			Input: &PtrPrimitive{
				BoolValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrPrimitive{
					BoolValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrPrimitiveTime(t *testing.T) {
	type PtrPrimitive struct {
		TimeValue *time.Time
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrPrimitive struct with value of time",
			Input: PtrPrimitive{
				TimeValue: &timeValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with value of falsy time",
			Input: PtrPrimitive{
				TimeValue: &falsyTimeValue,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrPrimitive struct with nil",
			Input: PtrPrimitive{
				TimeValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrPrimitive{
					TimeValue: nil,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrPrimitive struct with value of time",
			Input: &PtrPrimitive{
				TimeValue: &timeValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with value of falsy time",
			Input: &PtrPrimitive{
				TimeValue: &falsyTimeValue,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrPrimitive struct with nil",
			Input: &PtrPrimitive{
				TimeValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrPrimitive{
					TimeValue: nil,
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

func TestValidateOfRequiredValidatorWithNullableString(t *testing.T) {
	type Null struct {
		StringValue dbr.NullString
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Null struct with value of string",
			Input: Null{
				StringValue: nullableString,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of falsy string",
			Input: Null{
				StringValue: falsyNullableString,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of null",
			Input: Null{
				StringValue: nullNullableString,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					StringValue: nullNullableString,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer Null struct with value of string",
			Input: &Null{
				StringValue: nullableString,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of falsy string",
			Input: &Null{
				StringValue: falsyNullableString,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of null",
			Input: &Null{
				StringValue: nullNullableString,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					StringValue: nullNullableString,
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
func TestValidateOfRequiredValidatorWithNullableInt(t *testing.T) {
	type Null struct {
		IntValue dbr.NullInt64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"IntValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Null struct with value of int",
			Input: Null{
				IntValue: nullableInt,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of falsy int",
			Input: Null{
				IntValue: falsyNullableInt,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of null",
			Input: Null{
				IntValue: nullNullableInt,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					IntValue: nullNullableInt,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer Null struct with value of int",
			Input: &Null{
				IntValue: nullableInt,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of falsy int",
			Input: &Null{
				IntValue: falsyNullableInt,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of null",
			Input: &Null{
				IntValue: nullNullableInt,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					IntValue: nullNullableInt,
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
func TestValidateOfRequiredValidatorWithNullableFloat(t *testing.T) {
	type Null struct {
		FloatValue dbr.NullFloat64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"FloatValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Null struct with value of float",
			Input: Null{
				FloatValue: nullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of falsy float",
			Input: Null{
				FloatValue: falsyNullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of null",
			Input: Null{
				FloatValue: nullNullableFloat,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					FloatValue: nullNullableFloat,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer Null struct with value of float",
			Input: &Null{
				FloatValue: nullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of falsy float",
			Input: &Null{
				FloatValue: falsyNullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of null",
			Input: &Null{
				FloatValue: nullNullableFloat,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					FloatValue: nullNullableFloat,
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
func TestValidateOfRequiredValidatorWithNullableBool(t *testing.T) {
	type Null struct {
		BoolValue dbr.NullBool
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Null struct with value of bool",
			Input: Null{
				BoolValue: nullableBool,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of falsy bool",
			Input: Null{
				BoolValue: falsyNullableBool,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of null",
			Input: Null{
				BoolValue: nullNullableBool,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					BoolValue: nullNullableBool,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer Null struct with value of bool",
			Input: &Null{
				BoolValue: nullableBool,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of falsy bool",
			Input: &Null{
				BoolValue: falsyNullableBool,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of null",
			Input: &Null{
				BoolValue: nullNullableBool,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					BoolValue: nullNullableBool,
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
func TestValidateOfRequiredValidatorWithNullableTime(t *testing.T) {
	type Null struct {
		TimeValue dbr.NullTime
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer Null struct with value of time",
			Input: Null{
				TimeValue: nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of falsy time",
			Input: Null{
				TimeValue: falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer Null struct with value of null",
			Input: Null{
				TimeValue: nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Null{
					TimeValue: nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer Null struct with value of time",
			Input: &Null{
				TimeValue: nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of falsy time",
			Input: &Null{
				TimeValue: falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer Null struct with value of null",
			Input: &Null{
				TimeValue: nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &Null{
					TimeValue: nullNullableTime,
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

func TestValidateOfRequiredValidatorWithPtrNullableString(t *testing.T) {
	type PtrNull struct {
		StringValue *dbr.NullString
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrNull struct with value of string",
			Input: PtrNull{
				StringValue: &nullableString,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of falsy string",
			Input: PtrNull{
				StringValue: &falsyNullableString,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of null string",
			Input: PtrNull{
				StringValue: &nullNullableString,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: &nullNullableString,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer PtrNull struct with pointer string of nil",
			Input: PtrNull{
				StringValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					StringValue: nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer PtrNull struct with value of string",
			Input: &PtrNull{
				StringValue: &nullableString,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of falsy string",
			Input: &PtrNull{
				StringValue: &falsyNullableString,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of string",
			Input: &PtrNull{
				StringValue: &nullNullableString,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: &nullNullableString,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrNull struct with pointer string of nil",
			Input: &PtrNull{
				StringValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					StringValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrNullableInt(t *testing.T) {
	type PtrNull struct {
		IntValue *dbr.NullInt64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"IntValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrNull struct with value of int",
			Input: PtrNull{
				IntValue: &nullableInt,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of falsy int",
			Input: PtrNull{
				IntValue: &falsyNullableInt,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of null int",
			Input: PtrNull{
				IntValue: &nullNullableInt,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					IntValue: &nullNullableInt,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer PtrNull struct with pointer int of nil",
			Input: PtrNull{
				IntValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					IntValue: nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer PtrNull struct with value of int",
			Input: &PtrNull{
				IntValue: &nullableInt,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of falsy int",
			Input: &PtrNull{
				IntValue: &falsyNullableInt,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of int",
			Input: &PtrNull{
				IntValue: &nullNullableInt,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					IntValue: &nullNullableInt,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrNull struct with pointer int of nil",
			Input: &PtrNull{
				IntValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					IntValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrNullableFloat(t *testing.T) {
	type PtrNull struct {
		FloatValue *dbr.NullFloat64
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"FloatValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrNull struct with value of float",
			Input: PtrNull{
				FloatValue: &nullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of falsy float",
			Input: PtrNull{
				FloatValue: &falsyNullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of null float",
			Input: PtrNull{
				FloatValue: &nullNullableFloat,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					FloatValue: &nullNullableFloat,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer PtrNull struct with pointer float of nil",
			Input: PtrNull{
				FloatValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					FloatValue: nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer PtrNull struct with value of float",
			Input: &PtrNull{
				FloatValue: &nullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of falsy float",
			Input: &PtrNull{
				FloatValue: &falsyNullableFloat,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of float",
			Input: &PtrNull{
				FloatValue: &nullNullableFloat,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					FloatValue: &nullNullableFloat,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrNull struct with pointer float of nil",
			Input: &PtrNull{
				FloatValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					FloatValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrNullableBool(t *testing.T) {
	type PtrNull struct {
		BoolValue *dbr.NullBool
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"BoolValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrNull struct with value of bool",
			Input: PtrNull{
				BoolValue: &nullableBool,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of falsy bool",
			Input: PtrNull{
				BoolValue: &falsyNullableBool,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of null bool",
			Input: PtrNull{
				BoolValue: &nullNullableBool,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					BoolValue: &nullNullableBool,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer PtrNull struct with pointer bool of nil",
			Input: PtrNull{
				BoolValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					BoolValue: nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer PtrNull struct with value of bool",
			Input: &PtrNull{
				BoolValue: &nullableBool,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of falsy bool",
			Input: &PtrNull{
				BoolValue: &falsyNullableBool,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of bool",
			Input: &PtrNull{
				BoolValue: &nullNullableBool,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					BoolValue: &nullNullableBool,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrNull struct with pointer bool of nil",
			Input: &PtrNull{
				BoolValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					BoolValue: nil,
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
func TestValidateOfRequiredValidatorWithPtrNullableTime(t *testing.T) {
	type PtrNull struct {
		TimeValue *dbr.NullTime
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "non-pointer PtrNull struct with value of time",
			Input: PtrNull{
				TimeValue: &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of falsy time",
			Input: PtrNull{
				TimeValue: &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "non-pointer PtrNull struct with value of null time",
			Input: PtrNull{
				TimeValue: &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					TimeValue: &nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "non-pointer PtrNull struct with pointer time of nil",
			Input: PtrNull{
				TimeValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: PtrNull{
					TimeValue: nil,
				},
				Definition: definition,
			},
		},

		{
			Message: "pointer PtrNull struct with value of time",
			Input: &PtrNull{
				TimeValue: &nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of falsy time",
			Input: &PtrNull{
				TimeValue: &falsyNullableTime,
			},
			Expected: nil,
		},
		{
			Message: "pointer PtrNull struct with value of time",
			Input: &PtrNull{
				TimeValue: &nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					TimeValue: &nullNullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "pointer PtrNull struct with pointer time of nil",
			Input: &PtrNull{
				TimeValue: nil,
			},
			Expected: &validator.RequiredValidationError{
				Input: &PtrNull{
					TimeValue: nil,
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

func TestValidateOfRequiredValidatorWithMultiField(t *testing.T) {
	type Nullable struct {
		StringValue dbr.NullString
		IntValue    dbr.NullInt64
		FloatValue  dbr.NullFloat64
		BoolValue   dbr.NullBool
		TimeValue   dbr.NullTime
	}
	type RequiredValidatorTestCase struct {
		Message  string
		Input    interface{}
		Expected error
	}

	definition := validator.RequiredValidatorDefinition{
		Required: []string{"StringValue", "IntValue", "FloatValue", "BoolValue", "TimeValue"},
	}

	cases := []RequiredValidatorTestCase{
		{
			Message: "complete struct against required",
			Input: Nullable{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: nil,
		},
		{
			Message: "StringValue is missing",
			Input: Nullable{
				StringValue: nullNullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Nullable{
					StringValue: nullNullableString,
					IntValue:    nullableInt,
					FloatValue:  nullableFloat,
					BoolValue:   nullableBool,
					TimeValue:   nullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "InvValue is missing",
			Input: Nullable{
				StringValue: nullableString,
				IntValue:    nullNullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Nullable{
					StringValue: nullableString,
					IntValue:    nullNullableInt,
					FloatValue:  nullableFloat,
					BoolValue:   nullableBool,
					TimeValue:   nullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "FloatValue is missing",
			Input: Nullable{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullNullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Nullable{
					StringValue: nullableString,
					IntValue:    nullableInt,
					FloatValue:  nullNullableFloat,
					BoolValue:   nullableBool,
					TimeValue:   nullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "BoolValue is missing",
			Input: Nullable{
				StringValue: nullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullNullableBool,
				TimeValue:   nullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Nullable{
					StringValue: nullableString,
					IntValue:    nullableInt,
					FloatValue:  nullableFloat,
					BoolValue:   nullNullableBool,
					TimeValue:   nullableTime,
				},
				Definition: definition,
			},
		},
		{
			Message: "TimeValue is missing",
			Input: Nullable{
				StringValue: nullNullableString,
				IntValue:    nullableInt,
				FloatValue:  nullableFloat,
				BoolValue:   nullableBool,
				TimeValue:   nullNullableTime,
			},
			Expected: &validator.RequiredValidationError{
				Input: Nullable{
					StringValue: nullNullableString,
					IntValue:    nullableInt,
					FloatValue:  nullableFloat,
					BoolValue:   nullableBool,
					TimeValue:   nullNullableTime,
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
