package rest

import (
	"awesomeProject/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponderUseCase interface {
	SendMessage(message *domain.ServiceMessage) (*domain.ServiceMessage, error)
}
type Requester struct {
	request ResponderUseCase
}

type RequestData struct {
	Message string `json:"message"`
}

func New(useCase ResponderUseCase) *Requester {
	return &Requester{request: useCase}
}

func (q *Requester) ResponseHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Unable to parse", http.StatusBadRequest)
	}

	resp := domain.ServiceMessage{Response: data.Message, MaxToken: 40}
	var msg, err = q.request.SendMessage(&resp)

	if err != nil {
		http.Error(w, "Unable to parse", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, msg.Response)
}
