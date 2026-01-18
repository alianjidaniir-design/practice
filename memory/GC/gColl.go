package main

import (
	"fmt"
	"runtime"
	"time"
)

func printState(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloce", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
}
func main() {
	var mem runtime.MemStats
	printState(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("failed")
		}
	}

	printState(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("failed")
		}
		time.Sleep(5 * time.Second)
	}
	printState(mem)

}
