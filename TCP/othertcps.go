package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Please enter a port number!")
		return
	}
	SERVER := "localhost " + ":" + args[1]
	s, err := net.Listen("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

}
