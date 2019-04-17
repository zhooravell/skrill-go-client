package skrill

import "testing"

func TestDetailDescription_Valid(t *testing.T) {
	cn := DetailDescription("text")

	if err := cn.Valid(); err != nil {
		t.Error(err)
	}
}

func TestDetailDescription_Empty(t *testing.T) {
	cn := DetailDescription("")

	if err := cn.Valid(); err == nil {
		t.Fail()
	}
}
