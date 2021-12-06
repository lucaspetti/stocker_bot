package bot

import "stocker_bot/quote"

type CryptoSearchState struct {
	stockerBot StockBot
	dataGetter quote.DataGetter
}

const coinNotFound = "Sorry, coin was not found"

func (s CryptoSearchState) buildResponse(message string) (response string) {
	if message == "/back" {
		return s.stockerBot.enterInitialState()
	}

	// TODO: Render buttons or append "-eur" to the given message when searching
	response, err := s.dataGetter.GetData(message)

	if err != nil {
		return coinNotFound
	}
	return
}
