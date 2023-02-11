package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/luizmoitinho/bookstore_oauth_api/src/app"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	app.StartApplication()
}
