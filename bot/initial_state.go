package bot

type InitialState struct {
	stockerBot StockBot
}

const welcomeMessage = `Welcome to the Stocker Bot
Here are the possible commands:

/quote to start searching for stock data

/crypto to start search for crypto data
`

func (s InitialState) buildResponse(message string) (response string) {
	switch message {
	case "/quote":
		return s.stockerBot.enterQuoteState()
	case "/crypto":
		return s.stockerBot.enterCryptoState()
	default:
		return welcomeMessage
	}

}
