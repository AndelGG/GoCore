package app

import (
	"awesomeProject/internal/app/telegram"
	//"awesomeProject/internal/app/web"
	"awesomeProject/internal/config"
	"awesomeProject/internal/infrastructure/deepSeek"
	"awesomeProject/internal/useCases"
	"log/slog"
)

type Server struct {
	log *slog.Logger
	//WebServer *web.App
	TgServer *telegram.App
}

func New(cfg config.Config, log *slog.Logger) *Server {
	chatbot := deepSeek.New(cfg.ChatBotApiKey)

	useChat := useCases.NewChatBotResponder(chatbot, log)

	//req := web.New(cfg.WebPort, useChat, log)
	tg := telegram.New(useChat, log, cfg.TelegramApiKey)

	return &Server{TgServer: tg}
}
