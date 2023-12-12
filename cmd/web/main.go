package main

import (
	"fmt"
	"golangWebApp/pkg/handlers"
	"net/http"
)

const portNumber string = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
