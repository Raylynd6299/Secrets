package main

import (
	"log"

	routes "github.com/Raylynd6299/secrets/src/Routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env variables on project
	err := godotenv.Load(".local.env")
	if err != nil {
		log.Fatal("Error on load env File")
	}

	// Run and Create Server
	routes.CreateServer()
}
