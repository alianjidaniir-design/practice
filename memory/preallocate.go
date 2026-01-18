package main

import (
	"fmt"
)

func main() {
	myslice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		myslice = append(myslice, i)
	}
	fmt.Println(myslice)

	myMap := make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("k%d", i)
		myMap[key] = i
	}
	fmt.Println(myMap)

}
