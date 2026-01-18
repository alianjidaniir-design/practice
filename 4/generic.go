package main

import "fmt"

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type All int
type All2 interface {
	~int
}

func AddElements[T All2](s []T) T {
	sum := T(0)
	for _, v := range s {
		sum += v
	}
	return sum
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	println()
}

func comp[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func Numeri[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}
func f3[S ~[]E, E any](x S) int {
	return len(x)
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(Numeri(1, 2))
	fmt.Println(Numeri(1.111212122, 1.11211212))
	fmt.Println(Numeri(4.32, 4.23))
	fmt.Println(Numeri(-2, -2))
	fmt.Println(Numeri(-2, -3))
	s := []int{1, 2, 3, 4, 5}
	PrintSlice(s)
	s2 := []string{"ali", "anjidani", "khoshtip"}
	PrintSlice(s2)
	s3 := []any{1, "ali", 3, -29}
	PrintSlice(s3)
	x1 := comp(-2, -3)
	fmt.Println(x1)
	fmt.Println(comp("Ali", "Ali"))
	fmt.Println(x1)
	fmt.Println(comp("Anjidani", "Khoshtip"))
	fmt.Println(comp(4.3232323, 4.32231323))
	c := []All{1, 2, 3, 4, 5}
	fmt.Println(AddElements(c))
	fmt.Println(f2(x), f3(x))
	fmt.Println(f3([]string{"Ali", "Anjidani", "Khoshtip"}))
}
