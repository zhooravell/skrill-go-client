package skrill

import "testing"

func TestDetailText_Valid(t *testing.T) {
	cn := DetailText("text")

	if err := cn.Valid(); err != nil {
		t.Error(err)
	}
}

func TestDetailText_ValidEmpty(t *testing.T) {
	cn := DetailText("")

	if err := cn.Valid(); err == nil {
		t.Fail()
	}
}
