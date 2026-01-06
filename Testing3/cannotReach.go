package main

import (
	"fmt"
)

func S2() {
	return
	fmt.Println("Hello World")
}

func S1() {
	fmt.Println("In S1()")
	return
	fmt.Println("Leaving S(1)")
}

func main() {
	S1()
}
