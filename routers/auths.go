package routers

import (
	"go-api/connectors"
	"go-api/handlers"
	"go-api/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoute(r *mux.Router) {
	authRepo := repositories.RepoAuth(connectors.DB)
	h := handlers.HandlerAuth(authRepo)
	// Routing
	r.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/login", h.SelectUser).Methods(http.MethodPost)

}
