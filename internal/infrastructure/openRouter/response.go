package openRouter

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/lib/req"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatBot struct {
	Api string
}

func New(api string) *ChatBot {
	return &ChatBot{api}
}

func (c *ChatBot) RequestToChatBot(message *domain.ServiceMessage) (ResponseScheme, error) {

	const op = "infrastructure.openRouter.RequestToChatBot"

	req, err := req.CreateRequest(message, c.Api)
	if err != nil {
		return ResponseScheme{}, err
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var responseObject ResponseScheme
	if err := json.Unmarshal(body, &responseObject); err != nil {
		fmt.Println(fmt.Sprintf("body: %b", body))
		return ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}

	return responseObject, nil

}
