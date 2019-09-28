package main

import (
	"fmt"
	"internal/app/web"
	"net/http"
)

func main() {
	fmt.Println("Starting Web Application")
	webRoute := web.NewRouter()
	webRoute.Start()

	http.ListenAndServe(":8080", webRoute.Router)
}
