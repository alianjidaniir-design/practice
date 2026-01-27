package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/alianjidaniir-design/sqlite06"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {

	sqlite06.Filename = "test2.db"

	data, err := sqlite06.ListUsers()
	if err != nil {
		fmt.Println("ListUsers():", err)
		return
	}

	if len(data) != 0 {
		for _, v := range data {
			fmt.Println(v)
		}
	}

	randomUsername := strings.ToLower(getString(5))

	t := sqlite06.Userdata{
		Username:    randomUsername,
		Name:        "Ali",
		Surname:     "Asad",
		Description: "This is me!"}

	fmt.Println("Adding username:", randomUsername)
	id := sqlite06.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User with ID", id, "deleted!")
	}

	// Trying to delete the same user again!
	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	randomUsername = strings.ToLower(getString(5))
	randomName := getString(7)
	randomSurname := getString(10)
	dsc := time.Now().Format(time.RFC1123)

	t = sqlite06.Userdata{
		Username:    randomUsername,
		Name:        randomName,
		Surname:     randomSurname,
		Description: dsc}

	id = sqlite06.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	dsc = time.Now().Format(time.RFC1123)
	t.Description = dsc

	err = sqlite06.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
	os.Remove(sqlite06.Filename)

}
