package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"Ali": 2468}
	m["Food"] = 15
	fmt.Println("Before clear: ", m)
	clear(m)
	fmt.Println("After clear: ", m)
	s := make([]int, 0, 10)
	for i := 0; i < 7; i++ {
		s = append(s, i)
	}
	fmt.Println("Before clear: ", s, len(s), cap(s))
	clear(s)
	fmt.Println("After clear: ", s, len(s), cap(s))
}
