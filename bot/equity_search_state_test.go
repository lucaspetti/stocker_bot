package bot

import (
	"errors"
	"testing"
)

var SimulatedFetcherError = errors.New("Simulated error from Fetcher")

func TestQuoteSearchState(t *testing.T) {
	cases := []struct {
		Title                         string
		message                       string
		expectedResponse              string
		mockEnterInitialStateResponse string
		mockStockQuoteResponse        string
		errFetchingQuote              error
	}{
		{
			Title:                         "Sending /back command",
			message:                       "/back",
			expectedResponse:              "welcome",
			mockEnterInitialStateResponse: "welcome",
		},
		{
			Title:                  "When passing a ticker",
			message:                "TICKER",
			expectedResponse:       "PE: 30 \n PB: 50",
			mockStockQuoteResponse: "PE: 30 \n PB: 50",
		},
		{
			Title:            "Error fetching quote",
			message:          "",
			expectedResponse: tickerNotFound,
			errFetchingQuote: SimulatedFetcherError,
		},
	}

	for _, test := range cases {
		mockBot := &mockStockerBot{
			mockEnterInitialStateResponse: test.mockEnterInitialStateResponse,
		}

		mockDataGetter := &MockDataGetter{
			Response: test.mockStockQuoteResponse,
			Err:      test.errFetchingQuote,
		}

		quoteState := &EquitySearchState{
			stockerBot: mockBot,
			dataGetter: mockDataGetter,
		}

		got := quoteState.buildResponse(test.message)
		want := test.expectedResponse

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

type MockDataGetter struct {
	Response string
	Err      error
}

func (mock MockDataGetter) GetData(symbol string) (response string, err error) {
	if symbol == "" {
		return "", mock.Err
	}

	return mock.Response, nil
}
