package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"time"
)

func fib01(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fib02(n-1)) + int64(fib02(n-2))
}

func fib02(n int) int {
	fn := make(map[int]int)
	for i := 0; i < n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i <= int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func N2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(os.TempDir())
	cpuFilename := path.Join(os.TempDir(), "cpuProfileCla.out")
	cpu, err := os.Create(cpuFilename)
	if err != nil {
		fmt.Println(err)
	}
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()
	total := 0
	for i := 2; i < 10000; i++ {
		n := N1(i)
		if n {
			total++
		}

	}
	fmt.Println("Total", total)
	total = 0
	for i := 2; i < 10000; i++ {
		n := N2(i)
		if n {
			total++
		}
	}
	fmt.Println("Total", total)
	for i := 1; i < 90; i++ {
		n := fib01(i)
		fmt.Println("fibonachi:", n)
	}
	for i := 1; i < 90; i++ {
		n := fib02(i)
		fmt.Println("fibonachi:", n)
	}
	fmt.Println()
	runtime.GC()

	memoryFilename := path.Join(os.TempDir(), "memoryProfileCla.out")
	memory, err := os.Create(memoryFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memory.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(50 * time.Millisecond)
	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
