package main

import (
	"fmt"
	"unicode"
)

//'\u000A' is '\n'

func char() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00001234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) //integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) //character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) //UTF-8 bytes
	if unicode.IsLetter(rune(ch3)) {
		fmt.Printf("%c is a letter", ch3)
	} else {
		fmt.Printf("%c is not a letter", ch3)
	}
}

//replace r with '?' if r is non-ASCII
func replUnASCII(r rune) rune {
	if r > 255 {
		return '?'
	}
	return r
}

func isAscii(c rune) bool {
	if c > 255 {
		return false
	}
	return true
}
