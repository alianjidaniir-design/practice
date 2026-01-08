package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func R1(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", errors.New("invalid utf8_8")
	}
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--

	}
	return string(a), nil

}
func R2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", errors.New("invalid utf8_16")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {

	str := "Sedad"
	R1ret, err := R1(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(R1ret)
	R2ret, err := R2(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(R2ret)

}
