package main

import (
	"fmt"
	"time"
)

func createSlice() []int {
	return make([]int, 1000000)
}

func getValue(s []int) []int {
	val := s[:3]
	return val
}

func main() {
	for i := 0; i < 15; i++ {
		mmm := createSlice()
		val := getValue(mmm)
		fmt.Print(len(val), " ")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()
}
