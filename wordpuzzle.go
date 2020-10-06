// WordPuzzle example application to solve the nine-letter word puzzle.
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
	USAGE = `Usage: wordpuzzle -size [num] -mandatory [char] -letters [letters] [options]
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

// IsValidLetters returns true if string is 9 lowercase letters
func IsValidLetters(s string) bool {
	if len(s) != 9 {
		return false
	}
	for _, c := range s {
		if !unicode.IsLower(c) && unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsValidMandatory returns leading lowercase byte from the argument string
func IsValidMandatory(s string) (byte, error) {
	if len(s) == 1 {
		if m := rune(s[0]); unicode.IsLower(m) {
			return byte(m), nil
		}
		return 0, fmt.Errorf("expected lowercase letter got %s", s)
	}
	return 0, fmt.Errorf("%s is invalid for mandatory character", s)
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
	mandatoryChar := flag.String("mandatory", "", "Mandatory character for all words")
	size := flag.Int("size", 4, "Minimum word size (value from 1..9)")
	verboseFlag := flag.Bool("verbose", false, "Verbose mode")
	versionFlag := flag.Bool("version", false, "print wordpuzzle version")
	flag.Usage = usage

	flag.Parse()

	// print usage if no flags set
	if flag.NFlag() == 0 {
		usage()
		os.Exit(0)
	}

	// print version then exit
	if *versionFlag {
		version()
		os.Exit(0)
	}

	// test required parameter letters
	if !IsValidLetters(*letters) {
		fmt.Fprintf(os.Stderr, "Error: invalid letters %s\n", *letters)
		usage()
		os.Exit(1)
	}

	// test required parameter mandatory
	mandatory, ok := IsValidMandatory(*mandatoryChar)
	if ok != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid mandatory %s\n", *mandatoryChar)
		usage()
		os.Exit(1)
	}

	// test required parameter size
	if *size < 1 || *size > 9 {
		fmt.Fprintf(os.Stderr, "Error: invalid size %d\n", *size)
		usage()
		os.Exit(1)
	}

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
