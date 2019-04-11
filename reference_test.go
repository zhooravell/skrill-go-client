package skrill

import "testing"

func TestReference_Valid(t *testing.T) {
	ref := Reference("123")

	if err := ref.Valid(); err != nil {
		t.Error(err)
	}
}

func TestReference_ValidEmpty(t *testing.T) {
	ref := Reference("")

	if err := ref.Valid(); err == nil {
		t.Fail()
	}
}
