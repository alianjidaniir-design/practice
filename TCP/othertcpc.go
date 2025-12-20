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
	if len(args) <= 1 {
		fmt.Println("Please Usage a server:Port string!")
		return
	}
	connect := args[1]
	tcpADDr, err := net.ResolveTCPAddr("tcp", connect)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpADDr)
	if err != nil {
		fmt.Println("DialTCP:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("client >> :  ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		massage, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("server ->>: " + massage)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("End")
			return
		}
	}

}
