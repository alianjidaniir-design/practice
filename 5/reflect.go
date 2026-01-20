package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	A := T{12, "Ali", 32.45}
	fmt.Println(A)
	r := reflect.ValueOf(&A).Elem()
	fmt.Println(r.String())
	typeof := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tf := typeof.Field(i).Name
		fmt.Println(i, f, tf)
		k := reflect.TypeOf(f).Kind()
		fmt.Println(k, f)
		k2 := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k2 == reflect.Int {
			r.Field(i).SetInt(-9100)
		} else if k2 == reflect.Float64 {
			r.Field(i).SetFloat(-9100.5345)
		} else if k2 == reflect.String {
			r.Field(i).SetString("Good bye")
		}
	}
	fmt.Println("A", A)

}
