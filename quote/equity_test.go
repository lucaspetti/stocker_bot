package quote

import (
	"errors"
	"testing"

	finance "github.com/piquette/finance-go"
)

var SimulatedErrorFromGetEquity = errors.New("Simulated Error from Get Equity")
var SimulatedErrorFromGetQuote = errors.New("Simulated Error from Get Quote")

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

func TestGet(t *testing.T) {
	appliedTemplate := `
Name:      Mock Ticker
Symbol:    TICKR

Regular Market Price: 100 USD

EPS Trailing: 0.2
EPS Forward:  0.5

Trailing PE: 1.5
Forward PE:  2

Price to Book:     3
Market Cap:        1.20M

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
	}

	for _, test := range cases {
		mockGet := NewEquityGet(
			mockGetEquity,
			mockGetQuote,
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
