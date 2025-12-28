package main

import (
	"fmt"
	"os"

	"github.com/alianjidaniir-design/sqlite06"
)

func main() {
	sqlite06.Filename = "Ali.db"

	db, err := sqlite06.openConnection
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connection string", db)
	os.Remove(sqlite06.Filename)

}
