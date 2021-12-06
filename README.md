# Stocker Bot

A telegram bot to fetch stock market value indicators

## Running with Docker

First, make sure you have the .env file with the expected values set:

`TELEGRAM_APITOKEN`: The Api Token for your bot

`AUTHORIZED_USER_ID`: If you want to limit what user can message the bot, set this to the telegram USER_ID of your choice

To build the image and start the bot:

```docker-compose up```

### Tests

To run the tests:

```go test ./...```
