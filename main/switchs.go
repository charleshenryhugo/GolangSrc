package main

import "fmt"

func fallThrough1() {
	k := 6
	switch {
	case k == 4:
		fmt.Println("k <= 4")
		fallthrough
	case k == 5:
		fmt.Println("k <= 5")
		fallthrough
	case k == 6:
		fmt.Println("k <= 6")
		fallthrough
	case k == 7:
		fmt.Println("k <= 7")
		fallthrough
	case k == 8:
		fmt.Println("k <= 8")
		fallthrough
	default:
	}
}

func fallThrough2() {
	k := 6
	switch k {
	case 4:
		fmt.Println("k <= 4")
		fallthrough
	case 5:
		fmt.Println("k <= 5")
		fallthrough
	case 6:
		fmt.Println("k <= 6")
		fallthrough
	case 7:
		fmt.Println("k <= 7")
		fallthrough
	case 8:
		fmt.Println("k <= 8")
		fallthrough
	default:
	}
}

func fallThrough3() {
	switch k := 6; k {
	case 4:
		fmt.Println("k <= 4")
		fallthrough
	case 5:
		fmt.Println("k <= 5")
		fallthrough
	case 6:
		fmt.Println("k <= 6")
		fallthrough
	case 7:
		fmt.Println("k <= 7")
		fallthrough
	case 8:
		fmt.Println("k <= 8")
		fallthrough
	default:
	}
}
