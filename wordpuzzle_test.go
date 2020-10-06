package main

import (
	"testing"
)

const (
	size      = 4
	mandatory = byte('c')
	letters   = "adevcrsoi"
)

func TestWordTooShort(t *testing.T) {
	test := "ice"
	if IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Invalid word %s. Expected false", test)
	}
}

func TestWordTooLong(t *testing.T) {
	test := "adevcrsoiX"
	if IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Invalid word %s. Expected false", test)
	}
}

func TestValidWord(t *testing.T) {
	test := "voice"
	if !IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Valid word %s. Expected true", test)
	}
}

func TestNotValidWord(t *testing.T) {
	test := "voicex"
	if IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Valid word %s. Expected true", test)
	}
}

func TestRemoveIndex(t *testing.T) {
	test := "abc"
	if result := RemoveIndex([]byte(test), 0); string(result) != "bc" {
		t.Errorf("RemoveIndex from %s. Expected 'bc' got %s", test, result)
	}
}

func TestLettersValid(t *testing.T) {
	test := "adevcrsoi"
	if !IsValidLetters(test) {
		t.Errorf("Valid letters %s. Expected true", test)
	}
}
func TestLettersInvalid(t *testing.T) {
	test := "adevcrSoi"
	if IsValidLetters(test) {
		t.Errorf("Invalid letters %s. Expected false", test)
	}
}
func TestLettersTooLong(t *testing.T) {
	test := "adevcrsoix"
	if IsValidLetters(test) {
		t.Errorf("Invalid number of letters %d. Expected false", len(test))
	}
}

func TestIsValidMandatoryGood(t *testing.T) {
	test := "c"
	mandatory, ok := IsValidMandatory(test)
	if ok != nil {
		t.Errorf("Unexpected error: %t", ok)
	}
	if mandatory != 'c' {
		t.Errorf("Expected 'c' got %c", mandatory)
	}
}

func TestIsValidyMandatoryBad(t *testing.T) {
	test := "xx"
	mandatory, ok := IsValidMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != 0 {
		t.Errorf("Expected '' got %c", mandatory)
	}
}

func TestIsValidyMandatoryNotLetter(t *testing.T) {
	test := "$"
	mandatory, ok := IsValidMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != 0 {
		t.Errorf("Expected '' got %c", mandatory)
	}
}
func TestIsValidyMandatoryEmpty(t *testing.T) {
	test := ""
	mandatory, ok := IsValidMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != 0 {
		t.Errorf("Expected '0': got '%c'", mandatory)
	}
}
