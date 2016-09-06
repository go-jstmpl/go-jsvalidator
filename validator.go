package validator

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"unicode/utf8"
)

var (
	rDateTime = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$")
	rEmail    = regexp.MustCompile("^.+@.+\\..+$")
	rURI      = regexp.MustCompile("^[0-9a-zA-Z]+:\\/\\/.+$")
)

// EmptyError for Constructor methods
type EmptyError struct {
	message string
}

func (e EmptyError) Error() string {
	return e.message
}

// DuplicationError for Constructor methods
type DuplicationError struct {
	message string
}

func (e DuplicationError) Error() string {
	return e.message
}

// NoLengthError for Constructor methods
type NoLengthError struct {
	message string
}

func (e NoLengthError) Error() string {
	return e.message
}

// InvalidPatternError for Constructor methods
type InvalidPatternError struct {
	message string
}

func (e InvalidPatternError) Error() string {
	return e.message
}

// InvalidFormatError for Constructor methods
type InvalidFormatError struct {
	message string
}

func (e InvalidFormatError) Error() string {
	return e.message
}

// InvalidFieldTypeError for Required Validate method
type InvalidFieldTypeError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func (e InvalidFieldTypeError) Error() string {
	return fmt.Sprintf("input struct have invalid field against required '%v'", e.Definition.Required)
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

func (i IntMaximumValidationError) Error() string {
	return fmt.Sprintf("IntMaximumValidator: expected less than %d but actual %d with option exlusive %t",
		i.Definition.Maximum, i.Input, i.Definition.Exclusive)
}

func NewIntMaximumValidator(definition IntMaximumValidatorDefinition) (IntMaximumValidator, error) {
	return IntMaximumValidator{definition}, nil
}

func (m IntMaximumValidator) Validate(input int) error {
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

func (i IntMinimumValidationError) Error() string {
	return fmt.Sprintf("IntMinimumValidator: expected greater than %d but actual %d with option exlusive %t",
		i.Definition.Minimum, i.Input, i.Definition.Exclusive)
}

func NewIntMinimumValidator(definition IntMinimumValidatorDefinition) (IntMinimumValidator, error) {
	return IntMinimumValidator{definition}, nil
}

func (m IntMinimumValidator) Validate(input int) error {
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

func (f FloatMaximumValidationError) Error() string {
	return fmt.Sprintf("FloatMaximumValidator: expected less than %g but actual %g with option exlusive %t",
		f.Definition.Maximum, f.Input, f.Definition.Exclusive)
}

func NewFloatMaximumValidator(definition FloatMaximumValidatorDefinition) (FloatMaximumValidator, error) {
	return FloatMaximumValidator{definition}, nil
}

func (m FloatMaximumValidator) Validate(input float64) error {
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

func (f FloatMinimumValidationError) Error() string {
	return fmt.Sprintf("FloatMinimumValidator: expected greater than %g but actual %g with option exlusive %t",
		f.Definition.Minimum, f.Input, f.Definition.Exclusive)
}

func NewFloatMinimumValidator(definition FloatMinimumValidatorDefinition) (FloatMinimumValidator, error) {
	return FloatMinimumValidator{definition}, nil
}

func (m FloatMinimumValidator) Validate(input float64) error {
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

func (m MaxLengthValidationError) Error() string {
	return fmt.Sprintf("MaxLengthValidator: expected less than, or equal to, %d charactors but actual %d charactors",
		m.Definition.MaxLength, utf8.RuneCountInString(m.Input))
}

func NewMaxLengthValidator(definition MaxLengthValidatorDefinition) (MaxLengthValidator, error) {
	if definition.MaxLength < 0 {
		return MaxLengthValidator{}, NoLengthError{"the max length should be greater than, or equal to, 0"}
	}
	return MaxLengthValidator{definition}, nil
}

func (m MaxLengthValidator) Validate(input string) error {
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

func (m MinLengthValidationError) Error() string {
	return fmt.Sprintf("MinLengthValidator: expected greater than, or equal to, %d charactors but actual %d charactors",
		m.Definition.MinLength, utf8.RuneCountInString(m.Input))
}

func NewMinLengthValidator(definition MinLengthValidatorDefinition) (MinLengthValidator, error) {
	if definition.MinLength < 0 {
		return MinLengthValidator{}, fmt.Errorf("the minimum length should be greater than, or equal to, 0")
	}

	return MinLengthValidator{definition}, nil
}

func (m MinLengthValidator) Validate(input string) error {
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

func (p PatternValidationError) Error() string {
	return fmt.Sprintf("PatternValidator: input value '%s' does not match the regex pattern '%s'", p.Input, p.Definition.Pattern)
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

func (p PatternValidator) Validate(input string) error {
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

func (i IntEnumValidationError) Error() string {
	return fmt.Sprintf("IntEnumValidator: input value '%d' does not listed in enumerate '%v'", i.Input, i.Definition.Enumerate)
}

func NewIntEnumValidator(definition IntEnumValidatorDefinition) (IntEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return IntEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		key := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == key {
				return IntEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}

	return IntEnumValidator{definition}, nil

}

func (i IntEnumValidator) Validate(input int) error {
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

func (s StringEnumValidationError) Error() string {
	return fmt.Sprintf("StringEnumValidator: input value '%s' does not listed in enumerate '%v'", s.Input, s.Definition.Enumerate)
}

func NewStringEnumValidator(definition StringEnumValidatorDefinition) (StringEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return StringEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		key := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == key {
				return StringEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}

	return StringEnumValidator{definition}, nil
}

func (s StringEnumValidator) Validate(input string) error {
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

func (s FloatEnumValidationError) Error() string {
	return fmt.Sprintf("FloatEnumValidator: input value '%g' does not listed in enumerate '%v'", s.Input, s.Definition.Enumerate)
}

func NewFloatEnumValidator(definition FloatEnumValidatorDefinition) (FloatEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return FloatEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		key := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == key {
				return FloatEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}
	return FloatEnumValidator{definition}, nil

}

func (f FloatEnumValidator) Validate(input float64) error {
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

func (r RequiredValidationError) Error() string {
	return fmt.Sprintf("RequiredValidator: input sturct does not satisfy required values '%v'", r.Definition.Required)
}

func NewRequiredValidator(definition RequiredValidatorDefinition) (RequiredValidator, error) {
	required := definition.Required
	len := len(required)
	if len == 0 {
		return RequiredValidator{}, EmptyError{"the required value should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		key := required[i]
		for j := i + 1; j < len; j++ {
			if required[j] == key {
				return RequiredValidator{}, DuplicationError{"the required value should not be duplicated"}
			}
		}
	}

	return RequiredValidator{definition}, nil
}

func (r RequiredValidator) Validate(input interface{}) error {
	elem := reflect.ValueOf(input).Elem()

	for _, key := range r.definition.Required {
		e := elem.FieldByName(key)
		if !e.IsValid() {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		n, ok := e.Interface().(driver.Valuer)
		if !ok {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		v, err := n.Value()
		if err != nil {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		if v == nil {
			return &RequiredValidationError{
				Definition: r.definition,
				Input:      input,
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

func (i IntMaxItemsValidationError) Error() string {
	return fmt.Sprintf("IntMaxItemsValidator: the number of input values is greater than MaxItems:'%d'", i.Definition.MaxItems)
}

func NewIntMaxItemsValidator(definition IntMaxItemsValidatorDefinition) (IntMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return IntMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return IntMaxItemsValidator{definition}, nil
}

func (i IntMaxItemsValidator) Validate(input []int) error {
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

func (f FloatMaxItemsValidationError) Error() string {
	return fmt.Sprintf("FloatMaxItemsValidator: the number of input values is greater than MaxItems:'%d'", f.Definition.MaxItems)
}

func NewFloatMaxItemsValidator(definition FloatMaxItemsValidatorDefinition) (FloatMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return FloatMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return FloatMaxItemsValidator{definition}, nil
}

func (i FloatMaxItemsValidator) Validate(input []float64) error {
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

func (s StringMaxItemsValidationError) Error() string {
	return fmt.Sprintf("FloatMaxItemsValidator: the number of input values is greater than MaxItems:'%d'", s.Definition.MaxItems)
}

func NewStringMaxItemsValidator(definition StringMaxItemsValidatorDefinition) (StringMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return StringMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}

	return StringMaxItemsValidator{definition}, nil
}

func (i StringMaxItemsValidator) Validate(input []string) error {
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

func (f FormatValidationError) Error() string {
	return fmt.Sprintf("FormatValidator: value '%s' does not match the regex pattern for '%s'",
		f.Input, f.Definition.Format)
}

func NewFormatValidator(definition FormatValidatorDefinition) (FormatValidator, error) {
	switch definition.Format {
	case "date-time", "email", "hostname", "uri":
		return FormatValidator{definition}, nil
	}
	return FormatValidator{}, InvalidFormatError{"the format is not found"}
}

func (f FormatValidator) Validate(input string) error {
	switch f.definition.Format {
	case "date-time":
		ok := rDateTime.MatchString(input)
		if ok == true {
			return nil
		}
		break
	case "email":
		ok := rEmail.MatchString(input)
		if ok == true {
			return nil
		}
		break
	case "hostname":
		ok := isHostName(input)
		if ok == true {
			return nil
		}
		break
	case "uri":
		ok := rURI.MatchString(input)
		if ok == true {
			return nil
		}
		break
	}
	return &FormatValidationError{
		f.definition,
		input,
	}
}

// isHostName stolen from https://golang.org/src/net/dnsclient.go
func isHostName(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) > 255 {
		return false
	}

	last := byte(',')
	ok := false
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			ok = true
			partlen++
		case '0' <= c && c <= 9:
			partlen++
		case c == '-':
			if last == '.' {
				return false
			}
			partlen++
		case c == '.':
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return ok
}
