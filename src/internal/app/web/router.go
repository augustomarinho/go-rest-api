package web

import (
	"internal/app/controllers"
	"internal/app/repository"

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

	clientRepository := new(repository.ClientRepositoryImpl)
	clientController := controllers.NewClientController(clientRepository)
	router.Router.HandleFunc("/clients", clientController.GetClients).Methods("GET")
	router.Router.HandleFunc("/client", clientController.CreateClient).Methods("POST")
	router.Router.
		Path("/client").
		Methods("GET").
		Queries("email", "{email}").
		HandlerFunc(clientController.GetClientByEmail).
		Name("clientController.GetClientByEmail")
}
