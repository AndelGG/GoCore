package config

type Config struct {
	ChatBotApiKey  string
	TelegramApiKey string
	ResponseTokens int `yaml:"responseTokens" env-default:"40"`
	WebPort        int
}

// TODO: Make reading config

func MustLoad() Config {
	//if _, err := os.Stat("./config/config.yaml"); os.IsNotExist(err) {
	//	panic("config does not exist")
	//}

	//chatBotApiKey := os.Getenv("DEEPSEEK_API_KEY")
	//tgBotApiKey := os.Getenv("TELEGRAM_API_KEY")

	chatBotApiKey := ""
	tgBotApiKey := ""

	return Config{
		ChatBotApiKey:  chatBotApiKey,
		TelegramApiKey: tgBotApiKey,
		ResponseTokens: 40,
		WebPort:        8080,
	}
}
