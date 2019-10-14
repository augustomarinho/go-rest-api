package controllers

import (
	"bytes"
	"encoding/json"
	"internal/app/controllers"
	"internal/app/model"
	"internal/app/repository/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

func TestGetClientsHandlerWithSuccess(t *testing.T) {
	ja := jsonassert.New(t)

	req, err := http.NewRequest("GET", "/clients", nil)

	if err != nil {
		t.Fatal(err)
	}

	var clients = buildMockClients()
	clientRepositoryMock := new(mocks.ClientRepository)
	clientRepositoryMock.On("FindAll").Return(clients, nil)

	clientController := controllers.NewClientController(clientRepositoryMock)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.GetClients)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	ja.Assertf(rr.Body.String(), `[{"name":"<<PRESENCE>>", "email": "<<PRESENCE>>"}]`)
}

func TestCreateClientHandlerWithSuccess(t *testing.T) {

	expectedClient := buildClientModel()
	payload := buildJsonClient()

	clientRepositoryMock := new(mocks.ClientRepository)
	clientRepositoryMock.On("Save", &expectedClient).Return(nil)

	clientController := controllers.NewClientController(clientRepositoryMock)

	req, err := http.NewRequest("POST", "/client", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.CreateClient)
	handler.ServeHTTP(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusCreated, "they should be equal")
}

func buildClientModel() model.Client {
	client := model.NewUser("Augusto", "augustomarinho@conteudoatual.com.br")
	return *client
}

func buildMockClients() []model.Client {
	client := buildClientModel()
	var clients = make([]model.Client, 1, 1)
	clients[0] = client

	return clients
}

func buildJsonClient() []byte {
	clientModel := buildClientModel()
	jsonClient, _ := json.Marshal(clientModel)

	return jsonClient
}
