package usecases

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/lib/logger/sl"
	"fmt"
	"log/slog"
)

type ChatBotResponder struct {
	Responder ChatBot
	log       *slog.Logger
}

type ResponderUseCase interface {
	SendMessage(message *domain.Message) (*domain.Message, error)
}

type ChatBot interface {
	RequestToChatBot(message *domain.ServiceMessage) (domain.ResponseScheme, error)
}

func New(responder ChatBot, log *slog.Logger) *ChatBotResponder {
	return &ChatBotResponder{
		responder, log,
	}
}

func (u *ChatBotResponder) SendMessage(message *domain.ServiceMessage) (*domain.ServiceMessage, error) {

	op := "useCase.chatBotResponder.SendMessage"

	obj, err := u.Responder.RequestToChatBot(message)

	if err != nil {
		u.log.Warn("request to chatbot error", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	message.Request = obj.Choices[0].Message.Content

	return message, nil
}
