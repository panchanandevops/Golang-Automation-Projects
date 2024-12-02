
# Currency Converter

A simple currency converter API built in Go. This application allows users to convert a base currency into other supported currencies, using the **Free Currency API**. It supports a predefined set of currencies and retrieves up-to-date conversion rates.

## Project Structure

```
currency-converter
├── config
│   └── env.go
├── go.mod
├── go.sum
├── handlers
│   └── convert.go
├── main.go
└── utils
    └── validation.go
```

### **Main Files:**
- **`main.go`**: Entry point of the application where the HTTP server is set up and routes are defined.
- **`config/env.go`**: Loads environment variables, including the API key used for the Free Currency API.
- **`handlers/convert.go`**: Contains the logic to handle currency conversion requests and interact with the Free Currency API.
- **`utils/validation.go`**: Provides utility functions like currency validation to ensure that the base currency is supported.

## Installation

1. **Install dependencies:**
   The project uses Go modules to manage dependencies. Run the following to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

2. **Set up the `.env` file:**
   Create a `.env` file in the root of the project directory and add your Free Currency API key:
   ```
   CURRENCY_API_KEY=your_api_key_here
   ```

## How to Run

1. **Run the application:**
   Start the server by running:
   ```bash
   go run main.go
   ```
   The server will run on `http://localhost:8080`.

2. **API Usage:**
   The app exposes an endpoint for converting currencies:
   ```
   GET /convert/{base}
   ```

   Example:
   ```bash
   curl http://localhost:8080/convert/USD
   ```

   The response will be a JSON object containing the conversion rates for supported currencies:
   ```json
   {
     "CAD": 1.34,
     "EUR": 0.85,
     "AUD": 1.45,
     "CNY": 6.42,
     "INR": 74.12
   }
   ```

## Detailed Breakdown of Components

### **1. `main.go`**
This file is the entry point of the application. It sets up the HTTP server and defines the route for currency conversion requests (`/convert/{base}`). The environment variables are loaded using the `config.LoadEnv()` function.

### **2. `config/env.go`**
This file loads the environment variables, specifically the API key (`CURRENCY_API_KEY`) required for accessing the Free Currency API. It uses the `godotenv` package to read variables from a `.env` file.

### **3. `handlers/convert.go`**
This file handles the logic for converting currencies. The `ConvertCurrency` function:
- Retrieves the base currency from the URL.
- Validates the base currency using the `utils.IsSupportedCurrency` function.
- Fetches the conversion rates from the Free Currency API.
- Returns a JSON response containing the conversion rates for supported currencies.

### **4. `utils/validation.go`**
This file contains a helper function `IsSupportedCurrency`, which checks if the requested base currency is supported by comparing it against a list of predefined currencies.

## Dependencies

- **Go**: The Go programming language is required to build and run the project.
- **gorilla/mux**: A powerful URL router and dispatcher for Go.
- **godotenv**: Loads environment variables from a `.env` file.

You can install the dependencies by running:
```bash
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```