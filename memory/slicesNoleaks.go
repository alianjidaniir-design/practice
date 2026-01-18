package main

import (
	"fmt"
	"time"
)

func createSlice2() []int {

	return make([]int, 1000000)
}

func getValue2(s []int) []int {
	val2 := make([]int, 3)
	copy(val2, s)
	return val2
}

func main() {
	for i := 0; i < 15; i++ {
		mmm := createSlice2()
		val := getValue2(mmm)
		fmt.Print(len(val), " ")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()
}
