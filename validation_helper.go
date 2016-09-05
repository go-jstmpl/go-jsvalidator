package validator

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"unicode/utf8"
)

type EmptyError struct {
	message string
}

func (e EmptyError) Error() string {
	return e.message
}

type DuplicationError struct {
	message string
}

func (e DuplicationError) Error() string {
	return e.message
}

type NoLengthError struct {
	message string
}

func (e NoLengthError) Error() string {
	return e.message
}

type InvalidPatternError struct {
	message string
}

func (e InvalidPatternError) Error() string {
	return e.message
}

type InvalidFormatError struct {
	message string
}

func (e InvalidFormatError) Error() string {
	return e.message
}

type IntMaximumValidator struct {
	definition IntMaximumValidatorDefinition
}

type IntMaximumValidatorDefinition struct {
	Maximum   int  `json:"maximum"`
	Exclusive bool `json:"exclusive"`
}

type IntMaximumValidationError struct {
	Definition IntMaximumValidatorDefinition `json:"definition"`
	Input      int                           `json:"input"`
}

func NewIntMaximumValidator(definition IntMaximumValidatorDefinition) (IntMaximumValidator, error) {
	return IntMaximumValidator{definition}, nil
}

func (m IntMaximumValidator) Validate(input int) *IntMaximumValidationError {
	if !m.definition.Exclusive {
		if input <= m.definition.Maximum {
			return nil
		}
		return &IntMaximumValidationError{
			m.definition,
			input,
		}
	}

	if input < m.definition.Maximum {
		return nil
	}
	return &IntMaximumValidationError{
		m.definition,
		input,
	}
}

type IntMinimumValidator struct {
	definition IntMinimumValidatorDefinition
}

type IntMinimumValidatorDefinition struct {
	Minimum   int  `json:"minimum"`
	Exclusive bool `json:"exclusive"`
}

type IntMinimumValidationError struct {
	Definition IntMinimumValidatorDefinition `json:"definition"`
	Input      int                           `json:"input"`
}

func NewIntMinimumValidator(definition IntMinimumValidatorDefinition) (IntMinimumValidator, error) {
	return IntMinimumValidator{definition}, nil
}

func (m IntMinimumValidator) Validate(input int) *IntMinimumValidationError {
	if !m.definition.Exclusive {
		if input >= m.definition.Minimum {
			return nil
		}
		return &IntMinimumValidationError{
			m.definition,
			input,
		}
	}
	if input > m.definition.Minimum {
		return nil
	}
	return &IntMinimumValidationError{
		m.definition,
		input,
	}
}

type FloatMaximumValidator struct {
	definition FloatMaximumValidatorDefinition
}

type FloatMaximumValidatorDefinition struct {
	Maximum   float64 `json:"maximum"`
	Exclusive bool    `json:"exclusive"`
}

type FloatMaximumValidationError struct {
	Definition FloatMaximumValidatorDefinition `json:"definition"`
	Input      float64                         `json:"input"`
}

func NewFloatMaximumValidator(definition FloatMaximumValidatorDefinition) (FloatMaximumValidator, error) {
	return FloatMaximumValidator{definition}, nil
}

func (m FloatMaximumValidator) Validate(input float64) *FloatMaximumValidationError {
	if !m.definition.Exclusive {
		if input <= m.definition.Maximum {
			return nil
		}
		return &FloatMaximumValidationError{
			m.definition,
			input,
		}
	}
	if input < m.definition.Maximum {
		return nil
	}
	return &FloatMaximumValidationError{
		m.definition,
		input,
	}
}

type FloatMinimumValidator struct {
	definition FloatMinimumValidatorDefinition
}

type FloatMinimumValidatorDefinition struct {
	Minimum   float64 `json:"minimum"`
	Exclusive bool    `json:"exclusive"`
}

type FloatMinimumValidationError struct {
	Definition FloatMinimumValidatorDefinition `json:"definition"`
	Input      float64                         `json:"input"`
}

func NewFloatMinimumValidator(definition FloatMinimumValidatorDefinition) (FloatMinimumValidator, error) {
	return FloatMinimumValidator{definition}, nil
}

func (m FloatMinimumValidator) Validate(input float64) *FloatMinimumValidationError {
	if !m.definition.Exclusive {
		if input >= m.definition.Minimum {
			return nil
		}
		return &FloatMinimumValidationError{
			m.definition,
			input,
		}
	}

	if input > m.definition.Minimum {
		return nil
	}
	return &FloatMinimumValidationError{
		m.definition,
		input,
	}
}

type MaxLengthValidator struct {
	definition MaxLengthValidatorDefinition
}

type MaxLengthValidatorDefinition struct {
	MaxLength int `json:"max_length"`
}

type MaxLengthValidationError struct {
	Definition MaxLengthValidatorDefinition `json:"definition"`
	Input      string                       `json:"input"`
}

func NewMaxLengthValidator(definition MaxLengthValidatorDefinition) (MaxLengthValidator, error) {
	if definition.MaxLength < 0 {
		return MaxLengthValidator{}, NoLengthError{"the max length should be greater than, or equal to, 0"}
	}
	return MaxLengthValidator{definition}, nil
}

func (m MaxLengthValidator) Validate(input string) *MaxLengthValidationError {
	if utf8.RuneCountInString(input) <= m.definition.MaxLength {
		return nil
	}
	return &MaxLengthValidationError{
		m.definition,
		input,
	}
}

type MinLengthValidator struct {
	definition MinLengthValidatorDefinition
}

type MinLengthValidatorDefinition struct {
	MinLength int `json:"min_length"`
}

type MinLengthValidationError struct {
	Definition MinLengthValidatorDefinition `json:"definition"`
	Input      string                       `json:"input"`
}

func NewMinLengthValidator(definition MinLengthValidatorDefinition) (MinLengthValidator, error) {
	if definition.MinLength < 0 {
		return MinLengthValidator{}, fmt.Errorf("the minimum length should be greater than, or equal to, 0")
	}

	return MinLengthValidator{definition}, nil
}

func (m MinLengthValidator) Validate(input string) *MinLengthValidationError {
	if utf8.RuneCountInString(input) >= m.definition.MinLength {
		return nil
	}
	return &MinLengthValidationError{
		m.definition,
		input,
	}
}

type PatternValidator struct {
	definition PatternValidatorDefinition
}

type PatternValidatorDefinition struct {
	Pattern string `json:"pattern"`
}

type PatternValidationError struct {
	Definition PatternValidatorDefinition `json:"definition"`
	Input      string                     `json:"input"`
}

func NewPatternValidator(definition PatternValidatorDefinition) (PatternValidator, error) {
	if definition.Pattern == "" {
		return PatternValidator{}, EmptyError{"the pattern should not be empty"}
	}
	_, err := regexp.Compile(definition.Pattern)
	if err != nil {
		return PatternValidator{}, InvalidPatternError{fmt.Sprintf("invalid pattern: %s", definition.Pattern)}
	}
	return PatternValidator{definition}, nil
}

func (p PatternValidator) Validate(input string) *PatternValidationError {
	ok, err := regexp.MatchString(p.definition.Pattern, input)
	if ok == true && err == nil {
		return nil
	}
	return &PatternValidationError{
		p.definition,
		input,
	}
}

type IntEnumValidator struct {
	definition IntEnumValidatorDefinition
}

type IntEnumValidatorDefinition struct {
	Enumerate []int `json:"enum"`
}

type IntEnumValidationError struct {
	Definition IntEnumValidatorDefinition `json:"definition"`
	Input      int                        `json:"input"`
}

func NewIntEnumValidator(definition IntEnumValidatorDefinition) (IntEnumValidator, error) {
	enumerate := definition.Enumerate
	if len(enumerate) == 0 {
		return IntEnumValidator{}, EmptyError{"the enumurate should have at least one element"}
	}

	for i, e := range enumerate {
		for _, es := range enumerate[i+1:] {
			if e == es {
				return IntEnumValidator{}, DuplicationError{"the elements of enumurate should not be duplicated"}
			}
		}
	}
	return IntEnumValidator{definition}, nil

}

func (i IntEnumValidator) Validate(input int) *IntEnumValidationError {
	for _, e := range i.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &IntEnumValidationError{
		i.definition,
		input,
	}
}

type StringEnumValidator struct {
	definition StringEnumValidatorDefinition
}

type StringEnumValidatorDefinition struct {
	Enumerate []string `json:"enum"`
}

type StringEnumValidationError struct {
	Definition StringEnumValidatorDefinition `json:"definition"`
	Input      string                        `json:"input"`
}

func NewStringEnumValidator(definition StringEnumValidatorDefinition) (StringEnumValidator, error) {
	enumerate := definition.Enumerate
	if len(enumerate) == 0 {
		return StringEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i, e := range enumerate {
		for _, es := range enumerate[i+1:] {
			if e == es {
				return StringEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}
	return StringEnumValidator{definition}, nil
}

func (s StringEnumValidator) Validate(input string) *StringEnumValidationError {
	for _, e := range s.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &StringEnumValidationError{
		s.definition,
		input,
	}
}

type FloatEnumValidator struct {
	definition FloatEnumValidatorDefinition
}

type FloatEnumValidatorDefinition struct {
	Enumerate []float64 `json:"enum"`
}

type FloatEnumValidationError struct {
	Definition FloatEnumValidatorDefinition `json:"definition"`
	Input      float64                      `json:"input"`
}

func NewFloatEnumValidator(definition FloatEnumValidatorDefinition) (FloatEnumValidator, error) {
	enumerate := definition.Enumerate
	if len(enumerate) == 0 {
		return FloatEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i, e := range enumerate {
		for _, es := range enumerate[i+1:] {
			if e == es {
				return FloatEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}
	return FloatEnumValidator{definition}, nil

}

func (f FloatEnumValidator) Validate(input float64) *FloatEnumValidationError {
	for _, e := range f.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &FloatEnumValidationError{
		f.definition,
		input,
	}
}

type RequiredValidator struct {
	definition RequiredValidatorDefinition
}

type RequiredValidatorDefinition struct {
	Required []string `json:"pattern"`
}

type RequiredValidationError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func NewRequiredValidator(definition RequiredValidatorDefinition) (RequiredValidator, error) {
	required := definition.Required
	if len(required) == 0 {
		return RequiredValidator{}, EmptyError{"the required value should have at least one element"}
	}
	for idx, e := range required {
		for _, es := range required[idx+1:] {
			if e == es {
				return RequiredValidator{}, DuplicationError{"the required value should not be duplicated"}
			}
		}
	}
	return RequiredValidator{definition}, nil
}

func (r RequiredValidator) Validate(input interface{}) *RequiredValidationError {
	// convert reflect data to interfaced struct
	elem := reflect.ValueOf(input).Elem()
	size := elem.NumField()

	for s := 0; s < size; s++ {
		for _, require := range r.definition.Required {
			// getting required field
			if require == elem.Type().Field(s).Name {
				// change into a type that have null check function
				n, ok := elem.Field(s).Interface().(interface {
					Value() (driver.Value, error)
				})
				if ok != true {
					return &RequiredValidationError{
						Definition: r.definition,
						Input:      input,
					}
				}
				v, _ := n.Value()
				if v == nil {
					return &RequiredValidationError{
						Definition: r.definition,
						Input:      input,
					}
				}
			}
		}
	}
	return nil
}

type IntMaxItemsValidator struct {
	definition IntMaxItemsValidatorDefinition
}

type IntMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type IntMaxItemsValidationError struct {
	Definition IntMaxItemsValidatorDefinition `json:"definition"`
	Input      []int                          `json:"input"`
}

func NewIntMaxItemsValidator(definition IntMaxItemsValidatorDefinition) (IntMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return IntMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return IntMaxItemsValidator{definition}, nil
}

func (i IntMaxItemsValidator) Validate(input []int) *IntMaxItemsValidationError {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &IntMaxItemsValidationError{
		i.definition,
		input,
	}
}

type FloatMaxItemsValidator struct {
	definition FloatMaxItemsValidatorDefinition
}

type FloatMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type FloatMaxItemsValidationError struct {
	Definition FloatMaxItemsValidatorDefinition `json:"definition"`
	Input      []float64                        `json:"input"`
}

func NewFloatMaxItemsValidator(definition FloatMaxItemsValidatorDefinition) (FloatMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return FloatMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return FloatMaxItemsValidator{definition}, nil
}

func (i FloatMaxItemsValidator) Validate(input []float64) *FloatMaxItemsValidationError {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &FloatMaxItemsValidationError{
		i.definition,
		input,
	}
}

type StringMaxItemsValidator struct {
	definition StringMaxItemsValidatorDefinition
}

type StringMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type StringMaxItemsValidationError struct {
	Definition StringMaxItemsValidatorDefinition `json:"definition"`
	Input      []string                          `json:"input"`
}

func NewStringMaxItemsValidator(definition StringMaxItemsValidatorDefinition) (StringMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return StringMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}

	return StringMaxItemsValidator{definition}, nil
}

func (i StringMaxItemsValidator) Validate(input []string) *StringMaxItemsValidationError {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &StringMaxItemsValidationError{
		i.definition,
		input,
	}
}

type FormatValidator struct {
	definition FormatValidatorDefinition
}

type FormatValidatorDefinition struct {
	Format string `json:"format"`
}

type FormatValidationError struct {
	Definition FormatValidatorDefinition `json:"definition"`
	Input      string                    `json:"input"`
}

func NewFormatValidator(definition FormatValidatorDefinition) (FormatValidator, error) {
	switch definition.Format {
	case "date-time":
		return FormatValidator{definition}, nil
	case "email":
		return FormatValidator{definition}, nil
	case "uri":
		return FormatValidator{definition}, nil
	}
	return FormatValidator{}, InvalidFormatError{"the format is not found"}
}

func (f FormatValidator) Validate(input string) *FormatValidationError {
	const (
		rDateTime = "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$"
		rEmail    = "^.+@.+\\..+$"
		rURI      = "^https?:\\/\\/.+$"
	)

	switch f.definition.Format {
	case "date-time":
		ok, err := regexp.MatchString(rDateTime, input)
		if ok == true && err == nil {
			return nil
		}
		break
	case "email":
		ok, err := regexp.MatchString(rEmail, input)
		if ok == true && err == nil {
			return nil
		}
		break
	case "uri":
		ok, err := regexp.MatchString(rURI, input)
		if ok == true && err == nil {
			return nil
		}
		break
	}
	return &FormatValidationError{
		f.definition,
		input,
	}
}
