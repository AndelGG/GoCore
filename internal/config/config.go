package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ChatBotApiKey    string `yaml:"ChatBotApiKey"`
	TelegramApiKey   string `yaml:"TelegramApiKey"`
	OpenRouterApiKey string `yaml:"OpenRouterApiKey"`
	WebPort          int    `yaml:"WebPort"`
}

func MustLoad() Config {
	var conf Config

	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(file, conf); err != nil {
		panic(err)
	}

	return Config{
		ChatBotApiKey:    conf.ChatBotApiKey,
		TelegramApiKey:   conf.TelegramApiKey,
		OpenRouterApiKey: conf.OpenRouterApiKey,
		WebPort:          conf.WebPort,
	}
}
