package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test1.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SQLITE version", version)
	os.Remove("test1.db")

}
