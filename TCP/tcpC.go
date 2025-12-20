package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter host:port.")
		return
	}

	// for connection client to server
	connect := args[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		line, _ := reader.ReadString('\n')
		fmt.Fprintf(c, "%s\n", line)

		massage, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(" ->: " + massage)
		if strings.TrimSpace(string(line)) == "STOP" {
			fmt.Println("TCP client exiting... ")
			return
		}
	}
}
