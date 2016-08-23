package validator

import (
	"fmt"
	"regexp"
)

type MaximumError struct {
	Expected  int  `json:"expected"`
	Actual    int  `json:"actual"`
	Exclusive bool `json:"exclusive"`
}

func (m MaximumError) Error() string {
	return fmt.Sprintf("expected less than %d but actual %d with option exlusive %b", m.Expected, m.Actual, m.Exclusive)
}

type MinimumError struct {
	Expected  int  `json:"expected"`
	Actual    int  `json:"actual"`
	Exclusive bool `json:"exclusive"`
}

func (m MinimumError) Error() string {
	return fmt.Sprintf("expected more than %d but actual %d with option exlusive %b", m.Expected, m.Actual, m.Exclusive)
}

type MaxLengthError struct {
	Expected int `json:"expected"`
	Actual   int `json:"actual"`
}

func (m MaxLengthError) Error() string {
	return fmt.Sprintf("expected less than %d characters but actual %d characters", m.Expected, m.Actual)
}

type MinLengthError struct {
	Expected int `json:"expected"`
	Actual   int `json:"actual"`
}

func (m MinLengthError) Error() string {
	return fmt.Sprintf("expected more than %d characters but actual %d characters", m.Expected, m.Actual)
}

type PatternError struct {
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

func (m PatternError) Error() string {
	return fmt.Sprintf("expected pattern of regex `%d` but actual %d", m.Expected, m.Actual)
}

type EnumError struct {
	Expected []string `json:"expected"`
	Actual   string   `json:"actual"`
}

func (m EnumError) Error() string {
	return fmt.Sprintf("", m.Expected, m.Actual)
}

//if "exclusiveMaximum" is not present, or has boolean value false, then the instance is valid if it is lower than, or equal to, the value of "maximum";
//if "exclusiveMaximum" has boolean value true, the instance is valid if it is strictly lower than the value of "maximum".
func Maximum(val, max int, exc bool) error {
	if exc {
		if val >= max {
			return MaximumError{max, val, exc}
		}
		return nil
	}

	if val > max {
		return MaximumError{max, val, exc}
	}

	return nil
}

//if "exclusiveMinimum" is not present, or has boolean value false, then the instance is valid if it is greater than, or equal to, the value of "minimum";
//if "exclusiveMinimum" is present and has boolean value true, the instance is valid if it is strictly greater than the value of "minimum".
func Minimum(val, min int, exc bool) error {
	if exc {
		if val <= min {
			return MinimumError{min, val, exc}
		}
		return nil
	}

	if val < min {
		return MinimumError{min, val, exc}
	}

	return nil
}

func MaxLength(val string, max int) error {
	if max < 0 {
		return fmt.Errorf("maxLength must be greater than, or equal to, 0. maxLength: %d", max)
	}

	c := len(val)
	if c > max {
		return MaxLengthError{max, c}
	}

	return nil
}

func MinLength(val string, min int) error {
	if min < 0 {
		return fmt.Errorf("minLength must be greater than, or equal to, 0. minLenght: %d", min)
	}

	c := len(val)
	if c < min {
		return MinLengthError{min, c}
	}

	return nil
}

func Pattern(val, pattern string) error {
	ok, err := regexp.MatchString(pattern, val)
	if err != nil {
		return err
	}

	if ok == false {
		return PatternError{pattern, val}
	}

	return nil
}

func Enum(val string, enum []string) error {
	if len(enum) < 1 {
		return fmt.Errorf("enum must have at least one element. enum element count:%d", len(enum))
	}

	// element in the enum must be unique.
	for i, e := range enum {
		for j := i + 1; j < len(enum); j++ {
			if e == enum[j] {
				return fmt.Errorf("elements in the enum must be unique. `%s`", enum)
			}
		}
	}

	for _, e := range enum {
		if val == e {
			return nil
		}
	}

	return EnumError{enum, val}
}
