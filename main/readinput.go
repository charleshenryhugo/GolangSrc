package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput1() {
	var (
		firstName, lastName string
		s                   string
		i                   int
		f                   float64
		input               = "56.12 / 5212 / Go"
		format              = "%f / %d / %s"
	)

	fmt.Print("Enter name: ")
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Println("Hello! ", i, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read:", f, i, s)
}

func readInput2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input something: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("the input is", input)
	}
}

//read a text from keyboard input
//until letter c
//and output(return) numbers of (letters, words and lines)
func wordLetterCount(c byte) (nletter int, nword int, nline int) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter something: ")
	bytes, err := inputReader.ReadBytes(c)
	if err != nil {
		fmt.Println(err)
		return -1, -1, -1
	}
	var (
		nByte   = len(bytes)
		nSpace  int
		nLetter int
		nWord   int
		nLine   = 1
		flag    = true
	)
	for _, v := range bytes {
		if v == ' ' {
			flag = true
			nSpace++
		} else if v == '\n' {
			flag = true
			nLine++
		} else {
			if flag {
				nWord++
			}
			flag = false
		}
	}
	nLetter = nByte - nLine + 1
	fmt.Printf("nLetters = %d, nWord = %d, nLine = %d\n", nLetter, nWord, nLine)
	return nLetter, nWord, nLine
}
