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
	EpsTrailingTwelveMonths: 0.25,
	EpsForward:              0.57,
	TrailingPE:              1.502323,
	ForwardPE:               2.002323,
	PriceToBook:             3,
	MarketCap:               1200000,
}

var mockQuote = &finance.Quote{
	RegularMarketPrice: 100,
	CurrencyID:         "USD",
}

var mockValue = &ValueData{
	ROI5Y:                         10.002,
	RevenueGrowth5Y:               9.0098,
	EPSGrowth5Y:                   8.0012,
	PENormalizedAnnual:            7.0012,
	PEExclExtraTTM:                6.0012,
	BookValueGrowth5Y:             5.0523,
	RevenueShareGrowth5Y:          4.0452,
	LongTermDebtPerequityAnnual:   3.0,
	TotalDebtPerTotalEquityAnnual: 2.0,
	FOCFCagr5Y:                    1.0,
}

func TestGet(t *testing.T) {
	appliedTemplate := `
<b>Mock Ticker</b>
100 USD
<pre>
| Mkt Cap    | 1.20M
| EPS Trail. | 0.25
| EPS For    | 0.57
| EPS Gr5Y   | 8.00%
| Trail. PE  | 1.50
| Forward PE | 2.00
</pre>
<b>Value</b>
<pre>
| PB Ratio      | 3.00
| PB * PE       | 6.01
| BV Growth     | 5.05%
| ROI5Y         | 10.00
| Rev Gr5Y      | 9.01%
| RevShareGr 5Y | 4.05%
| FOCF agr 5Y   | 1.00
</pre>
<b>Debt</b>
<pre>
Long Term Debt/Eq Y   3.00
Total Debt/Total Eq Y 2.00
</pre>
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
