package tgClient

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/lib/jsonParser"
	"fmt"
	"net/http"
)

type ResponderUseCase interface {
	SendMessage(message *domain.ServiceMessage) (*domain.ServiceMessage, error)
}
type Requester struct {
	request ResponderUseCase
}

func New(useCase ResponderUseCase) *Requester {
	return &Requester{request: useCase}
}

func (q *Requester) WebHookHandler(w http.ResponseWriter, r *http.Request) {

	msg, err := jsonParser.JsonParser(r)
	if err != nil {
		// tg err
	}

	resp, err := q.request.SendMessage(&msg)
	if err != nil {
		// tg err
	}

	// tg resp
	fmt.Print("tg resp: ", resp)

	w.WriteHeader(http.StatusOK)
}
