package main

import (
	"github.com/firebase007/go-rest-api-with-fiber/database"
	"github.com/firebase007/go-rest-api-with-fiber/router"
	"github.com/gofiber/fiber" // import the fiber package
	"github.com/gofiber/fiber/middleware"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//connect to database.go
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// call the New() method - used ti instantiate a new fiber App
	app := fiber.New()
	// Middlleware
	app.Use(middleware.Logger())
	router.SetupRoutes(app)
	app.Listen("3000")
}
