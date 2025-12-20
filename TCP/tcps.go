package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: tcps <host:port>")
		return
	}

	// listening to Port

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer l.Close()

	// Accept to client

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.TrimSpace(string(netData)) == "STOP" {
		fmt.Println("Stopping")
		return
	}
	fmt.Print("-> ", string(netData))
	t := time.Now()
	mytime := t.Format("2006-01-02 15:04:05") + "\n"
	c.Write([]byte(mytime))

}
