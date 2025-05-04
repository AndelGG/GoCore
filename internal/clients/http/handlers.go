package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Requester struct {
	request MessageUseCase
}

type RequestData struct {
	Message string `json:"message"`
}

func New(useCase MessageUseCase) *Requester {
	return &Requester{useCase}
}

// TODO: разделить логику хэндлера

func (q *Requester) ResponseHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Unable to parse", http.StatusBadRequest)
	}

	fmt.Fprintf(w, q.request.Message(data.Message))
}
