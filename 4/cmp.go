package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(cmp.Compare("z", "b"))
	fmt.Println(cmp.Compare(-21, -22))
	fmt.Println(cmp.Compare(-22.233434374, -22.2343432))
	fmt.Println(cmp.Less("a", "b"))
	fmt.Println(cmp.Or("b", "a"))
}
