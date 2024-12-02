package main

import (
	"fmt"
	"log"
	"net/http"

	"weather-app/handlers"
	"weather-app/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnv()

	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", handlers.GetWeather).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
