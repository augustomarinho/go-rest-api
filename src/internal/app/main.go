package main

import (
	"fmt"
	"internal/app/entities"
	"internal/app/infrastructure"
	"internal/app/web"
	"net/http"
)

func main() {
	fmt.Println("Starting Web Application")
	connectOnDataBase()

	webRoute := web.NewRouter()
	webRoute.Start()

	http.ListenAndServe(":8080", webRoute.Router)
}

func connectOnDataBase() {
	dataBase := infrastructure.NewDatabase()
	dataBase.InitializeDatabase(&entities.ClientEntity{})
	defer dataBase.CloseConnection()
}
