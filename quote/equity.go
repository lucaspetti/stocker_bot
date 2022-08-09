package quote

import (
	"bytes"
	"errors"
	"html/template"
	"stocker_bot/numbers"

	finance "github.com/piquette/finance-go"
)

const equityTemplate = `
<b>{{.Data.LongName}}</b>
{{.Quote.RegularMarketPrice}} {{.Quote.CurrencyID}}
<pre>
| Mkt Cap    | {{.MarketCap}}
| EPS Trail. | {{.Data.EpsTrailingTwelveMonths}}
| EPS For    | {{.Data.EpsForward}}
| EPS Gr5Y   | {{printf "%.2f" .ValueData.EPSGrowth5Y}}%
| Trail. PE  | {{printf "%.2f" .Data.TrailingPE}}
| Forward PE | {{printf "%.2f" .Data.ForwardPE}}
</pre>
<b>Value</b>
<pre>
| PB Ratio      | {{printf "%.2f" .Data.PriceToBook}}
| BV Growth     | {{printf "%.2f" .ValueData.BookValueGrowth5Y}}%
| ROI5Y         | {{printf "%.2f" .ValueData.ROI5Y}}
| Rev Gr5Y      | {{printf "%.2f" .ValueData.RevenueGrowth5Y}}%
| RevShareGr 5Y | {{printf "%.2f" .ValueData.RevenueShareGrowth5Y}}%
| FOCF agr 5Y   | {{printf "%.2f" .ValueData.FOCFCagr5Y}}
</pre>
<b>Debt</b>
<pre>
Long Term Debt/Eq Y   {{printf "%.2f" .ValueData.LongTermDebtPerequityAnnual}}
Total Debt/Total Eq Y {{printf "%.2f" .ValueData.TotalDebtPerTotalEquityAnnual}}
</pre>
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
