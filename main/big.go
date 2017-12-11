package main

import (
	"fmt"
	"math"
	"math/big"
)

//bigIntCal shows some calculations with bigInts
func bigIntCal() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Println("Big Int:", ip)
}

//bigRatCal shows some calculations with bigRats
func bigRatCal() {
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Println("Big Rat", rq)
}
