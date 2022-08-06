package quote

import (
	"errors"
	"testing"

	finance "github.com/piquette/finance-go"
)

var SimulatedErrorFromGetEquity = errors.New("Simulated Error from Get Equity")
var SimulatedErrorFromGetQuote = errors.New("Simulated Error from Get Quote")
var SimulatedErrorFromGetValue = errors.New("Simulated Error from Get Value")

var mockEquity = &finance.Equity{
	LongName:                "Mock Ticker",
	EpsTrailingTwelveMonths: 0.2,
	EpsForward:              0.5,
	TrailingPE:              1.5,
	ForwardPE:               2.0,
	PriceToBook:             3,
	MarketCap:               1200000,
}

var mockQuote = &finance.Quote{
	Symbol:             "TICKR",
	RegularMarketPrice: 100,
	CurrencyID:         "USD",
}

var mockValue = &ValueData{
	ROI5Y:                         10.0,
	RevenueGrowth5Y:               9.0,
	EPSGrowth5Y:                   8.0,
	PENormalizedAnnual:            7.0,
	PEExclExtraTTM:                6.0,
	BookValueGrowth5Y:             5.0,
	RevenueShareGrowth5Y:          4.0,
	LongTermDebtPerequityAnnual:   3.0,
	TotalDebtPerTotalEquityAnnual: 2.0,
	FOCFCagr5Y:                    1.0,
}

func TestGet(t *testing.T) {
	appliedTemplate := `
Name:      Mock Ticker
Symbol:    TICKR

Market Price:  100 USD
Market Cap:    1.20M

EPS Trailing:    0.2
EPS Forward:     0.5
EPS Growth 5Y:   8

Trailing PE:     1.5
Forward PE:      2

Value Data:

Price to Book:                 3
Book Value Growth 5Y:          5
ROI 5Y:                        10
Revenue Growth 5Y:             9
Revenue Per Share Growth 5Y:   4
Free Operating Cash Flow 5Y:   1

Debt:

Long Term Debt Per Equity Annual:   3
Total Debt Per Total Equity Annual: 2

Click /back to go back to main menu
`

	cases := []struct {
		title            string
		symbol           string
		expectedResponse string
		expectedError    error
	}{
		{
			title:            "Successful response",
			symbol:           "TICKR",
			expectedResponse: appliedTemplate,
		},
		{
			title:         "Error from Get Equity",
			symbol:        "",
			expectedError: SimulatedErrorFromGetEquity,
		},
		{
			title:         "Error from Get Quote",
			symbol:        "quote_err",
			expectedError: SimulatedErrorFromGetQuote,
		},
		{
			title:         "Error from Get Value",
			symbol:        "value_err",
			expectedError: SimulatedErrorFromGetValue,
		},
	}

	for _, test := range cases {
		mockGet := NewEquityGet(
			mockGetEquity,
			mockGetQuote,
			mockGetValue,
		)

		got, err := mockGet.GetData(test.symbol)
		want := test.expectedResponse
		if err != nil && err != test.expectedError {
			t.Errorf("got unexpected error, expected %v, got %v", test.expectedError, err)
		}

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

func mockGetEquity(symbol string) (*finance.Equity, error) {
	if symbol == "" {
		return nil, SimulatedErrorFromGetEquity
	}

	return mockEquity, nil
}

func mockGetQuote(symbol string) (*finance.Quote, error) {
	if symbol == "quote_err" {
		return nil, SimulatedErrorFromGetQuote
	}

	return mockQuote, nil
}

func mockGetValue(symbol string) (*ValueData, error) {
	if symbol == "value_err" {
		return nil, SimulatedErrorFromGetValue
	}

	return mockValue, nil
}
