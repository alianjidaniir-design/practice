package main

import (
	"runtime"
)

type Data struct {
	i, j int
}

func main() {
	var N = 80000000
	var structure []Data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, Data{value, value})
	}
	runtime.GC()
	_ = structure[0]
}
