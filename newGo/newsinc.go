package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func initValue() {
	x = 5
}
func main() {
	function := sync.OnceFunc(initValue)
	for i := 0; i < 10; i++ {
		go function()
		fmt.Println(function)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		x = x + 2
	}
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		go function()
	}
	time.Sleep(time.Second)
	fmt.Println(x)
}
