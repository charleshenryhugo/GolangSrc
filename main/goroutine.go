package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func chanDeclare() {
	//declare a channel for int
	chInt := make(chan int)
	//channel for string
	chStr := make(chan string)

	//channel for int channels
	chanOfchans := make(chan chan int)
	//channel for func
	chanFunc := make(chan func())

	typeCheck(chInt, chStr, chanOfchans, chanFunc)

}

func goroutineTest() {
	Prime()
}

func Prime() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		log.Println("Prime() received ", prime)
		fmt.Println(prime)
		ch1 := make(chan int)
		go primeFilter(ch, ch1, prime)
		ch = ch1
	}
}

//Send sequence 2,3,4,... to channel "ch"
func generate(ch chan int) {
	for i := 2; i < 100; i++ {
		ch <- i
		log.Println("generate() :", i)
	}
}

//copy the values from channel 'in' to channel 'out',
//removing those divisible by 'prime'
func primeFilter(in, out chan int, prime int) {
	for {
		i := <-in
		log.Println("primeFilter() for ", prime, "received ", i)
		if i%prime != 0 {
			out <- i
			log.Println("primeFilter() output ", i)
		}
	}
}

//return a channel
func produce() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 1000; i++ {
			ch <- i
			log.Println("produce: ", i)
			time.Sleep(1e9)
		}
	}()
	return ch
}

func consume(ch chan int) {
	go func() {
		for n := range ch {
			log.Println("consume: ", n)
		}
	}()
}

func f1(ch chan interface{}) {
	log.Println(<-ch)
}

func longWait() {
	log.Println("Beginning longWait()")
	time.Sleep(5 * 1e9)
	log.Println("End of longWait()")
}

func shortWait() {
	log.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)
	log.Println("End of shortWait()")
}

func sendData(ch chan interface{}, chEnd chan bool) {
	for i := 0; i < 20; i++ {
		ch <- i
		log.Println("send:", i)
	}
}
func recvData1(ch chan interface{}, chEnd chan bool) {
	for i := 0; i < 20; i++ {
		log.Println("receiver1 received :", <-ch)
	}
	chEnd <- true
}

func recvData2(ch chan interface{}) {

	for {
		log.Println("receiver2 received :", <-ch)
	}
}

func paraSort(arr []int) {
	done := make(chan bool)
	doSort := func(a []int) {
		sort.Ints(a)
		done <- true //sort completed, send a signal
		log.Println("One sort finished")
	}

	len := len(arr)
	go doSort(arr[0 : len/2])
	go doSort(arr[len/2 : len])
	<-done
	<-done
	//log.Println(arr)
}
