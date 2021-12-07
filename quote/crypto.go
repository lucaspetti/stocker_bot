package quote

import (
	"bytes"
	"errors"
	"html/template"
	"stocker_bot/numbers"

	finance "github.com/piquette/finance-go"
)

const cryptoTemplate = `
Algorithm:      {{.Crypto.Algorithm}}
StartDate:      {{.Crypto.StartDate}}

Regular Market Price: {{.Quote.RegularMarketPrice}} {{.Quote.CurrencyID}}

MaxSupply         {{.MaxSupply}}
CirculatingSupply {{.CirculatingSupply}}

Click /back to go back to main menu
`

var ErrCoinNotFound = errors.New("Coin Not Found")

type CryptoGet struct {
	getCryptoFunc func(symbol string) (*finance.CryptoPair, error)
	getQuoteFunc  func(symbol string) (*finance.Quote, error)
}

func (g CryptoGet) GetCrypto(symbol string) (*finance.CryptoPair, error) {
	return g.getCryptoFunc(symbol)
}

func (g CryptoGet) GetQuote(symbol string) (*finance.Quote, error) {
	return g.getQuoteFunc(symbol)
}

func NewCryptoGet(
	getCrypto func(symbol string) (*finance.CryptoPair, error),
	getQuote func(symbol string) (*finance.Quote, error),
) *CryptoGet {
	return &CryptoGet{
		getCryptoFunc: getCrypto,
		getQuoteFunc:  getQuote,
	}
}

// CryptoData holds the data for a Crypto
type CryptoData struct {
	Crypto            finance.CryptoPair
	Quote             finance.Quote
	MaxSupply         string
	CirculatingSupply string
}

func (g CryptoGet) GetData(symbol string) (cryptoResponse string, err error) {
	crypto, err := g.GetCrypto(symbol)
	if err != nil {
		return "", err
	}

	if crypto == nil {
		return "", ErrCoinNotFound
	}

	quote, err := g.GetQuote(symbol)
	if err != nil {
		return "", err
	}

	if quote == nil {
		return "", ErrQuoteNotFound
	}

	data := CryptoData{
		Crypto:            *crypto,
		Quote:             *quote,
		MaxSupply:         numbers.FormatSuffix(int64(crypto.MaxSupply)),
		CirculatingSupply: numbers.FormatSuffix(int64(crypto.CirculatingSupply)),
	}

	buf := &bytes.Buffer{}
	template, err := template.New("crypto_text").Parse(cryptoTemplate)
	if err != nil {
		return "", err
	}

	err = template.Execute(buf, data)
	if err != nil {
		return "", err
	}

	cryptoResponse = buf.String()
	return
}
