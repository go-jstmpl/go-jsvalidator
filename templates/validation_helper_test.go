package validator

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	err := Maximum(100, 200, true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMinimum(t *testing.T) {
	err := Minimum(100, 10, true)
	if err != nil {
		t.Error(err)
	}

	err = Minimum(200, 200, true)
	if err == nil {
		t.Error(err)
	}

	err = Minimum(100, 10, false)
	if err != nil {
		t.Error(err)
	}

	err = Minimum(200, 200, false)
	if err != nil {
		t.Error(err)
	}

	err = Minimum(0, 0, false)
	if err != nil {
		t.Error(err)
	}

	err = Minimum(0, 0, true)
	if err == nil {
		t.Error(err)
	}
}

func TestMaxLength(t *testing.T) {
	err := MaxLength("hogehoge", -1)
	if err == nil {
		t.Error(err)
	}

	err = MaxLength("hogehoge", 8)
	if err != nil {
		t.Error(err)
	}

	err = MaxLength("hogehoge", 6)
	if err == nil {
		t.Error(err)
	}

	err = MaxLength("hogehoge", 1111)
	if err != nil {
		t.Error(err)
	}
}

func TestMinLength(t *testing.T) {
	err := MinLength("hogehoge", 6)
	if err != nil {
		t.Error(err)
	}

	err = MinLength("hogehoge", 8)
	if err != nil {
		t.Error(err)
	}

	err = MinLength("hogehoge", 10)
	if err == nil {
		t.Error(err)
	}

	err = MinLength("hogehoge", 0)
	if err != nil {
		t.Error(err)
	}
}

func TestPattern(t *testing.T) {
	err := Pattern("2016-05-09T19:45:32Z", "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$")
	if err != nil {
		t.Error(err)
	}

	err = Pattern("", "")
	if err != nil {
		t.Error(err)
	}
}

func TestEnum(t *testing.T) {
	en := []string{"amber", "red", "blue"}
	err := Enum("red", en)
	if err != nil {
		t.Error(err)
	}
}
