package skrill

import "testing"

func TestSignature_Valid(t *testing.T) {
	s := Signature("QWERTY")

	if err := s.Valid(); err != nil {
		t.Error(err)
	}
}

func TestSignature_ValidEmpty(t *testing.T) {
	s := Signature("")

	if err := s.Valid(); err == nil {
		t.Error(err)
	}
}

func TestSignature_ValidUpper(t *testing.T) {
	s := Signature("QwerQ")

	if err := s.Valid(); err == nil {
		t.Error(err)
	}
}
