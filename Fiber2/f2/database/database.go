package database

import (
	"database/sql"
	"fmt"
	"github.com/firebase007/go-rest-api-with-fiber/config"
	"strconv"
)

var db *sql.DB

func conncect() error {
	var err error
	p := config.Config("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Error parsing str ti int")
	}
	DB, err := sql.Open("SQLite", fmt.Sprintf("port=%d dbname=%s sslmode=disable", port,  config.Config("DB_NAME")))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
}
