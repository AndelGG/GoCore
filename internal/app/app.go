package app

import (
	"awesomeProject/internal/app/telegram"
	"awesomeProject/internal/app/web"
	"awesomeProject/internal/infrastructure/deepSeek"
	"awesomeProject/internal/usecases"
	"log/slog"
)

const tgBotHost = "api.telegram.org"

type Server struct {
	log       *slog.Logger
	WebServer *web.App
	TgServer  *telegram.App
}

func New(port int, token string, log *slog.Logger) *Server {
	chatbot := deepSeek.New(token)

	use := usecases.New(log, chatbot)

	req := web.New(port, use)
	tg := telegram.New(tgBotHost, use)

	return &Server{WebServer: req, TgServer: tg}
}
