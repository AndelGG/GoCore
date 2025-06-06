package telegram

import (
	"awesomeProject/internal/clients/tgClient"
	"awesomeProject/internal/controller"
	"awesomeProject/internal/controller/telegram"
	"awesomeProject/internal/domain"
	"log/slog"
	"time"
)

type App struct {
	fetcher   controller.Fetcher
	processor controller.Processor
	batchSize int
	log       *slog.Logger
}

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func New(token string, use domain.ResponderUseCase, log *slog.Logger) *App {

	telegramClient := tgClient.New(tgBotHost, token)

	eventsProcessor := telegram.New(telegramClient, use)

	return &App{
		fetcher:   eventsProcessor,
		processor: eventsProcessor,
		batchSize: batchSize,
		log:       log,
	}
}

func (a *App) Run() error {
	const op = "App.Tg.Run"

	log := a.log.With(slog.String("op", op))

	for {
		gotEvents, err := a.fetcher.Fetch(a.batchSize)
		if err != nil {
			log.Warn("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := a.handleEvents(gotEvents, log); err != nil {
			log.With(err)

			continue
		}
	}
}

/*
	1. Потеря событий: ретраи, возвращение в хранилище, фоллбэк, подтверждение
	2. Обработка всей пачки: остановка после ошибки, счетчик ошибок
	3. Параллельная обработка: sync.WaitGroup
*/

func (a *App) handleEvents(events []controller.Event, log *slog.Logger) error {
	for _, event := range events {
		log.Info("got new event %s", event.Text)
		if err := a.processor.Process(event); err != nil {
			log.Warn("cant handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
