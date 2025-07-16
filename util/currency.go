package util

var supportedCurrencies = map[string]bool{
	"INR": true,
	"USD": true,
	"EUR": true,
	"CAD": true,
}

func IsSupportedCurrency(currency string) bool {
	return supportedCurrencies[currency]
}
