package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}
type S2 struct {
	F1 int
	F2 S1
}

type Secret struct {
	searchValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("secret type")
	case Entry:
		fmt.Println("entry type")
	default:
		fmt.Println("No supported type:", T)
	}
}

func learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func Print(s interface{}) {
	fmt.Println(s)
}

func sw(s interface{}) {
	switch T := s.(type) {
	case int:
		fmt.Println(T * T)
	case string:
		fmt.Println("I love : ", T)
	case float64:
		fmt.Println(T - 1.21)
	default:
		fmt.Print("No supported type:", T)
		fmt.Printf(" %T\n", T)
	}
}

func returnNumber() interface{} {
	return 10
}

func main() {
	v1 := S1{F1: 14, F2: "good morning"}
	v2 := S2{F1: 35, F2: v1}
	Print(v2)
	Print(v1)
	Print(v1.F2)
	Print(1234)
	Print([]float64{1.343, 2.43, 3.34})
	A1 := Secret{searchValue: "Hello World"}
	A2 := Entry{F1: 99, F2: "good morning", F3: Secret{searchValue: "new cars"}}
	Teststruct(A1)
	Teststruct(A2)
	learn(A1)
	learn(A2)
	Teststruct(3)
	learn(3)
	sw(12)
	sw("Ali")
	sw(0)
	sw(12.45)
	sw([]int{1, 2, 7, 3})
	aa := returnNumber()
	number, ok := aa.(int)
	if ok {
		fmt.Println("type assertion successful")
	} else {
		fmt.Println("type assertion failed")
	}
	number++
	fmt.Println(number)

	value, ok := aa.(int64)
	if ok {
		fmt.Println("type assertion successful", value)
	} else {
		fmt.Println("type assertion failed", value)
	}

	i := aa.(int)
	fmt.Println(i)
	_ = aa.(bool)

}
