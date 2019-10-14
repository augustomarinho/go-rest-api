package controllers

import (
	"bytes"
	"internal/app/controllers"
	"internal/app/repository/mocks"
	testhelper "internal/test/helper"
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

	var clients = testhelper.BuildMockClients()
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

	expectedClient := testhelper.BuildClientModel()
	payload := testhelper.BuildJsonClient()

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

func TestGetClientByEmailHandlerWithSuccess(t *testing.T) {
	ja := jsonassert.New(t)
	expectedClient := testhelper.BuildClientModel()

	clientRepositoryMock := new(mocks.ClientRepository)
	clientRepositoryMock.On("FindByEmail", expectedClient.Email).Return(&expectedClient, nil)

	clientController := controllers.NewClientController(clientRepositoryMock)

	req, err := http.NewRequest("GET", "/client", nil)
	reqQuery := req.URL.Query()
	reqQuery.Add("email", expectedClient.Email)
	req.URL.RawQuery = reqQuery.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.GetClientByEmail)
	handler.ServeHTTP(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusOK, "they should be equal")
	ja.Assertf(rr.Body.String(), `{"name":"<<PRESENCE>>", "email": "<<PRESENCE>>"}`)
}

func TestGetClientByEmailHandler_ClientNotFound(t *testing.T) {
	expectedClient := testhelper.BuildClientModel()

	clientRepositoryMock := new(mocks.ClientRepository)
	clientRepositoryMock.On("FindByEmail", expectedClient.Email).Return(nil, nil)

	clientController := controllers.NewClientController(clientRepositoryMock)

	req, err := http.NewRequest("GET", "/client", nil)
	reqQuery := req.URL.Query()
	reqQuery.Add("email", expectedClient.Email)
	req.URL.RawQuery = reqQuery.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.GetClientByEmail)
	handler.ServeHTTP(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusNotFound, "they should be equal")
}
