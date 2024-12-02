package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"weather-app/services"

	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		http.Error(w, "API key is not set", http.StatusInternalServerError)
		return
	}

	geoData, err := services.GetGeolocation(city, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherData, err := services.GetWeatherData(geoData.Lat, geoData.Lon, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temperatureCelsius := weatherData.Main.Temp - 273.15
	response := map[string]interface{}{
		"weather":     weatherData.Weather[0].Description,
		"temperature": temperatureCelsius,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
