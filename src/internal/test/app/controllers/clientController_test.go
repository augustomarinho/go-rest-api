package controllers

import (
	"internal/app/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClientsHandlerWithSuccess(t *testing.T) {

	req, err := http.NewRequest("GET", "/clients", nil)

	if err != nil {
		t.Fatal(err)
	}

	clientController := controllers.NewClientController()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.GetClients)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"name":"Augusto","email":"augustomarinho@conteudoatual.com.br"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
