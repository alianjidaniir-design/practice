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

}
