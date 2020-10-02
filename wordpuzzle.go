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

const VERSION = "1.0.0"

// GetMandatory returns mandatory byte from command line argument string
func GetMandatory(in string) (byte, error) {
	var empty byte
	if len(in) > 0 {
		if m := rune(in[0]); unicode.IsLetter(m) && unicode.IsLower(m) {
			return byte(m), nil
		}
		return empty, fmt.Errorf("expected lowercase letter got %s", in)

	} else {
		return empty, fmt.Errorf("flag requires a parameter")
	}
}

// check for error and print message and exit with non-zero error code
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

// RemoveIndex removes value by index from a byte array
func RemoveIndex(s []byte, index int) []byte {
	return append(s[:index], s[index+1:]...)
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

	dictionary := flag.String("dictionary", "dictionary", "Dictionary to read words from")
	letters := flag.String("letters", "", "Nine letters to make words")
	mandatoryString := flag.String("mandatory", "", "Mandatory character for all words")
	size := flag.Int("size", 4, "Minimum word size (value from 1..9)")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	version := flag.Bool("version", false, "print wordpuzzle version")

	flag.Parse()

	// print version then exit
	if *version {
		fmt.Println("Version:", VERSION)
		os.Exit(0)
	}

	// test mandatory parameters
	if *size < 1 || *size > 9 || *mandatoryString == "" || *letters == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// mandatory should be a lower case letter
	mandatory, ok := GetMandatory(*mandatoryString)
	check(ok)

	// open dictionary
	file, ok := os.Open(*dictionary)
	check(ok)
	defer file.Close()

	// if verbose show all parameters
	if *verbose {
		fmt.Println("dictionary:", *dictionary)
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
