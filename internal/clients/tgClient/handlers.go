package tgClient

import (
	"awesomeProject/internal/domain"
	"context"
	"errors"
	"net/http"
)

var (
	ErrParse   = errors.New("the message does not match the pattern")
	ErrChatBot = errors.New("chatbot returns an error")
)

type TelegramHandler interface {
	WebHookHandler(w http.ResponseWriter, r *http.Request)
}

type TelegramReply interface {
	SendMessage(ctx context.Context, msg *domain.ServiceMessage) error
	CatchError(ctx context.Context, err error) error
	SendSticker(ctx context.Context, msg *domain.ServiceMessage) error
}

type ChatBotResponderUseCase interface {
	SendMessage(ctx context.Context, message *domain.ServiceMessage) (*domain.ServiceMessage, error)
}

type Requester struct {
	request ChatBotResponderUseCase
	respond TelegramReply
}

func New(request ChatBotResponderUseCase, respond TelegramReply) *Requester {
	return &Requester{request: request, respond: respond}
}

func (q *Requester) WebHookHandler(w http.ResponseWriter, r *http.Request) {
	var raw Update

	if err := json.NewDecoder(r.Body).Decode(raw); err != nil {
		return &domain.ServiceMessage{}, err
	}

	ctx := domain.ServiceMessage{}

	if err != nil {
		q.respond.CatchError(ErrParse)
	}

	go func() {

	}()

	w.WriteHeader(http.StatusOK)
}

func (u *Update) Parse(r *http.Request) (*domain.ServiceMessage, error) {

	text := u.Message.Text
	id := u.Message.Chat.ID
	maxToken := 40

	model := "deepseek-chat"

	message := domain.ServiceMessage{RequestText: text, ChatId: id, Model: model, MaxToken: maxToken}

	return &message, nil
}
