package exampleFunction

import "fmt"

func ExampleLengthRange() {
	fmt.Println(LengthRange("Ali   v"))
	fmt.Println(LengthRange("M"))
	// output:
	//7
	//7
}
