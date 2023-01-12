package main

import (
	"go-api/connectors"
	"go-api/migration"

	"github.com/joho/godotenv"
)

func main() {

	// Cors
	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	// var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	// var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	err := godotenv.Load()
	if err != nil {

	}

	connectors.DBInit()
	migration.RunMigration()

}
