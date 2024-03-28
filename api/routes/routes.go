package routes

import (
	"github.com/bhaveshparsad/fluxkart-identifier/api/controllers"
	"github.com/gorilla/mux"
)

var FluxKartRoutes = func(router *mux.Router) {
	router.HandleFunc("/identify", controllers.Identify).Methods("POST")
}
