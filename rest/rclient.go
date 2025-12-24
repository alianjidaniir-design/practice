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

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u8 = User{"fdg", "1234"}
var u7 = User{"3456", "123434"}
var u12 = User{"356546", "123454356"}

func deleteEndpoint(server string, user User) int {
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
		Timeout: 15 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error in request", err)
	}
	defer resp.Body.Close()

	if resp == nil {
		return http.StatusBadRequest
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/delete returned ", string(data))
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode
}

func getEndpoint(server string, user User) int {
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
		Timeout: 15 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer resp.Body.Close()
	if resp == nil {
		return resp.StatusCode
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Print("/get returned ", string(data))
	if err != nil {
		fmt.Println("Error", err)
	}
	return resp.StatusCode
}
func addEndpoint(server string, user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in request", err)
		return http.StatusBadRequest
	}
	u := bytes.NewReader(userMarshall)
	req, err := http.NewRequest(http.MethodPost, server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error", err)
		return http.StatusBadRequest
	}
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 15 * time.Second,
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
		Timeout: 15 * time.Second,
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
		Timeout: 15 * time.Second,
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
	httpCode := addEndpoint(server, u8)
	if httpCode != http.StatusOK {
		fmt.Println("u8 Return code! ", httpCode)
	} else {
		fmt.Println("u8 Data added", u8, httpCode)
	}
	httpCode = addEndpoint(server, u7)
	if httpCode != http.StatusOK {
		fmt.Println("u7 Return code!", httpCode)
	} else {
		fmt.Println("u7 Data added", u7, httpCode)
	}
	httpCode = addEndpoint(server, u12)
	if httpCode != http.StatusOK {
		fmt.Println("u12 Return code!", httpCode)
	} else {
		fmt.Println("u12 Data added", u12, httpCode)
	}
	//////////////////////
	//////////////////
	fmt.Println("/get")
	httpCode = getEndpoint(server, u8)
	fmt.Println("/get u8 return code", httpCode)
	httpCode = getEndpoint(server, u7)
	fmt.Println("/get u7 return code", httpCode)

	httpCode = getEndpoint(server, u12)
	fmt.Print("/get u12 return code", httpCode)
	/////////////
	fmt.Println("/delete")
	httpCode = deleteEndpoint(server, u8)
	fmt.Println("/delete u8 return code", httpCode)
	httpCode = deleteEndpoint(server, u7)
	fmt.Println("/delete u7 return code", httpCode)
	httpCode = deleteEndpoint(server, u12)
	fmt.Println("/delete u12 return code", httpCode)

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
