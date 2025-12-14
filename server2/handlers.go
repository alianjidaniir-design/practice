package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":1234"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving")
	w.WriteHeader(http.StatusOK)
	body := "Thanks for your request!\n"
	fmt.Fprintf(w, body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	parmStr := strings.Split(r.URL.Path, "/")
	fmt.Println("path", parmStr)
	if len(parmStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 not found")
		return
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	dataest := parmStr[2]
	err := DElete(dataest)
	if err != nil {
		fmt.Println(err)
		Body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "%s", Body)
		return
	}
	body := dataest + "deleted\n!"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", body)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := list()
	fmt.Fprintf(w, "%s", body)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := fmt.Sprintf("Total entries: %d\n", len(data))
	fmt.Fprintf(w, "%s", body)
}
