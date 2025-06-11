package telegram

import (
	r "awesomeProject/internal/lib/req"
	"fmt"
	"net/url"
	"path"
	"strconv"
)

const telegramApi = "api.telegram.org"

const (
	sendMessage = "sendMessage"
)

type Api struct {
	Api string
}

func New(api string) *Api {
	return &Api{api}
}

func (t Api) SendMessage(chatId int, message string) error {
	// TODO: hide
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", message)

	_, err := r.MakeGetRequest(t.makeTgAction(sendMessage), q)
	if err != nil {
		return err
	}

	return nil
}

func (t Api) SendSticker(chatId int, message string) error {
	return fmt.Errorf("sosi sosi")
}

func (t Api) makeTgAction(action string) string {

	u := url.URL{
		Scheme: "https",
		Host:   telegramApi,
		Path:   path.Join("bot"+t.Api, action),
	}

	return u.String()
}
