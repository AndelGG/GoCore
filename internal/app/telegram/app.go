package telegram

import (
	"awesomeProject/internal/clients/tgClient"
	"awesomeProject/internal/infrastructure/telegram"
	"awesomeProject/internal/useCases"
	"log/slog"
	"net/http"
	"strconv"
)

type App struct {
	request tgClient.TelegramHandler
	log     *slog.Logger
}

func New(chat tgClient.ChatBotResponderUseCase, log *slog.Logger, api string) *App {

	tg := telegram.New(api)

	reply := useCases.NewTelegramReply(log, tg)

	request := tgClient.New(chat, reply)

	return &App{request: request, log: log}
}

func (a *App) MustRun() {
	const op = "webApp.Run"

	port := 8000

	addr := createAddress(port)

	log := a.log.With(slog.String("op", op), slog.Int("port", port))

	http.HandleFunc("/", a.request.WebHookHandler)

	log.Info("starting TgServer", slog.String("address:", addr))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func createAddress(port int) string {
	return "localhost:" + strconv.Itoa(port)
}
