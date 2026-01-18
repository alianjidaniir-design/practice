package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2}
	s2 := []int{11, 22}
	s3 := []int{-1, -3}
	conCat := slices.Concat(s1, s2, s3)
	fmt.Println(conCat)

	v1 := []int{-1, 2, 5, 7, 9}
	fmt.Println("V1", v1)
	v2 := slices.Delete(v1, 0, 3)
	fmt.Println(v1)
	fmt.Println(v2)

}
