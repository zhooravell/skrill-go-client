package skrill

import "testing"

func TestSecretWord_Valid(t *testing.T) {
	sw := SecretWord("test123")

	if err := sw.Valid(); err != nil {
		t.Error(err)
	}
}

func TestPassword_ValidGreatThen10Char(t *testing.T) {
	sw := SecretWord("test123qwer5t6y")

	if err := sw.Valid(); err == nil {
		t.Error(err)
	}
}

func TestPassword_ValidContainsSpecialCharacters(t *testing.T) {
	sw := SecretWord("test@")

	if err := sw.Valid(); err == nil {
		t.Error(err)
	}
}

func TestPassword_ValidContainsSpecialCharacters_2(t *testing.T) {
	sw := SecretWord("$test")

	if err := sw.Valid(); err == nil {
		t.Error(err)
	}
}

func TestPassword_ValidContainsSpecialCharacters_4(t *testing.T) {
	sw := SecretWord("t%st")

	if err := sw.Valid(); err == nil {
		t.Error(err)
	}
}
