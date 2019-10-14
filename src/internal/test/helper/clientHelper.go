package testhelper

import (
	"encoding/json"
	"internal/app/model"
)

func BuildClientModel() model.Client {
	client := model.NewUser("Augusto", "augustomarinho@conteudoatual.com.br")
	return *client
}

func BuildMockClients() []model.Client {
	client := BuildClientModel()
	var clients = make([]model.Client, 1, 1)
	clients[0] = client

	return clients
}

func BuildJsonClient() []byte {
	clientModel := BuildClientModel()
	jsonClient, _ := json.Marshal(clientModel)

	return jsonClient
}
