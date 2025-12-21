package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func HandleConnection(c net.Conn, myCount int) {
	fmt.Print(",")
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)

		counter := "Client number: " + strconv.Itoa(myCount) + "\n"
		c.Write([]byte(string(counter)))

	}

	defer c.Close()

}

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide arg")

		os.Exit(5)
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConnection(c, count)
		count++
		break
	}

}
