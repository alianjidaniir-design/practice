package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("You are Using", runtime.GOOS, " ")
	fmt.Println("on a(n)", runtime.GOARCH, "machine")
	fmt.Println("with Go version)", runtime.Version())
}
