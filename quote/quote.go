package quote

import "errors"

type DataGetter interface {
	GetData(symbol string) (response string, err error)
}

var ErrQuoteNotFound = errors.New("Quote Not Found")

// TODO: Remove from here, define only struct
// const quoteTemplate = `
// Name:      {{.LongName}}
// Symbol:    {{.Symbol}}

// EPS Trailing: {{.EpsTrailingTwelveMonths}}
// EPS Forward:  {{.EpsForward}}

// Trailing PE: {{.TrailingPE}}
// Forward PE:  {{.ForwardPE}}

// Price to Book:     {{.PriceToBook}}
// SharesOutstanding: {{.SharesOutstanding}}
// Market Cap:        {{.MarketCap}}

// Click /back to go back to main menu
// `

// type DataFetcher interface {
// 	GetData(symbol string) (quoteResp string, err error)
// }

// var fetcher QuoteFetcher

// type QuoteFetcher struct{}

// type QuoteGetter interface {
// 	Get(symbol string) (*finance.Quote, error)
// }

// func (f QuoteFetcher) GetQuote(symbol string) (quoteResp string, err error) {
// 	quote, err := equity.Get(symbol)
// 	if err != nil {
// 		return "", err
// 	}

// 	if quote == nil {
// 		return "", ErrQuoteNotFound
// 	}

// 	buf := &bytes.Buffer{}
// 	template, err := template.New("quote_text").Parse(quoteTemplate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = template.Execute(buf, quote)
// 	if err != nil {
// 		panic(err)
// 	}

// 	quoteResp = buf.String()
// 	return
// }

// Quote is the basic quote structure shared across asset classes.
// Contains most fields that are common across all possible assets.

// type Quote struct {
// 	// Quote classifying fields.
// 	Symbol      string      `json:"symbol"`
// 	MarketState MarketState `json:"marketState"`
// 	QuoteType   QuoteType   `json:"quoteType"`
// 	ShortName   string      `json:"shortName"`

// 	// Regular session quote data.
// 	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
// 	RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose"`
// 	RegularMarketPrice         float64 `json:"regularMarketPrice"`
// 	RegularMarketTime          int     `json:"regularMarketTime"`
// 	RegularMarketChange        float64 `json:"regularMarketChange"`
// 	RegularMarketOpen          float64 `json:"regularMarketOpen"`
// 	RegularMarketDayHigh       float64 `json:"regularMarketDayHigh"`
// 	RegularMarketDayLow        float64 `json:"regularMarketDayLow"`
// 	RegularMarketVolume        int     `json:"regularMarketVolume"`

// 	// Quote depth.
// 	Bid     float64 `json:"bid"`
// 	Ask     float64 `json:"ask"`
// 	BidSize int     `json:"bidSize"`
// 	AskSize int     `json:"askSize"`

// 	// Pre-market quote data.
// 	PreMarketPrice         float64 `json:"preMarketPrice"`
// 	PreMarketChange        float64 `json:"preMarketChange"`
// 	PreMarketChangePercent float64 `json:"preMarketChangePercent"`
// 	PreMarketTime          int     `json:"preMarketTime"`

// 	// Post-market quote data.
// 	PostMarketPrice         float64 `json:"postMarketPrice"`
// 	PostMarketChange        float64 `json:"postMarketChange"`
// 	PostMarketChangePercent float64 `json:"postMarketChangePercent"`
// 	PostMarketTime          int     `json:"postMarketTime"`

// 	// 52wk ranges.
// 	FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
// 	FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent"`
// 	FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
// 	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`
// 	FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
// 	FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`

// 	// Averages.
// 	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
// 	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
// 	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
// 	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
// 	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
// 	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`

// 	// Volume metrics.
// 	AverageDailyVolume3Month int `json:"averageDailyVolume3Month"`
// 	AverageDailyVolume10Day  int `json:"averageDailyVolume10Day"`

// 	// Quote meta-data.
// 	QuoteSource               string `json:"quoteSourceName"`
// 	CurrencyID                string `json:"currency"`
// 	IsTradeable               bool   `json:"tradeable"`
// 	QuoteDelay                int    `json:"exchangeDataDelayedBy"`
// 	FullExchangeName          string `json:"fullExchangeName"`
// 	SourceInterval            int    `json:"sourceInterval"`
// 	ExchangeTimezoneName      string `json:"exchangeTimezoneName"`
// 	ExchangeTimezoneShortName string `json:"exchangeTimezoneShortName"`
// 	GMTOffSetMilliseconds     int    `json:"gmtOffSetMilliseconds"`
// 	MarketID                  string `json:"market"`
// 	ExchangeID                string `json:"exchange"`
// }
