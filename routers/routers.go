package routers

import "github.com/gorilla/mux"

func Router(r *mux.Router) {

	AuthRoute(r)

}
