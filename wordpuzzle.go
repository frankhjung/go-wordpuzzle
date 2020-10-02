package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Program version
const (
	USAGE = `Usage: wordpuzzle -size [num] -mandatory [char] -letters <letters> [options]
Solve 9-letter word puzzle.
Available options:`
	VERSION = "1.0.0"
)

// Usage prints program help
func usage() {
	fmt.Fprintln(os.Stderr, USAGE)
	flag.PrintDefaults()
	version()
}

// Version shows program version
func version() {
	fmt.Fprintln(os.Stderr, "wordpuzzle version:", VERSION)
}

// check for error and print message and exit with non-zero error code
func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", e)
		os.Exit(1)
	}
}

// GetMandatory returns leading lowercase byte from the argument string
func GetMandatory(s string) (byte, error) {
	var empty byte
	if len(s) == 1 {
		if m := rune(s[0]); unicode.IsLower(m) {
			return byte(m), nil
		}
		return empty, fmt.Errorf("expected lowercase letter got %s", s)
	}
	return empty, fmt.Errorf("%s is invalid for mandatory character", s)
}

// RemoveIndex removes value by index from a byte array
func RemoveIndex(s []byte, i int) []byte {
	return append(s[:i], s[i+1:]...)
}

// IsValidWord returns true if dictionary word is valid
func IsValidWord(size int, mandatory byte, letters string, word string) bool {
	// too small or too big so not valid
	if size > len(word) || len(word) > 9 {
		return false
	}
	// must contain mandatory letter
	if !strings.Contains(word, string(mandatory)) {
		return false
	}
	// test that all letters in word are valid
	working := []byte(letters)
	for _, letter := range word {
		if i := bytes.IndexByte(working, byte(letter)); i > -1 {
			// remove from working list if present
			working = RemoveIndex(working, i)
		} else {
			// letter not in working list so not a valid word
			return false
		}
	}
	return true
}

// main process commandline arguments and filter valid words from dictionary
func main() {

	dictionaryPath := flag.String("dictionary", "dictionary", "Dictionary to read words from")
	letters := flag.String("letters", "", "Nine letters to make words")
	mandatoryString := flag.String("mandatory", "", "Mandatory character for all words")
	size := flag.Int("size", 4, "Minimum word size (value from 1..9)")
	verboseFlag := flag.Bool("verbose", false, "Verbose mode")
	versionFlag := flag.Bool("version", false, "print wordpuzzle version")
	flag.Usage = usage

	flag.Parse()

	// print version then exit
	if *versionFlag {
		version()
		os.Exit(0)
	}

	// test mandatory parameters
	if *size < 1 || *size > 9 || *mandatoryString == "" || *letters == "" {
		usage()
		os.Exit(1)
	}

	// mandatory should be a lower case letter
	mandatory, ok := GetMandatory(*mandatoryString)
	check(ok)

	// open dictionary
	file, ok := os.Open(*dictionaryPath)
	check(ok)
	defer file.Close()

	// if verbose show all parameters
	if *verboseFlag {
		fmt.Println("dictionary:", *dictionaryPath)
		fmt.Println("letters:", *letters)
		fmt.Println("mandatory:", mandatory)
		fmt.Println("size:", *size)
		fmt.Println("other arguments:", flag.Args())
	}

	// read all words from dictionary, printing only valid words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if word := scanner.Text(); IsValidWord(*size, mandatory, *letters, word) {
			fmt.Println(word)
		}
	}
	check(scanner.Err())
}
