package validator

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	err := Maximum(100, 200, true)
	if err != nil {
		t.Log(err, "100, 200, true")
		t.Fail()
	}
	err = Maximum(200, 200, true)
	if err == nil {
		t.Log(err, "200, 200, true")
		t.Fail()
	}
	err = Maximum(100, 200, false)
	if err != nil {
		t.Log(err, "100, 200, false")
		t.Fail()
	}
	err = Maximum(200, 200, false)
	if err != nil {
		t.Log(err, "200, 200, false")
		t.Fail()
	}
	err = Maximum(0, 0, false)
	if err != nil {
		t.Log(err, "0, 0, false")
		t.Fail()
	}
	err = Maximum(0, 0, true)
	if err == nil {
		t.Log(err, "0, 0, true")
		t.Fail()
	}
}

func TestMinimum(t *testing.T) {
	err := Minimum(100, 10, true)
	if err != nil {
		t.Log(err, "100, 10, true")
		t.Fail()
	}
	err = Minimum(200, 200, true)
	if err == nil {
		t.Log(err, "200, 200, true")
		t.Fail()
	}
	err = Minimum(100, 10, false)
	if err != nil {
		t.Log(err, "100, 10, false")
		t.Fail()
	}
	err = Minimum(200, 200, false)
	if err != nil {
		t.Log(err, "200, 200, false")
		t.Fail()
	}
	err = Minimum(0, 0, false)
	if err != nil {
		t.Log(err, "0, 0, false")
		t.Fail()
	}
	err = Minimum(0, 0, true)
	if err == nil {
		t.Log(err, "0, 0, true")
		t.Fail()
	}
}

func TestMaxLength(t *testing.T) {
	err := MaxLength("hogehoge", 10)
	if err != nil {
		t.Fail()
	}

	err = MaxLength("hogehoge", 8)
	if err != nil {
		t.Fail()
	}

	err = MaxLength("hogehoge", 6)
	if err == nil {
		t.Fail()
	}

	err = MaxLength("hogehoge", 0)
	if err != nil {
		t.Fail()
	}
}

func TestMinLength(t *testing.T) {
	err := MinLength("hogehoge", 6)
	if err != nil {
		t.Fail()
	}

	err = MinLength("hogehoge", 8)
	if err != nil {
		t.Fail()
	}

	err = MinLength("hogehoge", 10)
	if err == nil {
		t.Fail()
	}

	err = MinLength("hogehoge", 0)
	if err != nil {
		t.Fail()
	}
}

func TestPattern(t *testing.T) {
	err := Pattern("seafood", `foo.*`)
	if err != nil {
		t.Fail()
	}

	err = Pattern("", "")
	if err != nil {
		t.Fail()
	}
}

func TestEnum(t *testing.T) {
	en := []string{"red", "amber", "green"}
	err := Enum("red", en...)
	if err != nil {
		t.Fail()
	}

	err = Enum("blue", en...)
	if err == nil {
		t.Fail()
	}

	err = Enum("", en...)
	if err == nil {
		t.Fail()
	}
}
