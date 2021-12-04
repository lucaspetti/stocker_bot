package stocks

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/piquette/finance-go/equity"
)

const quoteTemplate = `
Name:      {{.LongName}}
Symbol:    {{.Symbol}}

EPS Trailing: {{.EpsTrailingTwelveMonths}}
EPS Forward:  {{.EpsForward}}

Trailing PE: {{.TrailingPE}}
Forward PE:  {{.ForwardPE}}

Price to Book:     {{.PriceToBook}}
SharesOutstanding: {{.SharesOutstanding}}
Market Cap:        {{.MarketCap}}

Click /back to go back to main menu
`

var ErrQuoteNotFound = errors.New("Quote Not Found")

type QuoteFetcher interface {
	GetQuote(symbol string) (quoteResp string, err error)
}

var fetcher Fetcher

type Fetcher struct{}

func (f Fetcher) GetQuote(symbol string) (quoteResp string, err error) {
	quote, err := equity.Get(symbol)
	if err != nil {
		return "", err
	}

	if quote == nil {
		return "", ErrQuoteNotFound
	}

	buf := &bytes.Buffer{}
	template, err := template.New("quote_text").Parse(quoteTemplate)
	if err != nil {
		panic(err)
	}

	err = template.Execute(buf, quote)
	if err != nil {
		panic(err)
	}

	quoteResp = buf.String()

	return
}
