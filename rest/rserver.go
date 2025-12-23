package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var user User
var PORT = ":1234"
var DATA = make(map[string]string)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusNotFound)
	body := "Thanks\n"
	fmt.Fprintln(w, "%s", body)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format("2006-01-02 15:04:05")
	body := "The currect time is: " + t + "\n"
	fmt.Fprintln(w, "%s", body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "%s\n", "Method Not Allowed")
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	if user.Username == "" {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	DATA["username"] = user.Username
	log.Println(user.Username)
	w.WriteHeader(http.StatusCreated)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "%s\n", "Method Not Allowed")
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	_, ok := DATA[user.Username]
	if ok && user.Username != "" {
		log.Println("Found!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not Found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Error:", http.StatusNotFound)
	}
	return

}
