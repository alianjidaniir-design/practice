package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type User2 struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u1 = User2{"usr323", "1235t45"}
var u2 = User2{"Ali", "678910"}
var u3 = User2{"", "1234"}
var u4 = User2{"Ali", "123434"}

func deleteEndpoint(server string, user User2) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest
	}
	u := bytes.NewReader(userMarshall)
	req, err := http.NewRequest(http.MethodDelete, server+deleteEndPoint, u)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error in request", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/delete returned ", string(data))
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode
}

func getEndpoint(server string, user User2) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest
	}
	u := bytes.NewReader(userMarshall)
	req, err := http.NewRequest(http.MethodGet, server+getEndPoint, u)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/get returned ", string(data))
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode
}
func addEndpoint(server string, user User2) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusInternalServerError
	}
	u := bytes.NewReader(userMarshall)
	req, err := http.NewRequest(http.MethodPost, server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func timeEndpoint(server string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+timeEndPoint, nil)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest, ""
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode, string(data)
}

func slashEndpoint(server, URL string) (int, string) {
	req, err := http.NewRequest(http.MethodPost, server+URL, nil)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest, ""
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if resp == nil {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode, string(data)
}

const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

func main() {
	a := os.Args
	if len(a) != 2 {
		fmt.Println("Wrong number of arguments")
		fmt.Println("Need : server URl")
		return
	}
	server := a[1]

	fmt.Println("/add")
	httpCode := addEndpoint(server, u4)
	if httpCode != http.StatusOK {
		fmt.Println("u4 Return code! ", httpCode)
	} else {
		fmt.Println("u4 Data added", u1, httpCode)
	}
	httpCode = addEndpoint(server, u3)
	if httpCode != http.StatusOK {
		fmt.Println("u2 Return code!", httpCode)
	} else {
		fmt.Println("u2 Data added", u2, httpCode)
	}
	//////////////////////
	fmt.Println("/delete")
	httpCode = deleteEndpoint(server, u3)
	fmt.Println("/delete u3 return code", httpCode)

	httpCode = deleteEndpoint(server, u2)
	fmt.Println("/delete u2 return code", httpCode)
	//////////////////
	fmt.Println("/get")
	httpCode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code", httpCode)

	httpCode = getEndpoint(server, u3)
	fmt.Print("/get u3 return code", httpCode)
	//////////////
	fmt.Println("/time")
	httpCode, mytime := timeEndpoint(server)
	fmt.Println("/time returned: ", httpCode, "", mytime)
	time.Sleep(1 * time.Second)
	httpCode, mytime = timeEndpoint(server)
	fmt.Println("/time returned: ", httpCode, "", mytime)
	///////////
	fmt.Println("/")
	URL := "/"
	httpCode, response := slashEndpoint(server, URL)
	fmt.Println("/slash returned: ", httpCode, " with response: ", response)
	fmt.Println("/what")

	URL = "/what"
	httpCode, response = slashEndpoint(server, URL)
	fmt.Println("/slash returned: ", httpCode, " with response: ", response)
}
