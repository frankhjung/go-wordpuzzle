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
