package bot

import "stocker_bot/quote"

type EquitySearchState struct {
	stockerBot StockBot
	dataGetter quote.DataGetter
}

const tickerNotFound = "Sorry, ticker was not found"

func (s EquitySearchState) buildResponse(message string) (response string) {
	if message == "/back" {
		return s.stockerBot.enterInitialState()
	}

	// TODO: validate that ticker has only alphanumeric before sending request
	// r'[$][A-Za-z][\S]*

	response, err := s.dataGetter.GetData(message)

	if err != nil {
		return tickerNotFound
	}
	return
}
