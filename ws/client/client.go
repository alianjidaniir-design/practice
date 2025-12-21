package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var (
	SERVER       = ""
	PATH         = ""
	TIMESWAIT    = 0
	TIMESWAITMAX = 5
	in           = bufio.NewReader(os.Stdin)
)

func getinput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- result
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Enter SERVER + PATH")
		return
	}
	SERVER = args[1]
	PATH = args[2]
	fmt.Println("Connection to ", SERVER, "at", PATH)

	c1 := make(chan os.Signal, 1)
	signal.Notify(c1, os.Interrupt)

	input := make(chan string, 1)
	go getinput(input)

	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("Please give me input", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("write:", err)
				return
			}
			TIMESWAITMAX = 0
			go getinput(input)
		case <-c1:
			log.Println("Caught interrupt signal - quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}

}
