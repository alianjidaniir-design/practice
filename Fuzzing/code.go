package main

import (
	"fmt"
)

func AddInt(a, b int) int {
	for i := 0; i < a; i++ {
		b = b + i
	}
	return b
}
func main() {
	fmt.Println(AddInt(6, 4))
}
