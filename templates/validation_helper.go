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

func Maximum(val int, max int, exc bool) error {
	if exc {
		if val >= max {
			return MaximumError{max, val, exc}
		}
	} else {
		if val > max {
			return MaximumError{max, val, exc}
		}
	}

	return nil
}

func Minimum(val int, min int, exc bool) error {
	if exc {
		if val <= min {
			return MinimumError{min, val, exc}
		}
	} else {
		if val < min {
			return MinimumError{min, val, exc}
		}
	}

	return nil
}

func MaxLength(val string, max int) error {
	if max == 0 {
		return nil
	}

	c := len(val)
	if c > max {
		return MaxLengthError{max, c}
	}

	return nil
}

func MinLength(val string, min int) error {
	if min == 0 {
		return nil
	}
	c := len(val)
	if c < min {
		return MinLengthError{min, c}
	}

	return nil
}

func Pattern(val string, regx string) error {
	if regx == "" {
		return nil
	}

	ok, err := regexp.MatchString(regx, val)
	if err != nil || ok != true {
		return PatternError{regx, val}
	}

	return nil
}

func Enum(str string, accepts ...string) error {
	for _, v := range accepts {
		if str == v {
			return nil
		}
	}

	return EnumError{accepts, str}
}
