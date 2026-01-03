package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func main() {
	Arg := os.Args
	if len(Arg) != 2 {
		fmt.Printf("Usage: URL\n")
		return
	}
	URL := Arg[1]
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS SINFO: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Request data from server!")
	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
}
