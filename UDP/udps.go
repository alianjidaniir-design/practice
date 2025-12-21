package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		fmt.Println("Usage: UDP UDP_PORT")
		return
	}
	PORT := ":" + args[1]
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()
	buffer := make([]byte, 1024)
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(rg.Intn(100))
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Println(" -- ", string(buffer[0:n-1]))
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}
		data := []byte(strconv.Itoa(int(random(1, 1001))))
		fmt.Println(string(data))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
