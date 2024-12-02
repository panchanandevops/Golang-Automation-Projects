package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-app/models"
)

func GetGeolocation(city, apiKey string) (*models.GeoResponse, error) {
	geoURL := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(geoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch geolocation data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching geolocation data: %s", resp.Status)
	}

	var geoData []models.GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil || len(geoData) == 0 {
		return nil, fmt.Errorf("failed to parse geolocation data or city not found")
	}

	return &geoData[0], nil
}

func GetWeatherData(lat, lon float64, apiKey string) (*models.WeatherResponse, error) {
	weatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey)
	resp, err := http.Get(weatherURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather data: %s", resp.Status)
	}

	var weatherData models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return &weatherData, nil
}
