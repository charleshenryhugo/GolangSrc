package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func countCharacters() {
	str1 := "こんにちは"
	chars := utf8.RuneCountInString(str1)
	fmt.Println(str1[0:3])
	for i := 0; i < chars; i++ {
		fmt.Printf("%U: %c, length: %d\n", str1[i], str1[i], len(str1[0:i]))
	}
}

func stringDemo() {
	str := "Hi, Hi, haha, Hi, ha,ji,Hi"
	var old, new string = "Hi", "Hello"
	strnew := strings.Replace(str, old, new, -1)
	fmt.Println(strnew)
}

func stringSplitJoin(str string) {
	if len(str) == 0 {
		str = "The quick brown fox jumps over the lazy dog"
	}
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice by white space: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	sl2 := strings.Split(str, "o")
	fmt.Printf("Splitted in slice by letter \"o\": %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s -", val)
	}
	fmt.Println()

	sl3 := strings.Join(sl, "|")
	fmt.Printf("sl joined by \"|\": %s\n", sl3)
}

func stringConversion(str string) (int, error) {
	//fmt.Printf("The size of Int is : %d \n", strconv.IntSize)
	an, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Failed to conversion the string to int %v", err)
		return an, err
	}

	fmt.Printf("The size of Int is : %d \n", strconv.IntSize)
	fmt.Printf("An equals to : %d \n", an)
	an += 5
	NewS := strconv.Itoa(an)
	fmt.Printf("The new string is %s \n", NewS)
	return an, err
}

func funcParameter() (int, int) {
	str := "hahaha hahaha hahaha"
	index1 := strings.IndexFunc(str, unicode.IsSpace)
	fmt.Printf("the first space in \"hahaha hahaha hahaha\" is at %d\n", index1)
	index2 := strings.IndexFunc(str, isAscii)
	return index1, index2
}

//replace all non-ASCII characters in s with '?'
func replUnASCIIStr(s string) string {
	return strings.Map(replUnASCII, s)
}

func splitString(str string, ix int) (string, string) {
	if ix < 0 {
		return "", ""
	}
	if ix >= len(str) {
		return str, ""
	}
	return str[:ix], str[ix:]
}

func invString(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	len := len(runes)
	for ix := 0; ix <= (len-1)/2; ix++ {
		runes[ix], runes[len-1-ix] = runes[len-1-ix], runes[ix]
	}
	return string(runes)
}

func findDifRunes(runes []rune) []rune {
	len := len(runes)
	difRunes := make([]rune, 0, len)
	if len <= 1 {
		return difRunes
	}

	for ix := 1; ix < len; ix++ {
		if runes[ix] != runes[ix-1] {
			difRunes = append(difRunes, runes[ix])
		}
	}
	return difRunes
}
