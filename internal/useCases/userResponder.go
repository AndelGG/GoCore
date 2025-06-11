package useCases

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/lib/logger/sl"
	"fmt"
	"log/slog"
)

type MainResponser interface {
	SendMessage(chatId int, message string) error
}

type EventSender interface {
	SendEvent(msg *domain.ServiceMessage) error
}

type TelegramResponder interface {
	SendMessage(chatId int, message string) error
	SendSticker(chatId int, message string) error
}

type EventSenderUseCase struct {
	log     *slog.Logger
	Eventer MainResponser
}

type TelegramReplyUseCase struct {
	Log         *slog.Logger
	TgResponder TelegramResponder
}

func (e TelegramReplyUseCase) SendSticker(msg *domain.ServiceMessage) error {
	//TODO implement me
	panic("implement me")
}

func (e EventSenderUseCase) SendEvent(msg *domain.ServiceMessage) error {

	op := "useCase.chatBotResponder.SendMessage"
	err := e.Eventer.SendMessage(msg.ChatId, msg.RequestText)
	if err != nil {
		e.log.Warn("request to chatbot error", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (e TelegramReplyUseCase) SendMessage(msg *domain.ServiceMessage) error {

	// TODO: think about message object
	op := "useCase.chatBotResponder.SendMessage"
	err := e.TgResponder.SendMessage(msg.ChatId, msg.Response)
	if err != nil {
		e.Log.Warn("request to chatbot error", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func NewEventSender(log *slog.Logger, eventer MainResponser) *EventSenderUseCase {
	return &EventSenderUseCase{log: log, Eventer: eventer}
}

func NewTelegramReply(log *slog.Logger, tgResponder TelegramResponder) *TelegramReplyUseCase {
	return &TelegramReplyUseCase{Log: log, TgResponder: tgResponder}
}
