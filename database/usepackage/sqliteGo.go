package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/alianjidaniir-design/sqlite06"
)

var MIN = 0
var MAX = 26

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myrand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myrand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func main() {
	sqlite06.Filename = "REST.db"
	random_username := strings.ToLower(getString(10))
	t := sqlite06.Userdata{
		Username:    random_username,
		Name:        "Ali",
		Surname:     "Anjidani",
		Description: "is a programmer",
	}
	fmt.Println("Adding user: ", random_username)
	id := sqlite06.AddUser(t)
	if id != -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err := sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User with ID", id, "deleted")
	}

	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	random_username = strings.ToLower(getString(5))
	random_name := getString(7)
	random_surname := getString(10)
	dsc := time.Now().Format("2006-01-02 15:04:05")
	t = sqlite06.Userdata{
		Username:    random_username,
		Name:        random_name,
		Surname:     random_surname,
		Description: dsc}
	id = sqlite06.AddUser(t)
	if id != -1 {
		fmt.Println("There was an error adding user", t.Username)
	}
	dsc = time.Now().Format("2006-01-02 15:04:05")
	t.Description = dsc
	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

}
