package bot

import (
	"stocker_bot/quote"
	"strings"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/piquette/finance-go/crypto"
	equity "github.com/piquette/finance-go/equity"
	finance_quote "github.com/piquette/finance-go/quote"
)

type StockerBot struct {
	initialState state
	equitySearch state
	cryptoSearch state

	currentState state

	botApi *tgbotapi.BotAPI
}

type state interface {
	buildResponse(message string) (response string)
}

type StockBot interface {
	enterQuoteState() (response string)
	enterInitialState() (response string)
	enterCryptoState() (response string)
}

func NewStockerBot(
	botApi *tgbotapi.BotAPI,
	equityGetter quote.DataGetter,
	cryptoGetter quote.DataGetter,
) *StockerBot {
	stockerBot := &StockerBot{
		botApi: botApi,
	}

	initialState := &InitialState{
		stockerBot: stockerBot,
	}

	equitySearchState := &EquitySearchState{
		stockerBot: stockerBot,
		dataGetter: equityGetter,
	}

	cryptoSearchState := &CryptoSearchState{
		stockerBot: stockerBot,
		dataGetter: cryptoGetter,
	}

	stockerBot.currentState = initialState

	stockerBot.initialState = initialState
	stockerBot.equitySearch = equitySearchState
	stockerBot.cryptoSearch = cryptoSearchState
	return stockerBot
}

func (bot *StockerBot) enterQuoteState() (response string) {
	bot.currentState = bot.equitySearch

	return `Hello, now you can search for your desired stocks

Click /back to return to the main menu`
}

func (bot *StockerBot) enterCryptoState() (response string) {
	bot.currentState = bot.cryptoSearch

	return `Hello, now you can search for your desired coins

Click /back to return to the main menu`
}

func (bot *StockerBot) enterInitialState() (response string) {
	bot.currentState = bot.initialState

	return welcomeMessage
}

func Start(config Config) {
	bot, err := tgbotapi.NewBotAPI(config.telegramApiToken)
	if err != nil {
		// TODO: Retry message or handle error better
		panic(err)
	}

	finnhubConfig := finnhub.NewConfiguration()
	finnhubConfig.AddDefaultHeader("X-Finnhub-Token", config.finnhubApiToken)
	finnhubClient := finnhub.NewAPIClient(finnhubConfig).DefaultApi

	getValueFunc := quote.GetValueFunc(finnhubClient)
	equityGetter := quote.NewEquityGet(equity.Get, finance_quote.Get, getValueFunc)
	cryptoGetter := quote.NewCryptoGet(crypto.Get, finance_quote.Get)
	stockerBot := NewStockerBot(bot, equityGetter, cryptoGetter)

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
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

		if update.Message.From.ID != config.authorizedUserID {
			msg.Text = "Unauthorized user"
		} else {
			msgText := stockerBot.currentState.buildResponse(message)
			if strings.Contains(msgText, "<pre>") {
				msg.ParseMode = tgbotapi.ModeHTML
			}

			msg.Text = msgText
		}

		if _, err := bot.Send(msg); err != nil {
			// TODO: Retry message or handle error better
			panic(err)
		}
	}
}
