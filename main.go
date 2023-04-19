package main

import (
	"challenge_10/database"
	"challenge_10/router"
	"os"
)

func main() {

	osParameter := os.Args[1]

	switch osParameter {
	case "http":
		PORT := ":8080"

		// connect to database
		db := database.StartDB()

		// call router
		router.StartServer(db).Run(PORT)

	case "migrate":
		// migrate database
		println("Migrating Database")
		database.MigrationDB()
		println("Migration Database Success!")

	default:
		println("Command not Found!")
	}

}
