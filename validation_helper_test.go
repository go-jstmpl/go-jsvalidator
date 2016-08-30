package validator

import (
	"testing"
)

//There structs for make constructer test case
type TestInterface interface {
	Validate(string) bool
	Error(string) string
}

type TestCase struct {
	Case    TestInterface
	Message string
	Pass    bool
}

type IntEnumTestCase struct {
	Case    IntEnumValidator
	Message string
	Pass    bool
}

type StringEnumTestCase struct {
	Case    StringEnumValidator
	Message string
	Pass    bool
}

//This struct for make validator test case
type ValidatorTestCase struct {
	Case    interface{}
	Message string
	Pass    bool
}

// The purpose of method with a name like 'Test"Validation Keyword"' is testing constructor
// The purpose of method with a name like 'Test"Validation Keyword"Validator' is testing validator

func TestMaximumValidator(t *testing.T) {
	//Case exclusive is false
	validator, err := NewMaximumValidator(100, false)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []ValidatorTestCase{{
		Case:    99,
		Message: "'99' is less than Maximum therefore should be pass",
		Pass:    true,
	}, {
		Case:    100,
		Message: "'100' equal to Maximum and exclusive is false in this case should be pass",
		Pass:    true,
	}, {
		Case:    101,
		Message: "'101' is greater than Maximum therefore should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(int))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}

	//Case exclusive is true
	validator, err = NewMaximumValidator(100, true)
	if err != nil {
		t.Error(err.Error())
	}

	tests = []ValidatorTestCase{{
		Case:    99,
		Message: "'99' is less than Maximum therefore should be pass",
		Pass:    true,
	}, {
		Case:    100,
		Message: "'100' equal to Maximum and exclusive is false in this case should not be pass",
		Pass:    false,
	}, {
		Case:    101,
		Message: "'101' is greater than Maximum therefore should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(int))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}

func TestMinimumValidator(t *testing.T) {
	//Case exclusive is false
	validator, err := NewMinimumValidator(100, false)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []ValidatorTestCase{{
		Case:    101,
		Message: "'101' is greater than Minimum therefore should be pass",
		Pass:    true,
	}, {
		Case:    100,
		Message: "'100' equal to Minimum and exclusive is false in this case should be pass",
		Pass:    true,
	}, {
		Case:    99,
		Message: "'99' is less than Minimum therefore should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(int))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}

	//Case exclusive is true
	validator, err = NewMinimumValidator(100, true)
	if err != nil {
		t.Error(err.Error())
	}

	tests = []ValidatorTestCase{{
		Case:    101,
		Message: "'101' is greater than Minimum therefore should be pass",
		Pass:    true,
	}, {
		Case:    100,
		Message: "'100' equal to Minimum and exclusive is false in this case should not be pass",
		Pass:    false,
	}, {
		Case:    99,
		Message: "'99' is less than Minimum therefore should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(int))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}
func TestMaxLength(t *testing.T) {
	tests := []TestCase{{
		Case:    MaxLengthValidator{MaxLength: -1},
		Message: "MaxLength should be greater than 0",
		Pass:    false,
	}, {
		Case:    MaxLengthValidator{MaxLength: 0},
		Message: "MaxLength 0 should not be pass",
		Pass:    false,
	}, {
		Case:    MaxLengthValidator{MaxLength: 1},
		Message: "MaxLength 1 should be pass",
		Pass:    true,
	}}

	for _, test := range tests {
		_, err := NewMaxLengthValidator(test.Case.(MaxLengthValidator).MaxLength)
		//if Pass flag true, err should be empty
		if test.Pass == true && err != nil {
			t.Error(test.Message)
		}

		//if Pass flag false, err should not be empty
		if test.Pass == false && err == nil {
			t.Error(test.Message)
		}
	}

}

func TestMaxLengthValidator(t *testing.T) {
	validator, err := NewMaxLengthValidator(5)
	if err != nil {
		t.Error(err.Error())
	}

	tests := []ValidatorTestCase{{
		Case:    "あいうえ",
		Message: "4 charactors is less than maxLength therefore should be pass",
		Pass:    true,
	}, {
		Case:    "あいうえお",
		Message: "case value equal to maxLength should be pass",
		Pass:    true,
	}, {
		Case:    "あいうえおか",
		Message: "6 charactors is greater than maxLength therefore should not be pass",
		Pass:    false,
	}, {
		Case:    "abcde",
		Message: "case 5 characters in en should be pass",
		Pass:    true,
	}, {
		Case:    "",
		Message: "empty string should be pass",
		Pass:    true,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(string))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}

func TestMinLength(t *testing.T) {
	tests := []TestCase{{
		Case:    MinLengthValidator{MinLength: -1},
		Message: "MinLength should be greater than 0",
		Pass:    false,
	}, {
		Case:    MinLengthValidator{MinLength: 0},
		Message: "MinLength 0 should not be pass",
		Pass:    false,
	}, {
		Case:    MinLengthValidator{MinLength: 1},
		Message: "MinLength 1 should be pass",
		Pass:    true,
	}}

	for _, test := range tests {
		_, err := NewMinLengthValidator(test.Case.(MinLengthValidator).MinLength)
		//if Pass flag true, err should be empty
		if test.Pass == true && err != nil {
			t.Error(test.Message)
		}

		//if Pass flag false, err should not be empty
		if test.Pass == false && err == nil {
			t.Error(test.Message)
		}
	}
}

func TestMinLengthValidator(t *testing.T) {
	validator, err := NewMinLengthValidator(5)
	if err != nil {
		t.Error(err.Error())
	}
	tests := []ValidatorTestCase{{
		Case:    "あいうえおか",
		Message: "6 charactors is greater than minLength therefore should be pass",
		Pass:    true,
	}, {
		Case:    "あいうえお",
		Message: "case value equal to minLength should be pass",
		Pass:    true,
	}, {

		Case:    "あいうえ",
		Message: "4` charactors less than minLength therefore should not be pass",
		Pass:    false,
	}, {
		Case:    "abcde",
		Message: "case 5 characters in en should be pass",
		Pass:    true,
	}, {
		Case:    "あiうeお",
		Message: "case that includes japanese and english should be pass",
		Pass:    true,
	}, {
		Case:    "",
		Message: "empty string should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(string))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}

func TestPattern(t *testing.T) {
	tests := []TestCase{{
		Case:    PatternValidator{Pattern: ""},
		Message: "pattern should be empty",
		Pass:    true,
	}, {
		Case:    PatternValidator{Pattern: "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$"},
		Message: "pattern `^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$` should be pass regex compile",
		Pass:    true,
	}, {
		Case:    PatternValidator{Pattern: "[a-z"},
		Message: "pattern `[a-z` should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		_, err := NewPatternValidator(test.Case.(PatternValidator).Pattern)
		//if Pass flag true, err should be empty
		if test.Pass == true && err != nil {
			t.Error(test.Message)
		}
		//if Pass flag false, err should not be empty
		if test.Pass == false && err == nil {
			t.Error(test.Message)
		}
	}
}

func TestPatternValidator(t *testing.T) {
	validator, err := NewPatternValidator("^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$")
	if err != nil {
		t.Error(err.Error())
	}

	tests := []ValidatorTestCase{{
		Case:    "555-1212",
		Message: "should be pass",
		Pass:    true,
	}, {
		Case:    "(888)555-1212",
		Message: "should be pass",
		Pass:    true,
	}, {
		Case:    "(888)555-1212 ext. 532",
		Message: "should not be pass",
		Pass:    false,
	}, {
		Case:    "(800)FLOWERS",
		Message: "should not be pass",
		Pass:    false,
	}}
	for _, test := range tests {
		ok := validator.Validate(test.Case.(string))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}

func TestIntEnum(t *testing.T) {
	tests := []IntEnumTestCase{{
		Case:    IntEnumValidator{Enumerate: []int{}},
		Message: "enum accepts list shoud not be empty",
		Pass:    false,
	}, {
		Case:    IntEnumValidator{Enumerate: []int{10, 20, 10}},
		Message: "enum accepts element should be unique",
		Pass:    false,
	}, {
		Case:    IntEnumValidator{Enumerate: []int{10, 20, 30}},
		Message: "enum accepts list [10, 20, 30] should be pass",
		Pass:    true,
	}}

	for _, test := range tests {
		_, err := NewIntEnumValidator(test.Case.Enumerate)
		//if Pass flag true, err should be empty
		if test.Pass == true && err != nil {
			t.Error(test.Message)
		}

		//if Pass flag false, err should not be empty
		if test.Pass == false && err == nil {
			t.Error(test.Message)
		}
	}
}

func TestIntEnumvalidator(t *testing.T) {
	validator, err := NewIntEnumValidator([]int{11, 13, 17, 19})
	if err != nil {
		t.Error(err.Error())
	}

	tests := []ValidatorTestCase{{
		Case:    13,
		Message: "There is '13' in enumerates therefore should be pass",
		Pass:    true,
	}, {
		Case:    14,
		Message: "There is not '14' in enumerates therefore should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(int))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}

func TestStringEnum(t *testing.T) {
	tests := []StringEnumTestCase{{
		Case:    StringEnumValidator{Enumerate: []string{}},
		Message: "enum accepts list shoudl not be empty",
		Pass:    false,
	}, {
		Case:    StringEnumValidator{Enumerate: []string{"test", "test"}},
		Message: "enum accepts element should be unique",
		Pass:    false,
	}, {
		Case:    StringEnumValidator{Enumerate: []string{"test1", "test2", "test3"}},
		Message: "enum accepts list [test1, test2, test3] should be pass",
		Pass:    true,
	}}

	for _, test := range tests {
		_, err := NewStringEnumValidator(test.Case.Enumerate)
		// if Pass flag true, err should be empty
		if test.Pass == true && err != nil {
			t.Error(test.Message)
		}

		// if Pass flag false, err should not be empty
		if test.Pass != true && err == nil {
			t.Error(test.Message)
		}
	}
}

func TestStringEnumValidator(t *testing.T) {
	validator, err := NewStringEnumValidator([]string{"red", "amgber", "blue"})
	if err != nil {
		t.Error(err.Error())
	}

	tests := []ValidatorTestCase{{
		Case:    "red",
		Message: "There is `red` in enumerates therefore should be pass",
		Pass:    true,
	}, {
		Case:    "green",
		Message: "There is not 'green' in enumerates therefore should not be pass",
		Pass:    false,
	}, {
		Case:    "",
		Message: "empty string case should not be pass",
		Pass:    false,
	}}

	for _, test := range tests {
		ok := validator.Validate(test.Case.(string))
		//if Pass flag true, ok should be true
		if test.Pass == true && ok != true {
			t.Error(test.Message)
		}

		//if Pass flag false, ok should not be true
		if test.Pass == false && ok == true {
			t.Error(test.Message)
		}
	}
}
