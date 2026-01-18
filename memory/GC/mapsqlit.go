package main

import "runtime"

func main() {
	var N = 80000000
	sqlit := make([]map[int]int, 2000)
	for i := range sqlit {
		sqlit[i] = make(map[int]int)
	}
	for i := 0; i < N; i++ {
		value := int(i)
		sqlit[i%2000][value] = value
	}
	runtime.GC()
	_ = sqlit[0]
}
