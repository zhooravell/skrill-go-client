package skrill

import "testing"

func TestSid_Valid(t *testing.T) {
	s := Sid("test123")

	if err := s.Valid(); err != nil {
		t.Error(err)
	}
}

func TestSid_ValidEmpty(t *testing.T) {
	s := Sid("")

	if err := s.Valid(); err == nil {
		t.Error(err)
	}
}
