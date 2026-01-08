package main

import (
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	const nGo = "/sched/goroutines:goroutines"

	getMeteric := make([]metrics.Sample, 1)
	getMeteric[0].Name = nGo
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(4 * time.Second)
		}()
		metrics.Read(getMeteric)
		if getMeteric[0].Value.Kind() == metrics.KindBad {
			fmt.Printf("metric %q no longer supported", nGo)

		}
		mVal := getMeteric[0].Value.Uint64()
		fmt.Println("Number of goroutines:", mVal)
	}
	wg.Wait()
	metrics.Read(getMeteric)
	mVal := getMeteric[0].Value.Uint64()
	fmt.Println("Before exiting Number of goroutines:", mVal)

}
