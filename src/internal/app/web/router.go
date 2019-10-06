package web

import (
	"internal/app/controllers"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() *Router {
	router := new(Router)
	return router
}

func (router *Router) Start() {
	router.Router = mux.NewRouter().StrictSlash(true)

	clientController := controllers.NewClientController()
	router.Router.HandleFunc("/clients", clientController.GetClients).Methods("GET")
	router.Router.HandleFunc("/client", clientController.CreateClient).Methods("POST")
}
