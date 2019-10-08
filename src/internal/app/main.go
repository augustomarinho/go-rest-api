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
	database := connectOnDataBase()
	defer database.CloseConnection()

	webRoute := web.NewRouter()
	webRoute.Start()

	http.ListenAndServe(":8080", webRoute.Router)
}

func connectOnDataBase() *infrastructure.Database {
	dataBase := infrastructure.ConnectDatabase()
	dataBase.InitializeDatabase(&entities.ClientEntity{})
	return dataBase
}
