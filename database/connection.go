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
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("INSERT INTO book VALUES(NULL, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t, dsc)
	if err != nil {
		return err
	}
	return nil
}
func selectData(db *sql.DB, n int) error {
	rows, err := db.Query("SELECT * FROM book WHERE id > ? ", n)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var dt string
		var description string
		err = rows.Scan(&id, &dt, &description)
		if err != nil {
			return err
		}

		data, err := time.Parse("2006-01-02 15:04:05", dt)

		if err != nil {
			return err
		}
		fmt.Println(id, data, description)
	}
	return nil

}

func main() {
	os.Remove(dbname)
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	const create string = `
CREATE TABLE IF NOT EXISTS book (
 id INTEGER NOT NULL PRIMARY KEY ,
 time text NOT NULL,
 description TEXT);`

	_, err = db.Exec(create)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < 11; i = i + 1 {
		dsc := "Description " + strconv.Itoa(i)
		err = insertData(db, dsc)
		if err != nil {
			fmt.Println(err)
		}

	}

	err = selectData(db, 5)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(1 * time.Second)

	ti := time.Now().Format("2006-01-02 15:04:05")
	db.Exec("UPDATE book SET  time = ? WHERE id > ?", ti, 7)

	err = selectData(db, 8)
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := db.Prepare("DELETE from book WHERE id = ?")
	_, err = stmt.Exec(8)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = selectData(db, 7)
	if err != nil {
		fmt.Println(err)
		return
	}
	query, err := db.Query("SELECT count(*) as count from book")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer query.Close()
	count := -100
	for query.Next() {
		_ = query.Scan(&count)
	}
	fmt.Println(count)
}
