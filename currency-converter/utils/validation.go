package utils

// IsSupportedCurrency checks if a currency is supported.
func IsSupportedCurrency(currency string, supportedCurrencies []string) bool {
	for _, c := range supportedCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}
