package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	
	"time"

	"github.com/gorilla/mux"
)

var rMux = mux.NewRouter()

var PORT = ":1234"

func main() {
	a := os.Args
	if len(a) > 1 {
		PORT = ":" + a[1]
	}
	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	rMux.NotFoundHandler = http.HandlerFunc(DefaultHandler)
	notallowed := notAllowedHandler{}
	rMux.MethodNotAllowedHandler = notallowed
	rMux.HandleFunc("/time", TimeHandler)
	getMux := rMux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/getall", GetAllHandler)
	getMux.HandleFunc("/getid/{username}", GetIDHandler)
	getMux.HandleFunc("/logged", LoggedUserHandler)
	getMux.HandleFunc("/username/{id:[0-9]+}", GetUserDataHandlern)
	putMux := rMux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", UpdataHandler)
	postMux := rMux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", AddHandler)
	postMux.HandleFunc("login", LoginHandler)
	postMux.HandleFunc("logout", LogoutHandler)
	deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/delete", DeleteHandler)
	deleteMux.HandleFunc("/username/{id:[0-9]+}", DeleteHandler)

	go func() {
		log.Printf("Listening on %s\n", port)
		err = s.ListenAndServe()
		if err != nil {
			log.Println("Error stating server: %s\n", err)
			return
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Got signal:", sig)
	time.Sleep(3 * time.Second)
	s.Shutdown(nil)

}
