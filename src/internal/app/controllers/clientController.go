package controllers

import (
	"internal/app/model"
	"internal/app/rest"
	"net/http"
)

type ClientController struct {
}

func NewClientController() *ClientController {

	controller := new(ClientController)
	return controller
}

func (controller *ClientController) GetClients(w http.ResponseWriter, r *http.Request) {
	client := model.NewUser("Augusto", "augustomarinho@conteudoatual.com.br")
	var clients [1]*model.Client
	clients[0] = client

	rest.RespondWithJSON(w, http.StatusOK, clients)
}
