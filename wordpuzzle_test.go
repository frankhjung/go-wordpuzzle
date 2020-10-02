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
	word := "ice"
	if IsValidWord(size, mandatory, letters, word) {
		t.Errorf("Invalid word %s. Expected false", word)
	}
}

func TestWordTooLong(t *testing.T) {
	word := "adevcrsoia"
	if IsValidWord(size, mandatory, letters, word) {
		t.Errorf("Invalid word %s. Expected false", word)
	}
}

func TestValidWord(t *testing.T) {
	word := "voice"
	if !IsValidWord(size, mandatory, letters, word) {
		t.Errorf("Valid word %s. Expected true", word)
	}
}

func TestNotValidWord(t *testing.T) {
	word := "voicex"
	if IsValidWord(size, mandatory, letters, word) {
		t.Errorf("Valid word %s. Expected true", word)
	}
}

func TestRemoveIndex(t *testing.T) {
	word := []byte("abc")
	if result := RemoveIndex(word, 0); string(result) != "bc" {
		t.Errorf("RemoveIndex from %s. Expected 'bc' got %s", word, result)
	}
}

func TestGetMandatory(t *testing.T) {
	parameter := "cd"
	mandatory, ok := GetMandatory(parameter)
	if ok != nil {
		t.Errorf("Unexpected error: %t", ok)
	}
	if mandatory != 'c' {
		t.Errorf("Expected 'c' got %c", mandatory)
	}
}

func TestGetMandatoryEmpty(t *testing.T) {
	var parameter string
	var empty byte
	mandatory, ok := GetMandatory(parameter)
	if ok == nil {
		t.Errorf("Expected error: got %v and %c", ok, mandatory)
	}
	if mandatory != empty {
		t.Errorf("Expected '%c': got '%c'", empty, mandatory)
	}
}
