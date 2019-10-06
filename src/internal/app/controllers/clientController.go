package controllers

import (
	"encoding/json"
	"internal/app/model"
	"internal/app/repository"
	"internal/app/rest"
	"net/http"
)

type ClientController struct {
	repository repository.ClientRepository
}

func NewClientController() *ClientController {

	controller := new(ClientController)
	controller.repository = new(repository.ClientRepositoryImpl)

	return controller
}

func (controller *ClientController) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := controller.repository.FindAll()

	if err != nil {
		rest.RespondWithError(w, http.StatusInternalServerError, "Problems to find all clients")
		return
	}
	rest.RespondWithJSON(w, http.StatusOK, clients)
}

func (controller *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client *model.Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		rest.RespondWithError(w, http.StatusCreated, "Wrong Post Message Payload")
	}

	controller.repository.Save(client)

	rest.RespondWithHttpCode(w, http.StatusCreated)
}
