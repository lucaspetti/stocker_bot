package bot

type QuoteSearchState struct {
	stockerBot StockBot
}

const tickerNotFound = "Sorry, ticker was not found"

func (s QuoteSearchState) buildResponse(message string) (response string) {
	if message == "/back" {
		return s.stockerBot.enterInitialState()
	}

	// TODO: validate that ticker has only alphanumeric before sending request
	// r'[$][A-Za-z][\S]*

	response, err := s.stockerBot.getStockQuote(message)

	if err != nil {
		return tickerNotFound
	}
	return
}
