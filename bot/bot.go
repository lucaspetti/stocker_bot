package bot

import (
	"stocker_bot/stocks"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StockerBot struct {
	initialState state
	quoteSearch  state

	currentState state

	botApi       *tgbotapi.BotAPI
	stockFetcher stocks.QuoteFetcher
}

type state interface {
	buildResponse(message string) (response string)
}

type StockBot interface {
	enterQuoteState() (response string)
	enterInitialState() (response string)
	getStockQuote(symbol string) (response string, err error)
}

func NewStockerBot(botApi *tgbotapi.BotAPI, stockFetcher stocks.QuoteFetcher) *StockerBot {
	stockerBot := &StockerBot{
		botApi:       botApi,
		stockFetcher: stockFetcher,
	}

	initialState := &InitialState{
		stockerBot: stockerBot,
	}

	quoteSearch := &QuoteSearchState{
		stockerBot: stockerBot,
	}

	stockerBot.currentState = initialState

	stockerBot.initialState = initialState
	stockerBot.quoteSearch = quoteSearch
	return stockerBot
}

func (bot *StockerBot) enterQuoteState() (response string) {
	bot.currentState = bot.quoteSearch

	return "Hello, now you can search for your desired stocks"
}

func (bot *StockerBot) enterInitialState() (response string) {
	bot.currentState = bot.initialState

	return welcomeMessage
}

func (bot *StockerBot) getStockQuote(symbol string) (response string, err error) {
	return bot.stockFetcher.GetQuote(symbol)
}

func Start(apiToken string, authorizedUserID int) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		// TODO: Retry message or handle error better
		panic(err)
	}
	stockerBot := NewStockerBot(bot, stocks.Fetcher{})

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		// TODO: Retry message or handle error better
		panic(err)
	}

	for update := range updates {
		// Only look at messages for now and discard any other updates.
		if update.Message == nil {
			continue
		}

		message := update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.From.ID != authorizedUserID {
			msg.Text = "Unauthorized user"
		} else {
			msg.Text = stockerBot.currentState.buildResponse(message)
		}

		if _, err := bot.Send(msg); err != nil {
			// TODO: Retry message or handle error better
			panic(err)
		}
	}
}
