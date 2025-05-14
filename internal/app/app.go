package app

import (
	"awesomeProject/internal/app/telegram"
	"awesomeProject/internal/app/web"
	"awesomeProject/internal/config"
	"awesomeProject/internal/infrastructure/deepSeek"
	"awesomeProject/internal/useCases"
	"log/slog"
)

type Server struct {
	log       *slog.Logger
	WebServer *web.App
	TgServer  *telegram.App
}

func New(cfg config.Config, log *slog.Logger) *Server {
	chatbot := deepSeek.New(cfg.ChatBotApiKey)

	use := useCases.New(chatbot, log)

	req := web.New(cfg.WebPort, use, log)
	tg := telegram.New(cfg.TelegramApiKey, use)

	return &Server{WebServer: req, TgServer: tg}
}
