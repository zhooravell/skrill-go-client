package skrill

import "testing"

func TestLanguage_Valid(t *testing.T) {
	l := Language("EN")

	if err := l.Valid(); err != nil {
		t.Error(err)
	}
}

func TestLanguage_ValidEmpty(t *testing.T) {
	l := Language("UA")

	if err := l.Valid(); err == nil {
		t.Fail()
	}
}
