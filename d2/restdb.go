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
