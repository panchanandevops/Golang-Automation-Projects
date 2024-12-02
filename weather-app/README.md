
# Weather App

A simple weather app built with Go, which retrieves weather information for a specified city using the OpenWeatherMap API. The app fetches geolocation and weather data, then provides weather details such as description and temperature in Celsius.

## Project Structure

```
weather-app
├── go.mod
├── go.sum
├── handlers
│   └── weather.go
├── main.go
├── models
│   └── weather_models.go
├── services
│   └── weather_service.go
└── utils
    └── env_loader.go
```

### **Main Files:**
- **`main.go`**: Entry point of the application where the HTTP server is set up and routes are defined.
- **`models/weather_models.go`**: Contains the data structures used for the API response and geolocation.
- **`services/weather_service.go`**: Handles the logic for fetching geolocation and weather data using OpenWeatherMap API.
- **`handlers/weather.go`**: Contains the handler that processes the incoming requests and sends the appropriate weather data.
- **`utils/env_loader.go`**: Loads environment variables from the `.env` file using the `godotenv` package.

## Installation


1. **Install dependencies:**
   The project uses Go modules to manage dependencies. Run the following to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

2. **Set up the `.env` file:**
   Create a `.env` file in the root of the project directory and add your OpenWeatherMap API key:
   ```
   WEATHER_API_KEY=your_api_key_here
   ```

## How to Run

1. **Run the application:**
   Start the server by running:
   ```bash
   go run main.go
   ```
   The server will run on `http://localhost:8080`.

2. **API Usage:**
   The app exposes an endpoint for fetching the weather of a specific city:
   ```
   GET /weather/{city}
   ```

   Example:
   ```bash
   curl http://localhost:8080/weather/London
   ```

   The response will be a JSON object containing the weather description and temperature in Celsius:
   ```json
   {
     "weather": "clear sky",
     "temperature": 15.0
   }
   ```

## Detailed Breakdown of Components

### **1. `main.go`**
This is the entry point of the application. It sets up the router and defines the route to fetch weather data for a specified city. The environment variables are loaded using the `utils.LoadEnv()` function.

### **2. `models/weather_models.go`**
This file contains the data structures used to parse the JSON response from the OpenWeatherMap API:
- `GeoResponse`: Contains latitude and longitude information for a city.
- `WeatherResponse`: Contains weather description and temperature data for the city.

### **3. `services/weather_service.go`**
This file handles the interaction with the OpenWeatherMap API. It contains two key functions:
- `GetGeolocation`: Retrieves the latitude and longitude of the city using the OpenWeatherMap geolocation API.
- `GetWeatherData`: Fetches the weather data for the given latitude and longitude, returning the weather description and temperature.

### **4. `handlers/weather.go`**
This file contains the handler function `GetWeather`, which is called when the API endpoint is accessed. It retrieves the weather data for a city by first calling the `GetGeolocation` function to get the latitude and longitude and then the `GetWeatherData` function to fetch weather details. The result is returned as a JSON response with the weather description and temperature in Celsius.

### **5. `utils/env_loader.go`**
This file is responsible for loading environment variables from the `.env` file using the `godotenv` package. The API key for OpenWeatherMap is stored in the `.env` file, and this file ensures it's available to the application at runtime.

## Dependencies

- **Go**: The Go programming language is required to build and run the project.
- **gorilla/mux**: A powerful URL router and dispatcher for Go.
- **godotenv**: Loads environment variables from a `.env` file.

You can install the dependencies by running:
```bash
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```