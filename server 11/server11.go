package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serveing : %s\n", r.URL.Path)
	fmt.Println("Served", r.Host)

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05")
	Body := "Great"
	fmt.Fprintf(w, Body)
	fmt.Fprintln(w, t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arg := os.Args
	if len(arg) != 1 {
		PORT = ":" + arg[1]
	}
	fmt.Println("Using port:", PORT)
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/time", timeHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
