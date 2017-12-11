package main

//. Matches any character except a newline
//[.] Matches .
//* Matches the preceding pattern element >=0 times, the same as {1,}
//+ Matches the proceding pattern element >0 times, the same as {0,}
//? Matches the proceding pattern element 0 or 1 time, the same as {0,1}
//{M,N} Matches the proceding pattern element at least M times and at most N times
//^ Matches the beginning of a line or string
//$ Matches the end of a line or string
//\A Matches the beginning of a string(but not an internal line)
//\z Matches the end of a string(but not an internal line)
//\w the same as [A-Za-z0-9_]
//\W the same as [^A-Za-z0-9_]
//\d the same as [0-9]
//\D the same as [^0-9]
//\s Matches a whitespace charactor
//\S Matches anything but a whitespace

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var (
	digitRegexp = regexp.MustCompile(`[0-9]+`)
)

//findFileDigits finds alldigits in a file
//and return a []byte
func findFileDigits(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

//findStrDigits finds all digits in a string
//and return a []byte
func findStrDigits(str string) []byte {
	b := digitRegexp.FindAll([]byte(str), len(str))
	c := make([]byte, 0)

	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

//doubleAllFloats doubles all the float-type-digits
//found in searchIn
func doubleAllFloats(searchIn string) string {
	f := func(s string) string {
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			v = 0.0
		}
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	pat := `[0-9]+[.][0-9]+`

	if ok, _ := regexp.MatchString(pat, searchIn); !ok {
		fmt.Println("Match not found!")
		return ""
	}
	dRegexp := regexp.MustCompile(pat)
	str1 := dRegexp.ReplaceAllString(searchIn, "####.####")
	str2 := dRegexp.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str1)
	fmt.Println(str2)
	return str2
}

func scanBytes(bytes []byte) {
	for _, v := range bytes {
		fmt.Printf("%c ", v)
	}
}
