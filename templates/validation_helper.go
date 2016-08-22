package validator

import (
	"fmt"
	"regexp"
)

/*
if "exclusiveMaximum" is not present, or has boolean value false, then the instance is valid if it is lower than, or equal to, the value of "maximum";
if "exclusiveMaximum" has boolean value true, the instance is valid if it is strictly lower than the value of "maximum".
*/
type MaximumError struct {
	Expected  int  `json:"expected"`
	Actual    int  `json:"actual"`
	Exclusive bool `json:"exclusive"`
}

func (m MaximumError) Error() string {
	return fmt.Sprintf("expected less than %d but actual %d with option exlusive %b", m.Expected, m.Actual, m.Exclusive)
}

/*
if "exclusiveMinimum" is not present, or has boolean value false, then the instance is valid if it is greater than, or equal to, the value of "minimum";

if "exclusiveMinimum" is present and has boolean value true, the instance is valid if it is strictly greater than the value of "minimum".
*/
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

func Pattern(val, regx string) error {
	ok, err := regexp.MatchString(regx, val)
	if err != nil {
		return err
	}

	if ok {
		return PatternError{regx, val}
	}

	return nil
}

func Enum(str string, accepts []string) error {
	if len(accepts) < 1 {
		return fmt.Errorf("enum must have at least one element. enum element count:%d", len(accepts))
	}

	for i, e := range accepts {
		if i == len(accepts)-1 {
			break
		}

		for j := i; j < len(accepts); i++ {
			if e == accepts[j+1] {
				return fmt.Errorf("elements in the enum must be unique. `%s`", accepts)
			}
		}
	}

	for _, v := range accepts {
		if str == v {
			return nil
		}
	}

	return EnumError{accepts, str}
}
