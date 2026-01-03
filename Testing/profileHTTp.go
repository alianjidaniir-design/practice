package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World", r.URL.Path)
	fmt.Printf("loaded: %s\n", r.Host)
}

func mytimeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05")
	Body := "Very very good"
	fmt.Fprintf(w, "Serving: %s\n", Body)
	fmt.Fprintf(w, "Time: %s\n", t)
	fmt.Fprintf(w, "Serving %s\n", r.URL.Path)
	fmt.Printf("Serving %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Need to else argument")
	} else {
		PORT = ":" + arg[1]
		fmt.Println("Using port:", PORT)
	}
	r := http.NewServeMux()
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/time", mytimeHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	//////
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
