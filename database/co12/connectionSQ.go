package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var dbname = "ali.db"

func insertData(db *sql.DB, dsc string) error {
	cT := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("INSERT INTO book VALUES(NULL,?,?);")
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(cT, dsc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func selectData(db *sql.DB, n int) error {
	rows, err := db.Query("SELECT * FROM book WHERE id > ? ;", n)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var desc string
		err = rows.Scan(&id, &name, &desc)
		if err != nil {
			fmt.Println(err)
			return err
		}
		data, err := time.Parse("2006-01-02 15:04:05", name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("id:%d name:%s desc:%s \n", id, data, desc)
	}
	return nil
}

func main() {

	os.Remove(dbname)

	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer db.Close()

	const create string = `
CREATE TABLE IF NOT EXISTS book (
id INTEGER NOT NULL PRIMARY KEY,
time TEXT NOT NULL,
desc TEXT);`

	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Insert 10 rows to the book table

	for i := 1; i < 11; i++ {
		dc := "Description: " + strconv.Itoa(i)
		err = insertData(db, dc)
		if err != nil {
			fmt.Println("Insert data:", err)
		}
		// select multiple rows
		err = selectData(db, 5)
		if err != nil {
			fmt.Println("Select data:", err)
		}

		//update Data
		cT := time.Now().Format("2006-01-02 15:04:05")
		db.Exec("UPDATE book SET time= ? WHERE id > ?", cT, 7)

		err = selectData(db, 8)
		if err != nil {
			fmt.Println("Select data:", err)
			return
		}
		//Deleting
		stmt, err := db.Prepare("DELETE FROM book WHERE id = ?")
		_, err = stmt.Exec(8)
		if err != nil {
			fmt.Println("Delete book:", err)
			return
		}
		err = selectData(db, 7)
		if err != nil {
			fmt.Println("Select data:", err)
			return
		}
		//counting rows in table
		query, err := db.Query("SELECT count(*) as count from book")
		if err != nil {
			fmt.Println("Query:", err)
			return
		}
		defer query.Close()

		count := -100
		for query.Next() {
			_ = query.Scan(&count)
		}
		fmt.Println("count(*):", count)
	}

}
