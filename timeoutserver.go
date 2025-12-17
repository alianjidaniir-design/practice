package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready for serving %s\n", r.URL.Path)
	fmt.Println("Served ", r.Host)
	fmt.Println(r.Method, r.URL.Path)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05")
	Body := "this is very Great"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "Ready for serving %s\n", r.URL.Path)
	fmt.Printf("Served time:%s\n ", r.Host)
}
func main() {
	PORT := ":8001"
	arg := os.Args
	if len(arg) != 1 {
		PORT = ":" + arg[1]
	}
	fmt.Println("This is port number :", PORT)

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
