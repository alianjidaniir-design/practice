package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

func se(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(w, r)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("defaultHandler Serving: ", r.URL.Path, "from", r.Host, "with method", r.Method)
	w.WriteHeader(http.StatusNotFound)
	body := r.URL.Path + "is not supporting. but very thanks for visiting!\n"
	fmt.Fprintf(w, "%s", body)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host, "with method", r.Method)
	w.WriteHeader(http.StatusNotFound)
	body := "Method not allowed!\n"
	fmt.Fprintf(w, "%s", body)
}
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	t := time.Now().Format("2006-01-02 15:04:05")
	Body := "time is: " + t + "\n"
	fmt.Fprintf(w, "%s", Body)
}
