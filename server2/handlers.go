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
	log.Println("Serving:", r.URL.Path, "from", r.Host)
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
		fmt.Fprintf(w, "%s", Body)
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

func insertHandler(w http.ResponseWriter, r *http.Request) {
	parmStr := strings.Split(r.URL.Path, "/")
	fmt.Println("path", parmStr)
	if len(parmStr) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "400 bad request , Not enough arguments:"+r.URL.Path)
		return
	}
	dataest := parmStr[2]
	dataStr := parmStr[3:]
	data := make([]float64, 0)

	for _, v := range dataStr {
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			data = append(data, val)
		}
	}
	entry := Procces(dataest, data)
	err := insert(&entry)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotModified)
		Body := "There is nothing to insert!\n"
		fmt.Fprintf(w, "%s", Body)
	} else {
		w.WriteHeader(http.StatusOK)
		Body := "New Record added successfully!\n"
		fmt.Fprintf(w, "%s", Body)
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	parmStr := strings.Split(r.URL.Path, "/")
	fmt.Println("path", parmStr)
	if len(parmStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 not found: "+r.URL.Path)
		return
	}
	var body string
	dataest := parmStr[2]
	t := search(dataest)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		body = "Could not be found!\n" + dataest + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		body = fmt.Sprintf("%s %d %f %f\n", t.Name, t.Len, t.Mean, t.StdDev)
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "%s", body)

}
