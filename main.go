package main

import (
	"go-api/connectors"
	"go-api/migration"
	"go-api/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Cors
	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	err := godotenv.Load()
	if err != nil {

	}

	connectors.DBInit()
	migration.RunMigration()

	r := mux.NewRouter()
	routers.Router(r.PathPrefix("/api/v1").Subrouter())
	log.Println("Server Running!")

	srverr := http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
	log.Println(srverr)

}
