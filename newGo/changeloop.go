package main

import (
	"fmt"
	"time"
)

func main() {

	for x := range 3 {
		fmt.Print(" ", x)
	}

	fmt.Println()

	value := []int{1, 3, 4, 6, 9}

	for _, val := range value {
		go func() {
			fmt.Print(" ", val)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}
