package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func (h notAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler(w, r)
}

type notAllowedHandler struct{}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handler)
	notAllowed := notAllowedHandler{}
	r.MethodNotAllowedHandler = notAllowed
	getMux := r.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", handler)
	getMux.HandleFunc("/getall", handler)
	getMux.HandleFunc("/getid", handler)
	getMux.HandleFunc("/logged", handler)
	getMux.HandleFunc("/username/{id:[0-9]+}", handler)

	putMux := r.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/add", handler)
	putMux.HandleFunc("/login", handler)
	putMux.HandleFunc("/logout", handler)

	deleteMux := r.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", handler)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("pt", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("pr", pathRegexp)
		}
		qt, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("qt", strings.Join(qt, ","))
		}
		qren, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("qren", strings.Join(qren, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("methods", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/", r)
}
