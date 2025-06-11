package tgClient

import (
	"awesomeProject/internal/domain"
	"encoding/json"
	"net/http"
)

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"` // Message can be nil
}

type UpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

func (u *Update) Parse(r *http.Request) (*domain.ServiceMessage, error) {

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return &domain.ServiceMessage{}, err
	}

	text := u.Message.Text
	id := u.Message.Chat.ID
	maxToken := 40

	model := "deepseek-chat"

	message := domain.ServiceMessage{RequestText: text, ChatId: id, Model: model, MaxToken: maxToken}

	return &message, nil
}
