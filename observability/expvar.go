package main

import (
	"expvar"
	"fmt"
	"net/http"
)

func main() {

	intVar := expvar.NewInt("intVar")
	intVar.Set(1234)
	expvar.Publish("customFunction", expvar.Func(func() interface{} {
		return "Hi from Mastering Go"
	}))
	http.Handle("/debug/expvars", expvar.Handler())
	go func() {
		fmt.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error listening on port 8080", err)
		}
	}()
	intVar.Add(10)
	select {}
}
