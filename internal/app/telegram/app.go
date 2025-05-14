package telegram

import (
	"awesomeProject/internal/clients/tgClient"
	"awesomeProject/internal/controller"
	"awesomeProject/internal/controller/telegram"
	"awesomeProject/internal/domain"
	"fmt"
	"log"
	"time"
)

type App struct {
	fetcher   controller.Fetcher
	processor controller.Processor
	batchSize int
}

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func New(token string, use domain.ResponderUseCase) *App {

	telegramClient := tgClient.New(tgBotHost, token)

	eventsProcessor := telegram.New(telegramClient, use)

	return &App{
		fetcher:   eventsProcessor,
		processor: eventsProcessor,
		batchSize: batchSize,
	}
}

func (a *App) MustRun() error {
	fmt.Println("TelegramBot is running")
	for {
		gotEvents, err := a.fetcher.Fetch(a.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := a.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

/*
	1. Потеря событий: ретраи, возвращение в хранилище, фоллбэк, подтверждение
	2. Обработка всей пачки: остановка после ошибки, счетчик ошибок
	3. Параллельная обработка: sync.WaitGroup
*/

func (a *App) handleEvents(events []controller.Event) error {
	for _, event := range events {
		log.Printf("got new event %s", event.Text)
		if err := a.processor.Process(event); err != nil {
			log.Printf("cant handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
