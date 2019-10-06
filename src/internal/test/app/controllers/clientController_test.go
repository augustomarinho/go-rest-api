package controllers

import (
	"internal/app/controllers"
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

	clientController := controllers.NewClientController()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(clientController.GetClients)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	ja.Assertf(rr.Body.String(), `[{"name":"<<PRESENCE>>", "email": "<<PRESENCE>>"}]`)
}
