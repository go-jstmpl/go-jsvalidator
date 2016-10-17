package strings

import (
	"fmt"
	"regexp"
)

var (
	rDateTime = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(?:\\.\\d{1,9})?(?:[+-]\\d{2}:\\d{2}|Z)$")
	rEmail    = regexp.MustCompile("^.+@.+\\..+$")
	rURI      = regexp.MustCompile("^[0-9a-zA-Z]+:\\/\\/.+$")
)

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
	return fmt.Sprintf("input value '%s' does not match the pattern for '%s'\n",
		f.Input, f.Definition.Format)
}

func NewFormatValidator(definition FormatValidatorDefinition) (FormatValidator, error) {
	switch definition.Format {
	case "date-time", "email", "hostname", "uri", "password-0Aa":
		return FormatValidator{definition}, nil
	}
	return FormatValidator{}, InvalidFormatError{"the format is not found"}
}

func (f FormatValidator) Validate(input string) error {
	switch f.definition.Format {
	case "date-time":
		ok := rDateTime.MatchString(input)
		if ok {
			return nil
		}
		break
	case "email":
		ok := rEmail.MatchString(input)
		if ok {
			return nil
		}
		break
	case "hostname":
		ok := isHostName(input)
		if ok {
			return nil
		}
		break
	case "uri":
		ok := rURI.MatchString(input)
		if ok {
			return nil
		}
		break
	case "password-0Aa":
		ok := isPassword0Aa(input)
		if ok {
			return nil
		}
		break
	}
	return &FormatValidationError{
		f.definition,
		input,
	}
}

// isHostName is cited from https://golang.org/src/net/dnsclient.go
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

func isPassword0Aa(pw string) bool {
	large := false
	small := false
	number := false

	for _, r := range pw {
		switch {
		case '0' <= r && r <= '9':
			number = true
			continue
		case 'A' <= r && r <= 'Z':
			large = true
			continue
		case 'a' <= r && r <= 'z':
			small = true
			continue
		case '!' <= r && r <= '~': // when other ASCII characters
			continue
		default:
			return false
		}
	}
	return large && small && number
}
