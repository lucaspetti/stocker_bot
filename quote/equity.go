package quote

import (
	"bytes"
	"errors"
	"html/template"
	"stocker_bot/numbers"

	finance "github.com/piquette/finance-go"
)

// TODO: Format as table
const equityTemplate = `
Name:      {{.Data.LongName}}
Symbol:    {{.Quote.Symbol}}

Market Price:  {{.Quote.RegularMarketPrice}} {{.Quote.CurrencyID}}
Market Cap:    {{.MarketCap}}

EPS Trailing:       {{.Data.EpsTrailingTwelveMonths}}
EPS Forward:       {{.Data.EpsForward}}
EPS Growth 5Y:   {{printf "%.2f" .ValueData.EPSGrowth5Y}}%

Trailing PE:      {{printf "%.2f" .Data.TrailingPE}}
Forward PE:     {{printf "%.2f" .Data.ForwardPE}}

Value Data:

Price to Book:                             {{printf "%.2f" .Data.PriceToBook}}
Book Value Growth 5Y:             {{printf "%.2f" .ValueData.BookValueGrowth5Y}}%
ROI 5Y:                                        {{printf "%.2f" .ValueData.ROI5Y}}
Revenue Growth 5Y:                  {{printf "%.2f" .ValueData.RevenueGrowth5Y}}%
Revenue Per Share Growth 5Y:   {{printf "%.2f" .ValueData.RevenueShareGrowth5Y}}%
Free Operating Cash Flow 5Y:   {{printf "%.2f" .ValueData.FOCFCagr5Y}}

Debt:

Long Term Debt/Eq Annual:     {{printf "%.2f" .ValueData.LongTermDebtPerequityAnnual}}
Total Debt/Total Eq Annual: {{printf "%.2f" .ValueData.TotalDebtPerTotalEquityAnnual}}

Click /back to go back to main menu
`

var ErrEquityNotFound = errors.New("Quote Not Found")

type EquityGet struct {
	getEquityFunc func(symbol string) (*finance.Equity, error)
	getQuoteFunc  func(symbol string) (*finance.Quote, error)
	getValueFunc  func(symbol string) (*ValueData, error)
}

func (g EquityGet) GetEquity(symbol string) (*finance.Equity, error) {
	return g.getEquityFunc(symbol)
}

func (g EquityGet) GetQuote(symbol string) (*finance.Quote, error) {
	return g.getQuoteFunc(symbol)
}

func (g EquityGet) GetValue(symbol string) (*ValueData, error) {
	return g.getValueFunc(symbol)
}

func NewEquityGet(
	getEquity func(symbol string) (*finance.Equity, error),
	getQuote func(symbol string) (*finance.Quote, error),
	getValue func(symbol string) (*ValueData, error),
) *EquityGet {
	return &EquityGet{
		getEquityFunc: getEquity,
		getQuoteFunc:  getQuote,
		getValueFunc:  getValue,
	}
}

type EquityData struct {
	Quote     finance.Quote
	Data      finance.Equity
	MarketCap string
	ValueData ValueData
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

	valueData, err := g.GetValue(symbol)
	if err != nil {
		return "", err
	}

	data := EquityData{
		Data:      *equity,
		Quote:     *quote,
		ValueData: *valueData,
		MarketCap: numbers.FormatSuffix(equity.MarketCap),
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
