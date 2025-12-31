package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

func se(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(w, r)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("defaultHandler Serving: ", r.URL.Path, "from", r.Host, "with method", r.Method)
	w.WriteHeader(http.StatusNotFound)
	body := r.URL.Path + "is not supporting. but very thanks for visiting!\n"
	fmt.Fprintf(w, "%s", body)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host, "with method", r.Method)
	w.WriteHeader(http.StatusNotFound)
	body := "Method not allowed!\n"
	fmt.Fprintf(w, "%s", body)
}
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	t := time.Now().Format("2006-01-02 15:04:05")
	Body := "time is: " + t + "\n"
	fmt.Fprintf(w, "%s", Body)
}
func AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	users := []User{}
	err = json.Unmarshal(d, &users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(users)
	if !IsUserAdmin(users[0]) {
		log.Println("Command issued by non-admin user", users[0].Username)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := InsertUser(users[1])
	if !result {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteHandler is for deleting an existing user + DELETE
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID not found in request")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var user = User{}
	err := user.FromJson(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !IsUserAdmin(user) {
		log.Println("User", user.Username, "is not admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id", err)
		return
	}

	t := FindUserID(intID)
	if t.Username != "" {
		log.Println("for delete:", t)
		deleteed := DeleteUser(intID)
		if deleteed {
			log.Println("User", intID, "deleted", id)
			w.WriteHeader(http.StatusOK)
			return
		} else {
			log.Println("User", intID, "not deleted", id)
			w.WriteHeader(http.StatusNotFound)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	if len(d) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty body")
		return
	}
	fmt.Println("Get All:", string(d))
	user := User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println("Get", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !IsUserAdmin(user) {
		log.Println("User", user.Username, "is not admin")
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	err = se(ListAllUsers(), w)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// GetIDHandler returns the ID of an existing user
func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	username, ok := mux.Vars(r)["username"]
	if !ok {
		log.Println("Username not found in request")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	if len(d) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty body")
		return
	}
	user := User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println("Get", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(user)
	if !IsUserAdmin(user) {
		log.Println("User", user.Username, "is not admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := FindUserUsername(username)
	if t.ID != 0 {
		w.WriteHeader(http.StatusFound)
		err = t.ToJson(w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println("GetIDHandler( terminated OK.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		log.Println("User " + username + " not found")
	}
}

// GetUserDataHandler + GET returns the full record of a user
func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID not found in request")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	t := FindUserID(intID)
	if t.ID != 0 {
		err = t.ToJson(w)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)

		}
		return
	}
	log.Println("User", intID, "not found")
	w.WriteHeader(http.StatusNotFound)
}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	if len(d) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty body")
		return
	}
	var users = []User{}
	err = json.Unmarshal(d, &users)
	if err != nil {
		log.Println("Update", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !IsUserAdmin(users[0]) {
		log.Println("Command issued by non-admin user: ", users[0].Username)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(users)
	t := FindUserUsername(users[1].Username)
	t.Username = users[1].Username
	t.Password = users[1].Password
	t.Admin = users[1].Admin

	if !UpdateUser(t) {
		log.Println("User", users[1].Username, "not updated")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("User", users[1].Username, "updated")
	w.WriteHeader(http.StatusOK)

}

// LoginHandler is for updating the LastLogin time of a user
// And changing the Active field to true

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	if len(d) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty body")
		return
	}
	var user = User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Input user", user)

	if !IsUserValid(user) {
		log.Println("User", user.Username, "is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := FindUserUsername(user.Username)
	log.Println("logging in: ", t)

	t.LastLogin = time.Now().Unix()
	t.Active = 1
	if UpdateUser(t) {
		log.Println("User updated:", t)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("Update failed:", t)
		w.WriteHeader(http.StatusBadRequest)
	}
}

// LogoutHandler is for logging out a user
// And changing the Active field to false
func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("LogoutHandler Serving:", r.URL.Path, "from", r.Host)

	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := FindUserUsername(user.Username)
	log.Println("Logging out:", t.Username)
	t.Active = 0
	if UpdateUser(t) {
		log.Println("User updated:", t)
		rw.WriteHeader(http.StatusOK)
	} else {
		log.Println("Update failed:", t)
		rw.WriteHeader(http.StatusBadRequest)
	}
}

// LoggedUsersHandler returns the list of all logged in users
func LoggedUsersHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoggedUsersHandler Serving:", r.URL.Path, "from", r.Host)
	var user = User{}
	err := user.FromJson(r.Body)

	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = se(ReturnLoggedUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
