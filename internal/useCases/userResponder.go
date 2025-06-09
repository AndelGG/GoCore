package useCases

import "awesomeProject/internal/domain"

type EventSender interface {
	SendEvent(msg *domain.ServiceMessage) error
}

type TelegramResponser interface {
	SendMessage(msg *domain.ServiceMessage, chatId string) error
}

type EventSenderUseCase struct {
}

func (e EventSenderUseCase) SendEvent(msg *domain.ServiceMessage) error {
	return nil
}

func NewEventSenderUseCase() {

}
