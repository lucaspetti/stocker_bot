package quote

import (
	"errors"
	"testing"

	finance "github.com/piquette/finance-go"
)

var SimulatedErrorFromGetCrypto = errors.New("Error fetching crypto")
var mockCryptoPair = &finance.CryptoPair{
	Algorithm:         "Mock",
	StartDate:         0,
	MaxSupply:         1000,
	CirculatingSupply: 999,
}

func TestGetCrypto(t *testing.T) {
	appliedTemplate := `
Algorithm:      Mock
StartDate:      0

Regular Market Price: 100 USD

MaxSupply         1000
CirculatingSupply 999

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
			symbol:           "COIN",
			expectedResponse: appliedTemplate,
		},
		{
			title:         "Error from Get Crypto",
			expectedError: SimulatedErrorFromGetCrypto,
		},
		{
			title:         "Error from Get Quote",
			symbol:        "quote_err",
			expectedError: SimulatedErrorFromGetQuote,
		},
	}

	for _, test := range cases {
		mockGet := NewCryptoGet(
			mockGetCrypto,
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

func mockGetCrypto(symbol string) (*finance.CryptoPair, error) {
	if symbol == "" {
		return nil, SimulatedErrorFromGetCrypto
	}

	return mockCryptoPair, nil
}
