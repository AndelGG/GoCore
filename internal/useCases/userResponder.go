package useCases

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
)

type MainResponser interface {
	SendMessage(ctx context.Context, chatId int, message string) error
}

type EventSender interface {
	SendEvent(ctx context.Context, msg *domain.ServiceMessage) error
}

type TelegramResponder interface {
	SendMessage(ctx context.Context, chatId int, message string) error
	SendSticker(ctx context.Context, chatId int, message string) error
}

type EventSenderUseCase struct {
	log     *slog.Logger
	Eventer MainResponser
}

type TelegramReplyUseCase struct {
	Log         *slog.Logger
	TgResponder TelegramResponder
}

func (e TelegramReplyUseCase) SendSticker(ctx context.Context, msg *domain.ServiceMessage) error {
	//TODO implement me
	panic("implement me")
}

func (e TelegramReplyUseCase) CatchError(ctx context.Context, err error) error {

	// TODO: think about message object
	op := "useCase.TelegramReplyUseCase.CatchError"
	err := e.TgResponder.SendMessage(ctx, msg.ChatId, err.Error())
	if err != nil {
		e.Log.Warn("reply tg error", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (e EventSenderUseCase) SendEvent(ctx context.Context, msg *domain.ServiceMessage) error {

	op := "useCase.EventSenderUseCase.SendEvent"
	err := e.Eventer.SendMessage(ctx, msg.ChatId, msg.Response)
	if err != nil {
		e.log.Warn("catch event", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (e TelegramReplyUseCase) SendMessage(ctx context.Context, msg *domain.ServiceMessage) error {

	// TODO: think about message object
	op := "useCase.TelegramReplyUseCase.SendMessage"
	err := e.TgResponder.SendMessage(ctx, msg.ChatId, msg.Response)
	if err != nil {
		e.Log.Warn("send tg message", sl.Err(err))
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
