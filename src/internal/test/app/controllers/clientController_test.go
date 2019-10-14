package controllers

import (
	"internal/app/controllers"
	"internal/app/model"
	"internal/app/repository/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
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

func buildMockClients() []model.Client {
	client := model.NewUser("Augusto", "augustomarinho@conteudoatual.com.br")
	var clients = make([]model.Client, 1, 1)
	clients[0] = *client

	return clients
}
