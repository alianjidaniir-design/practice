package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "ali.db")
	if err != nil {
		fmt.Println("Error connection: ", err)
		return
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		fmt.Println("version: ", err)
		return
	}
	fmt.Println("SQLite3", version)
	err = os.Remove("ali.db")
	if err != nil {
		return
	}
}
