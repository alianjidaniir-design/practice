package main

import (
	"fmt"
	"os"
	"strconv"
)

func dd(x int) (int, int) {
	return x * x, 5 * x
}
func sortTwo(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

func minMax(x, y int) (min int, max int) {
	if x < y {
		min = x
		max = y
		return min, max
	}
	min = y
	max = x
	return min, max
}

func main() {

	jsd := os.Args
	if len(jsd) == 3 {
		a1, _ := strconv.Atoi(jsd[1])
		a2, _ := strconv.Atoi(jsd[2])
		fmt.Println(minMax(a1, a2))
	}

	n := 123
	d, s := dd(n)
	fmt.Println(d, s)
	fmt.Println(sortTwo(d, s))
}
