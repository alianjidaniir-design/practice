package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	src := rand.NewSource(seed)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Uint64())
	}
}
