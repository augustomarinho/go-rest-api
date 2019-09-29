package main

import (
	"fmt"
	"internal/app/infrastructure"
	"internal/app/model"
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
	dataBase.Connect()
	dataBase.InitializeDatabase(&model.Client{})
	defer dataBase.CloseConnection()
}
