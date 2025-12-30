package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Filename = "REST.db"

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", Filename)
	if err != nil {
		return nil
	}
	return db
}

func (p *User) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *User) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func DeleteUser(ID int) bool {
	db := OpenConnection()
	if db == nil {
		log.Fatal("Error opening database")
		return false
	}
	defer db.Close()
	t := FindUserID(ID)
	if t.ID == 0 {
		log.Println("User", ID, " does not exist")
		return false
	}
	stm, err := db.Prepare("DELETE FROM users WHERE UserID = $1")
	if err != nil {
		log.Println("Deleteuser:", err)
		return false
	}
	_, err = stm.Exec(ID)
	if err != nil {
		log.Println("Deleteuser:", err)
		return false
	}
	return true
}

func InsertUser(u User) bool {
	db := OpenConnection()
	if db == nil {
		log.Fatal("Error opening database")
		return false
	}
	defer db.Close()
	if IsUserValid(u) {
		log.Println("User", u.username, " already exists!")
		return false
	}
	stm, err := db.Prepare("INSERT INTO users(username, password , lastlogin,admin,active) values ($1, $2 ,$3,$4,$5)")
	if err != nil {
		log.Println("InsertUser:", err)
		return false
	}
	stm.Exec(u.Username, u.Password, u.Lastlogin, u.Admin, u.Active)
	return true
}

func listAllUsers() []Users {
	db := OpenConnection()
	if db == nil {
		log.Fatal("Error opening database")
		return nil
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		log.Println("listAllUsers:", err)
		return []User{}
	}
	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return []User{}
		}
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}
	log.Println("All:", all)
	return all
}

func ListLogged() []User {
	db := OpenConnection()
	if db == nil {
		log.Fatal("Error opening database")
		return []USer{}
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE active \n")
	if err != nil {
		log.Println(err)
		return []Users{}
	}
	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5 int
	var c6 int
	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {

			log.Println(err)
		}
		temp := Users{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)

	}
	log.Println("All:", all)
	return all

}

func IsUserValid(u User) bool {
	db := OpenConnection()
	if db == nil {
		log.Fatal("Error opening database")
		return false
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE username = $1 \n", u.Username)
	if err != nil {
		fmt.Println(err)
	}
	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5 int
	var c6 int
	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = Users{c1, c2, c3, c4, c5, c6}

	}
	if u.Username == temp.Username && u.Password == temp.Password {
		return true
	}
	return false
}
