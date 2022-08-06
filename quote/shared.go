package quote

import "errors"

type DataGetter interface {
	GetData(symbol string) (response string, err error)
}

var (
	ErrQuoteNotFound = errors.New("Quote Not Found")
	// ErrCompanyBasicFinancialsNotFound is returned when company is not found on Finnhub
	ErrCompanyBasicFinancialsNotFound = errors.New("Company Basic Financials Not Found")
	ErrMetricsNotFound                = errors.New("Metrics Not Found")
)
