package main

import (
	"fmt"
	"os"
	"path"
	"runtime/trace"
	"time"
)

func main() {
	filepath := path.Join(os.TempDir(), "trace.out")
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("trace.out is nil")
		}
	}
	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("trace.out is nil")
		}
		time.Sleep(time.Millisecond)
	}

}
