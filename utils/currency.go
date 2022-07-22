package util

// import "github.com/go-playground/locales/currency"

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD , EUR , CAD:
		return true
	}
	return false
}