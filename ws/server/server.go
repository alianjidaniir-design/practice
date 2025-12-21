package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Server!\n")
	fmt.Fprintf(w, "Please use /ws for Websocket!")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		mt, massage, err := ws.ReadMessage()
		if err != nil {
			log.Println("From", r.Host, "read:", err)
			break
		}
		log.Println("Received", string(massage))
		err = ws.WriteMessage(mt, massage)
		if err != nil {
			log.Println("WriteMessage:", err)
			break
		}
	}
}

func main() {
	arg := os.Args
	if len(arg) > 1 {
		PORT = ":" + arg[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/ws", wsHandler)
	log.Println("Listening to TCP Port ", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}

}
