package main

import (
	"fmt"
	"log"
	"net/http"

	"currency-converter/config"
	"currency-converter/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize the router
	r := mux.NewRouter()
	r.HandleFunc("/convert/{base}", handlers.ConvertCurrency).Methods("GET")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
