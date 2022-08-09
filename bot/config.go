package bot

type Config struct {
	telegramApiToken string
	authorizedUserID int64
	finnhubApiToken  string
}

func NewConfig(
	authorizedUserID int64,
	telegramApiToken,
	finnhubApiToken string,
) *Config {
	return &Config{
		telegramApiToken: telegramApiToken,
		authorizedUserID: authorizedUserID,
		finnhubApiToken:  finnhubApiToken,
	}
}
