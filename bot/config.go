package bot

type Config struct {
	telegramApiToken string
	authorizedUserID int
	finnhubApiToken  string
}

func NewConfig(
	authorizedUserID int,
	telegramApiToken,
	finnhubApiToken string,
) *Config {
	return &Config{
		telegramApiToken: telegramApiToken,
		authorizedUserID: authorizedUserID,
		finnhubApiToken:  finnhubApiToken,
	}
}
