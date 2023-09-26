package main

import (
	"m-banking-api/app"
	"m-banking-api/database"
	"m-banking-api/router"
)

func main() {
	// Initialize the app
	app.Setup()

	// Initialize the database connection
	database.InitDB()

	// Map URL routes
	router.SetupRoutes()

	// Start the server
	app.Router.Run(":5500")
}
