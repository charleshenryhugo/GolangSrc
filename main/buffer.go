package main

import (
	"bytes"
)

//Append () appends a new []byte data to slice,
//growing slice as needed
//slower than "append()"
func Append(slice, data []byte) []byte {
	buffer := bytes.NewBuffer(slice)
	buffer.Write(data)
	return buffer.Bytes()
}

//AppendByte () appends a new byte data to slice,
//growing slice as needed
//used like "append()"
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n >= cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}

//divide a []byte slice into two slice,
//and return a slice of 2 slices
func divSlice(buf []byte, n int) [][]byte {
	return [][]byte{buf[:n], buf[n:]}
}
