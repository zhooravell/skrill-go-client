package skrill

import (
	"strings"
	"testing"
)

func TestCompanyName_Valid(t *testing.T) {
	cn := CompanyName("Test company")

	if err := cn.Valid(); err != nil {
		t.Error(err)
	}
}

func TestCompanyName_ValidEmpty(t *testing.T) {
	cn := CompanyName("")

	if err := cn.Valid(); err == nil {
		t.Fail()
	}
}

func TestCompanyName_ValidInvalidMaxLength(t *testing.T) {
	cn := CompanyName(strings.Repeat("a", 31))

	if err := cn.Valid(); err == nil {
		t.Fail()
	}
}
