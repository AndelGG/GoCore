package tgClient

import (
	"awesomeProject/internal/domain"
	"fmt"
	"net/http"
)

type TelegramHandler interface {
	WebHookHandler(w http.ResponseWriter, r *http.Request)
}

type TelegramReply interface {
	SendMessage(msg *domain.ServiceMessage) error
	SendSticker(msg *domain.ServiceMessage) error
}

type ChatBotResponderUseCase interface {
	SendMessage(message *domain.ServiceMessage) (*domain.ServiceMessage, error)
}

type Requester struct {
	request ChatBotResponderUseCase
	respond TelegramReply
}

func New(request ChatBotResponderUseCase, respond TelegramReply) *Requester {
	return &Requester{request: request, respond: respond}
}

func (q *Requester) WebHookHandler(w http.ResponseWriter, r *http.Request) {
	var u Update

	rawMsg, err := u.Parse(r)

	if err != nil {
		panic(err)
		// tg err
	}

	go func() {
		resp, err := q.request.SendMessage(rawMsg)
		if err != nil {
			// tg err
		}

		err = q.respond.SendMessage(resp)
		if err != nil {
			// tg err
		}

		fmt.Println("tg resp: ", resp)

	}()

	fmt.Println("ok")
	w.WriteHeader(http.StatusOK)
}
