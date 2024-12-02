package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"currency-converter/utils"

	"github.com/gorilla/mux"
)

type CurrencyResponse struct {
	Data map[string]float64 `json:"data"`
}

// SupportedCurrencies contains a list of currencies that can be converted.
var SupportedCurrencies = []string{"USD", "CAD", "EUR", "AUD", "CNY", "INR"}

// ConvertCurrency handles the HTTP request for currency conversion.
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	baseCurrency := mux.Vars(r)["base"]

	// Validate base currency
	if !utils.IsSupportedCurrency(baseCurrency, SupportedCurrencies) {
		http.Error(w, "Invalid base currency. Supported currencies are: "+strings.Join(SupportedCurrencies, ", "), http.StatusBadRequest)
		return
	}

	// Fetch the API key
	apiKey := os.Getenv("CURRENCY_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key is not set", http.StatusInternalServerError)
		return
	}

	// Build API request URL
	currencyList := strings.Join(SupportedCurrencies, ",")
	url := fmt.Sprintf("https://api.freecurrencyapi.com/v1/latest?apikey=%s&base_currency=%s&currencies=%s", apiKey, baseCurrency, currencyList)

	// Make the API call
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch data from the API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check API response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from API: "+resp.Status, resp.StatusCode)
		return
	}

	// Parse API response
	var currencyResponse CurrencyResponse
	err = json.NewDecoder(resp.Body).Decode(&currencyResponse)
	if err != nil {
		http.Error(w, "Failed to parse API response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Remove base currency from the results
	delete(currencyResponse.Data, baseCurrency)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currencyResponse.Data)
}
