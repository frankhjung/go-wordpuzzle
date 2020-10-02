package main

import (
	"testing"
)

const (
	size      = 4
	mandatory = byte('c')
	letters   = "adevcrsoi"
)

var (
	empty byte
	test  string
)

func TestWordTooShort(t *testing.T) {
	test = "ice"
	if IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Invalid word %s. Expected false", test)
	}
}

func TestWordTooLong(t *testing.T) {
	test = "adevcrsoia"
	if IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Invalid word %s. Expected false", test)
	}
}

func TestValidWord(t *testing.T) {
	test = "voice"
	if !IsValidWord(size, mandatory, letters, test) {
		t.Errorf("Valid word %s. Expected true", test)
	}
}

func TestNotValidWord(t *testing.T) {
	test = "voicex"
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

func TestGetMandatory(t *testing.T) {
	test = "c"
	mandatory, ok := GetMandatory(test)
	if ok != nil {
		t.Errorf("Unexpected error: %t", ok)
	}
	if mandatory != 'c' {
		t.Errorf("Expected 'c' got %c", mandatory)
	}
}

func TestGetMandatoryInvalid(t *testing.T) {
	test = "xx"
	mandatory, ok := GetMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != empty {
		t.Errorf("Expected '' got %c", mandatory)
	}
}

func TestGetMandatoryNotLetter(t *testing.T) {
	test = "$"
	mandatory, ok := GetMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != empty {
		t.Errorf("Expected '' got %c", mandatory)
	}
}
func TestGetMandatoryEmpty(t *testing.T) {
	test = ""
	mandatory, ok := GetMandatory(test)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != empty {
		t.Errorf("Expected '%c': got '%c'", empty, mandatory)
	}
}
