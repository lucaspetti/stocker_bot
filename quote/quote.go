package quote

import "errors"

type DataGetter interface {
	GetData(symbol string) (response string, err error)
}

var ErrQuoteNotFound = errors.New("Quote Not Found")
