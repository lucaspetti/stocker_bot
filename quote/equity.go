package quote

import (
	"bytes"
	"errors"
	"html/template"

	finance "github.com/piquette/finance-go"
)

const equityTemplate = `
Name:      {{.Data.LongName}}
Symbol:    {{.Quote.Symbol}}

Regular Market Price: {{.Quote.RegularMarketPrice}} {{.Quote.CurrencyID}}

EPS Trailing: {{.Data.EpsTrailingTwelveMonths}}
EPS Forward:  {{.Data.EpsForward}}

Trailing PE: {{.Data.TrailingPE}}
Forward PE:  {{.Data.ForwardPE}}

Price to Book:     {{.Data.PriceToBook}}
SharesOutstanding: {{.Data.SharesOutstanding}}
Market Cap:        {{.Data.MarketCap}}

Click /back to go back to main menu
`

var ErrEquityNotFound = errors.New("Quote Not Found")

type EquityGet struct {
	getEquityFunc func(symbol string) (*finance.Equity, error)
	getQuoteFunc  func(symbol string) (*finance.Quote, error)
}

func (g EquityGet) GetEquity(symbol string) (*finance.Equity, error) {
	return g.getEquityFunc(symbol)
}

func (g EquityGet) GetQuote(symbol string) (*finance.Quote, error) {
	return g.getQuoteFunc(symbol)
}

func NewEquityGet(
	getEquity func(symbol string) (*finance.Equity, error),
	getQuote func(symbol string) (*finance.Quote, error),
) *EquityGet {
	return &EquityGet{
		getEquityFunc: getEquity,
		getQuoteFunc:  getQuote,
	}
}

type EquityData struct {
	Quote finance.Quote
	Data  finance.Equity
}

func (g EquityGet) GetData(symbol string) (equityResponse string, err error) {
	equity, err := g.GetEquity(symbol)
	if err != nil {
		return "", err
	}

	if equity == nil {
		return "", ErrEquityNotFound
	}

	quote, err := g.GetQuote(symbol)
	if err != nil {
		return "", err
	}

	if quote == nil {
		return "", ErrQuoteNotFound
	}

	data := EquityData{
		Data:  *equity,
		Quote: *quote,
	}

	buf := &bytes.Buffer{}
	template, err := template.New("equity_text").Parse(equityTemplate)
	if err != nil {
		return "", err
	}

	err = template.Execute(buf, data)
	if err != nil {
		return "", err
	}

	equityResponse = buf.String()
	return
}
