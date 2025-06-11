package useCases

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/infrastructure/deepSeek"
	"awesomeProject/internal/lib/logger/sl"
	"fmt"
	"log/slog"
)

type ChatBotResponder struct {
	Responder ChatBot
	log       *slog.Logger
}

type ChatBot interface {
	RequestToChatBot(text, model string, maxToken int) (deepSeek.ResponseScheme, error)
}

func NewChatBotResponder(responder ChatBot, log *slog.Logger) *ChatBotResponder {
	return &ChatBotResponder{
		responder, log,
	}
}

func (u ChatBotResponder) SendMessage(message *domain.ServiceMessage) (*domain.ServiceMessage, error) {

	op := "useCase.chatBotResponder.SendMessage"
	obj, err := u.Responder.RequestToChatBot(message.RequestText, message.Model, message.MaxToken)

	if err != nil {
		u.log.Warn("request to chatbot error", sl.Err(err))
		return message, fmt.Errorf("%s: %w", op, err)
	}

	message.Response = obj.Choices[0].Message.Content
	return message, nil
}
