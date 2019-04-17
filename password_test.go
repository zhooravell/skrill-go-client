package skrill

import "testing"

func TestPassword_Valid(t *testing.T) {
	p := Password("Q1w2e3r4t5y6u7i8")

	if err := p.Valid(); err != nil {
		t.Error(err)
	}
}

func TestPassword_ValidLessThen8Char(t *testing.T) {
	p := Password("q1w2")

	if err := p.Valid(); err == nil {
		t.Error(err)
	}
}

func TestPassword_ValidLessThen1AlphabeticChar(t *testing.T) {
	p := Password("123456789")

	if err := p.Valid(); err == nil {
		t.Error(err)
	}
}

func TestPassword_ValidLessThen1NonAlphabeticChar(t *testing.T) {
	p := Password("Qwertyiop")

	if err := p.Valid(); err == nil {
		t.Error(err)
	}
}
