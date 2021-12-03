package bot

import (
	"stocker_bot/stocks"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(apiToken string, authorizedUserID int) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	parser := NewParser(stocks.Fetcher{})

	if err != nil {
		panic(err)
	}

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
		panic(err)
	}

	for update := range updates {
		// Only look at messages for now and discard any other updates.
		if update.Message == nil {
			continue
		}

		message := update.Message.Text
		var msgText string

		if update.Message.From.ID != authorizedUserID {
			msgText = "Unauthorized user"
		} else {
			msgText = parser.Parse(message)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			// TODO: Retry message or handle error better
			panic(err)
		}
	}
}
