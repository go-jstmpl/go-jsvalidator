package validator

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type MaximumValidator struct {
	Maximum   int  `json:"maximum"`
	Exclusive bool `json:"exclusive"`
}

func (m MaximumValidator) Error(val int) string {
	return fmt.Sprintf("expected less than %d but actual %d with option exlusive %b",
		m.Maximum, val, m.Exclusive)
}

func NewMaximumValidator(maximum int, exclusive bool) (*MaximumValidator, error) {
	v := &MaximumValidator{
		Maximum:   maximum,
		Exclusive: exclusive,
	}

	return v, nil
}

//If "exclusiveMaximum" is not present, or has boolean value false,
//then the instance is valid if it is less than,
//or equal to, the value of "minimum".
//If "exlusiveMaximum" is present and has boolean value true,
//the instance is valid if it is strictly less than the value of "maximum".
func (m MaximumValidator) Validate(val int) bool {
	if !m.Exclusive {
		if val <= m.Maximum {
			return true
		}
		return false
	}

	if val < m.Maximum {
		return true
	}
	return false
}

type MinimumValidator struct {
	Minimum   int  `json:"minimum"`
	Exclusive bool `json:"exclusive"`
}

func (m MinimumValidator) Error(val int) string {
	return fmt.Sprintf("expected greater than %d but actual %d with option exlusive %b",
		m.Minimum, val, m.Exclusive)
}

func NewMinimumValidator(minimum int, exclusive bool) (*MinimumValidator, error) {
	v := &MinimumValidator{
		Minimum:   minimum,
		Exclusive: exclusive,
	}

	return v, nil
}

//If "exclusiveMinimum" is not present, or has boolean value false,
//then the instance is valid if it is greater than,
//or equal to, the value of "minimum"
//if "exclusiveMinimum" is present and has boolean value true,
//the instance is valid if it is strictly greater than the value of "minimum".
func (m MinimumValidator) Validate(val int) bool {
	if !m.Exclusive {
		if val >= m.Minimum {
			return true
		}
		return false
	}

	if val > m.Minimum {
		return true
	}
	return false
}

type MaxLengthValidator struct {
	MaxLength int `json:"maxlength"`
}

func (m MaxLengthValidator) Error(val string) string {
	return fmt.Sprintf("expected less than, or equal to, %d charactors but actual %d",
		m.MaxLength, utf8.RuneCountInString(val))
}

func NewMaxLengthValidator(maxLength int) (*MaxLengthValidator, error) {
	if maxLength <= 0 {
		return nil, fmt.Errorf("maxLength must be greater than 0. maxLength: %d", maxLength)
	}
	v := &MaxLengthValidator{
		MaxLength: maxLength,
	}

	return v, nil
}

func (m MaxLengthValidator) Validate(val string) bool {
	if utf8.RuneCountInString(val) <= m.MaxLength {
		return true
	}

	return false
}

type MinLengthValidator struct {
	MinLength int `json:"minLength"`
}

func (m MinLengthValidator) Error(val string) string {
	return fmt.Sprintf("expected greater than, or equal to, %d charactors but actual %d charactors",
		m.MinLength, utf8.RuneCountInString(val))
}

func NewMinLengthValidator(minLength int) (*MinLengthValidator, error) {
	if minLength <= 0 {
		return nil, fmt.Errorf("minLength must be less than 0. minLength: %d", minLength)
	}
	v := &MinLengthValidator{
		MinLength: minLength,
	}

	return v, nil
}

func (m MinLengthValidator) Validate(val string) bool {
	if utf8.RuneCountInString(val) >= m.MinLength {
		return true
	}

	return false
}

type PatternValidator struct {
	Pattern string
}

func (m PatternValidator) Error(val string) string {
	return fmt.Sprintf("%s expected matching pattern: %s", val, m.Pattern)
}

func NewPatternValidator(pattern string) (*PatternValidator, error) {
	_, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	v := &PatternValidator{
		Pattern: pattern,
	}

	return v, nil
}

func (m PatternValidator) Validate(val string) bool {
	ok, err := regexp.MatchString(m.Pattern, val)
	if err != nil {
		return false
	}

	if ok == false {
		return false
	}

	return true
}

type IntEnumValidator struct {
	Enumerate []int
}

func (m IntEnumValidator) Error(val string) string {
	return fmt.Sprintf("%s expected include emunerate list: %v", val, m.Enumerate)
}

func NewIntEnumValidator(enumerate []int) (*IntEnumValidator, error) {
	if len(enumerate) == 0 {
		return nil, errors.New("enumerate element should not be empty")
	}

	// unique test
	for idx, e := range enumerate {
		for _, es := range enumerate[idx+1:] {
			if e == es {
				return nil, errors.New("enumerate element should be unique")
			}
		}
	}

	v := &IntEnumValidator{
		Enumerate: enumerate,
	}

	return v, nil
}

func (m IntEnumValidator) Validate(val int) bool {
	for _, e := range m.Enumerate {
		if val == e {
			return true
		}
	}

	return false
}

type StringEnumValidator struct {
	Enumerate []string
}

func (m StringEnumValidator) Error(val string) string {
	return fmt.Sprintf("%s expected include enumerate list: %v", val, m.Enumerate)
}

func NewStringEnumValidator(enumerate []string) (*StringEnumValidator, error) {
	if len(enumerate) == 0 {
		return nil, errors.New("enumerate element should not be empty")
	}

	// unique test
	for idx, e := range enumerate {
		for _, es := range enumerate[idx+1:] {
			if e == es {
				return nil, errors.New("enumerate element should be unique")
			}
		}
	}

	v := &StringEnumValidator{
		Enumerate: enumerate,
	}

	return v, nil
}

func (m StringEnumValidator) Validate(val string) bool {
	for _, e := range m.Enumerate {
		if val == e {
			return true
		}
	}

	return false
}
